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
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"sync/atomic"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/codec"
	ssdexec "github.com/chronowave/chronowave/ssd/exec"
	ssdidx "github.com/chronowave/chronowave/ssd/index"
	"github.com/chronowave/chronowave/ssql"
)

func Build(json string, timestamp string, keys []string) error {
	data, err := ioutil.ReadFile(json)
	if err != nil {
		return err
	}

	return createIndex(0, data, timestamp, keys)
}

func createIndex(nid int64, data []byte, timestamp string, keys []string) error {
	parsed, err := codec.ParseJson(data)
	if err != nil {
		return err
	}

	indexed, err := ssdidx.Build(parsed)
	if err != nil {
		return err
	}

	dir := filepath.Join(Directory, index)
	tmp, err := ioutil.TempFile(dir, "tmp")
	if err != nil {
		return err
	}
	defer tmp.Close()

	data, err = ssdidx.EncodeIndexBlock(indexed)
	if err != nil {
		return err
	}

	_, err = tmp.Write(data)
	if err != nil {
		return err
	}
	if err = tmp.Close(); err != nil {
		return err
	}

	for {
		if nid == 0 {
			nid = atomic.AddInt64(&seq, 1)
		}
		name := fmt.Sprintf("%016X", nid)
		path := filepath.Join(Directory, index, name[:4], name[8:12])
		os.MkdirAll(path, os.ModePerm)
		if f, err := os.OpenFile(filepath.Join(path, name), os.O_CREATE|os.O_EXCL, os.ModePerm); err == nil {
			f.Close()
			os.Rename(tmp.Name(), f.Name())
			break
		} else if !os.IsExist(err) {
			panic(err)
		}
	}

	attributes := make([]*ssql.Attribute, len(keys)+1)
	expr := make([]*ssql.Expr, len(keys)+1)
	attributes[0] = &ssql.Attribute{Name: "ts"}
	expr[0] = &ssql.Expr{Field: &ssql.Expr_Tuple{
		Tuple: &ssql.Tuple{Name: "ts", Path: timestamp},
	}}

	inverted := make([]map[string]bool, len(keys))
	for i, key := range keys {
		inverted[i] = map[string]bool{}
		name := strconv.FormatInt(int64(i), 10)
		attributes[i+1] = &ssql.Attribute{Name: name}
		expr[i+1] = &ssql.Expr{Field: &ssql.Expr_Tuple{
			Tuple: &ssql.Tuple{Name: name, Path: key},
		}}
	}

	stmt := &ssql.Statement{Find: attributes, Where: expr}

	rs := ssdexec.Exec(indexed, stmt)

	var (
		min = int64(math.MaxInt64)
		max = int64(math.MinInt64)
	)

	for i := range rs.RowId {
		if rs.Column[0].RowIdx[i] > 0 {
			if int64(rs.Column[0].Value[i]) > max {
				max = int64(rs.Column[0].Value[i])
			}

			if int64(rs.Column[0].Value[i]) < min {
				min = int64(rs.Column[0].Value[i])
			}
		}
	}

	if min == math.MinInt64 || max == math.MinInt64 {
		os.Remove(tmp.Name())
		return errors.New("invalid timestamp path [" + timestamp + "]")
	}

	err = insertWave(nid, min, max)
	if err != nil {
		return err
	}

	for i := 1; i < len(rs.ColumnType); i++ {
		var onValue func(v uint64)
		switch rs.ColumnType[i] {
		case ssd.TEXT:
			onValue = func(v uint64) {
				b := rs.Text[v]
				if len(b) > 0 {
					inverted[i-1][string(rs.Text[v])] = true
				}
			}
		case ssd.INT64:
			onValue = func(v uint64) {
				inverted[i-1][strconv.FormatUint(v, 16)] = true
			}
		default:
			continue
		}

		for j := range rs.RowId {
			if rs.Column[i].RowIdx[j] > 0 {
				onValue(rs.Column[i].Value[j])
			}
		}
	}

	for i, path := range keys {
		key := make([]string, len(inverted[i]))
		j := 0
		for k := range inverted[i] {
			key[j] = k
			j++
		}
		err = insertWaveLoc(path, key, nid)
		if err != nil {
			return err
		}
	}

	return nil
}
