/*
 *  Copyright 2020 ChronoWave Authors
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  Package parser declares an expression parser with support for macro
 *  expansion.
 */

package embed

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	hfmi "github.com/rleiwang/hfmi/ctor"

	"github.com/chronowave/chronowave/ssd/codec"
	ssdidx "github.com/chronowave/chronowave/ssd/index"
	"github.com/chronowave/chronowave/ssql/parser"

	"github.com/rs/zerolog/log"
)

const (
	sz = 256
)

func init() {
	hfmi.SetSegmentCache(2 * 1024 * 1024)
}

type WaveStream struct {
	wal      string
	cnt      uint64
	batch    []string
	c        chan []string
	timer    *time.Timer
	dbcloser func()
	lock     sync.Mutex
}

func NewWave(dir, ts string, keys []string) *WaveStream {
	Directory = dir
	closer, err := Verify()
	if err != nil {
		panic(err)
	}

	buildWalIndex()
	timer := time.NewTimer(15 * time.Second)
	go func() {
		for range timer.C {
			buildWalIndex()
		}
	}()

	c := make(chan []string, 2048)
	go func() {
		for files := range c {
			buildIndex(ts, keys, files)
		}
	}()

	return &WaveStream{
		wal:      filepath.Join(dir, wal),
		batch:    make([]string, 0, sz),
		c:        c,
		timer:    timer,
		dbcloser: closer,
	}
}

func (s *WaveStream) Close() {
	s.timer.Stop()
	close(s.c)
	s.dbcloser()
}

func (s *WaveStream) OnNewDocument(json []byte) error {
	n := filepath.Join(s.wal, strconv.FormatUint(atomic.AddUint64(&s.cnt, 1), 10))
	err := ioutil.WriteFile(n, json, os.ModePerm)
	if err != nil {
		return err
	}
	s.lock.Lock()
	defer s.lock.Unlock()

	s.batch = append(s.batch, n)
	if len(s.batch) >= sz {
		tmp := s.batch
		s.batch = make([]string, 0, sz)
		s.batch = append(s.batch, tmp[sz:]...)
		s.c <- tmp[:sz]
	}
	return nil
}

func (s *WaveStream) Purge(ctx context.Context, time time.Time) error {
	return Purge(ctx, time)
}

func (s *WaveStream) Query(ctx context.Context, query string) ([]byte, error) {
	stmt, errs := parser.Parse(query)
	if len(errs) > 0 {
		return nil, errors.New("syntax error")
	}

	return Query(ctx, stmt), nil
}

func buildIndex(timestamp string, keys []string, files []string) {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("runtime error in building index %v", r)
		}
	}()
	if id, err := BuildFromFiles(files, timestamp, keys); err == nil {
		walIndex = nil
		for _, f := range files {
			os.Remove(f + "." + strconv.FormatInt(id, 10))
		}
	} else {
		log.Info().Msgf("restore due to error %v\n", err)
		for _, f := range files {
			os.Rename(f+"."+strconv.FormatInt(id, 10), f)
		}
	}
}

func buildWalIndex() {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("runtime error in building index %v", r)
		}
	}()

	files, err := ioutil.ReadDir(filepath.Join(Directory, wal))
	if err != nil {
		log.Err(err).Msg("reading WAL files")
		return
	}

	buf := pool.Get().([]byte)
	defer pool.Put(buf)

	w := bytes.NewBuffer(buf[:0])
	for _, f := range files {
		if f.IsDir() || len(filepath.Ext(f.Name())) > 0 {
			continue
		}
		if r, err := os.OpenFile(filepath.Join(Directory, wal, f.Name()), os.O_RDONLY, os.ModePerm); err == nil {
			io.Copy(w, r)
			r.Close()
		}
	}

	if w.Len() == 0 {
		return
	}

	parsed, err := codec.ParseJson(w.Bytes())
	if err != nil {
		log.Err(err).Msg("parsing WAL files")
		return
	}

	indexed, err := ssdidx.Build(parsed)
	if err != nil {
		log.Err(err).Msg("building index from WAL files")
	} else {
		walIndex = indexed
	}
}
