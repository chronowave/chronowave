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
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strconv"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/codec"
	ssdexec "github.com/chronowave/chronowave/ssd/exec"
	ssdidx "github.com/chronowave/chronowave/ssd/index"
	"github.com/chronowave/chronowave/ssd/operator"
	"github.com/chronowave/chronowave/ssql"

	"github.com/rs/zerolog/log"
)

type void struct{}

var (
	empty    = []byte("[]")
	walIndex *ssd.IndexedBlock
	guard    = make(chan void, runtime.NumCPU())
)

func Query(ctx context.Context, stmt *ssql.Statement) []byte {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("runtime error in query %v, trace: %v", r, string(debug.Stack()))
		}
	}()

	wid, err := selectIndexIds(stmt)
	if err != nil {
		log.Err(err).Msg("query for ids")
	}

	files := make([]string, len(wid))
	for i, id := range wid {
		name := fmt.Sprintf("%016X", id)
		files[i] = filepath.Join(Directory, index, name[:4], name[8:12], name)
	}

	rss := make([]*ssd.ResultSet, len(wid)+1)
	done := make(chan void, len(rss))
	queryWAL(ctx, stmt, 0, rss, done)
	for i, f := range files {
		queryIndexFile(ctx, f, stmt, i+1, rss, done)
	}

	for range rss {
		<-done
	}

	j := 0
	for i := range rss {
		if rss[i] != nil {
			rss[j] = rss[i]
			j++
		}
	}

	if j == 0 {
		return empty
	}

	rs := operator.Merge(rss[:j])

	if len(stmt.OrderBy) > 0 {
		operator.OrderBy(rs, stmt)
	}

	return codec.MarshalResultSet(rs, stmt.GetLimit())
}

func selectIndexIds(stmt *ssql.Statement) ([]int64, error) {
	for _, expr := range stmt.Where {
		if expr.GetTuple() != nil && expr.GetTuple().GetPredicate() != nil {
			pred := expr.GetTuple().GetPredicate()
			switch pred.(type) {
			case *ssql.Tuple_Timeframe:
				beg := pred.(*ssql.Tuple_Timeframe).Timeframe.First.Value.(*ssql.Operand_Int).Int
				end := pred.(*ssql.Tuple_Timeframe).Timeframe.Second.Value.(*ssql.Operand_Int).Int
				return selectWave(beg, end)
			case *ssql.Tuple_Key:
				var key string
				switch pred.(*ssql.Tuple_Key).Key.First.Value.(type) {
				case *ssql.Operand_Int:
					key = strconv.FormatInt(pred.(*ssql.Tuple_Key).Key.First.Value.(*ssql.Operand_Int).Int, 16)
				case *ssql.Operand_Text:
					key = pred.(*ssql.Tuple_Key).Key.First.Value.(*ssql.Operand_Text).Text
				}
				return selectKey(expr.GetTuple().GetPath(), key)
			}
		}
	}

	return []int64{}, nil
}

func queryIndexFile(ctx context.Context, f string, stmt *ssql.Statement, i int, rss []*ssd.ResultSet, done chan void) {
	select {
	case guard <- void{}:
		go func() {
			runtime.LockOSThread()
			defer func() {
				runtime.UnlockOSThread()
				<-guard
				done <- void{}
				if err := recover(); err != nil {
					log.Error().Msgf("worker has error '%v', trace %v", err, string(debug.Stack()))
				}
			}()

			data, err := ioutil.ReadFile(f)
			if err != nil {
				fmt.Printf("skipped index file %v due to %v\n", f, err)
				return
			}

			idx, err := ssdidx.DecodeIndexBlock(data)
			if err != nil {
				fmt.Printf("skipped index file %v, decode err %v\n", f, err)
				return
			}

			if rs := ssdexec.Exec(idx, stmt); rs != nil {
				rss[i] = rs
			}
		}()
	case <-ctx.Done():
	}
}

func queryWAL(ctx context.Context, stmt *ssql.Statement, i int, rss []*ssd.ResultSet, done chan void) {
	select {
	case guard <- void{}:
		go func() {
			runtime.LockOSThread()
			defer func() {
				runtime.UnlockOSThread()
				<-guard
				done <- void{}
				if err := recover(); err != nil {
					log.Error().Msgf("worker has error '%v', trace %v", err, string(debug.Stack()))
				}
			}()
			indexed := walIndex
			if indexed != nil {
				rss[i] = ssdexec.Exec(indexed, stmt)
			}
		}()
	case <-ctx.Done():
	}
}
