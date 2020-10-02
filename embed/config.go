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
	"os"
	"path/filepath"
)

var (
	Directory string
)

var (
	wal       = "wal"
	index     = "index"
	donothing = func() {}
)

func Verify() (func(), error) {
	err := os.MkdirAll(filepath.Join(Directory, wal), os.ModePerm)
	if err != nil {
		return donothing, err
	}
	err = os.MkdirAll(filepath.Join(Directory, index), os.ModePerm)
	if err != nil {
		return donothing, err
	}

	return func() { closeSqlite() }, openSqlite(Directory)
}
