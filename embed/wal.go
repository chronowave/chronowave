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
	"io"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

var (
	pool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 10*1024*1024)
		},
	}
)

func BuildFromFiles(files []string, timestamp string, keys []string) (int64, error) {
	buf := pool.Get().([]byte)
	defer pool.Put(buf)

	nid := atomic.AddInt64(&seq, 1)
	w := bytes.NewBuffer(buf[:0])
	for _, f := range files {
		r, err := os.OpenFile(f, os.O_RDONLY, os.ModePerm)
		if err != nil {
			continue
		}
		io.Copy(w, r)
		r.Close()
		os.Rename(f, f+"."+strconv.FormatInt(nid, 10))
	}

	return nid, createIndex(nid, w.Bytes(), timestamp, keys)
}
