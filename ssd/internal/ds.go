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

package internal

import (
	"github.com/chronowave/chronowave/ssql"
)

var (
	FragmentSize = 128
)

type Document struct {
	Entity    []uint16
	Attribute []uint16
}

type Column struct {
	Key           []byte
	Ok            bool
	Name          string
	Group         bool
	Type          byte
	Func          *ssql.Function
	ColumnText    *ColumnText
	ColumnFloat64 *ColumnFloat64
	ColumnInt64   *ColumnInt64
	ColumnBool    *ColumnBool
	ColumnNull    *ColumnNull
	ColumnJson    *ColumnJson
}

type ColumnBool struct {
	Columnar []bool
	Rows     []uint32
	Cols     []uint32
}

type ColumnNull struct {
	Rows []uint32
	Cols []uint32
}

type ColumnInt64 struct {
	Columnar []int64
	Rows     []uint32
	Cols     []uint32
}

type ColumnFloat64 struct {
	Columnar []float64
	Rows     []uint32
	Cols     []uint32
}

type ColumnText struct {
	Columnar map[uint32][]byte
	Rows     []uint32
	Cols     []uint32
}

type ColumnJson struct {
	Rows []uint32
	Cols [][]byte
}

type Attribute struct {
	Code      []byte
	Offset    uint16
	ValueType byte
	Value     uint64
}
