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

package ssd

import (
	"bytes"

	"github.com/rleiwang/hfmi"

	"github.com/chronowave/chronowave/ssd/aggregator"
)

const (
	MaxDoc = 1 << 16
)

type Columnar struct {
	Float64 []float64
	Int64   []int64
	Bool    []bool
	Text    []uint32
}

type ParsedBlock struct {
	Count    uint32
	Meta     *EntityMeta
	Entity   *bytes.Buffer
	Columnar Columnar
	Content  *TextColumnar
}

type HeaderISA struct {
	Entity    []uint16
	Attribute []uint16
}

type HeaderLookupTable struct {
	Text    HeaderISA
	Float64 HeaderISA
	Int64   HeaderISA
	Bool    HeaderISA
	Null    HeaderISA
}

type IndexedBlock struct {
	ID       uint64
	EntityID []uint32
	Meta     *EntityMeta
	Entity   hfmi.FMI
	Columnar Columnar
	Content  hfmi.FMI
	HLT      HeaderLookupTable
	HeaderDA []uint32
	FragDA   []uint32
}

type Column struct {
	Name   string
	RowIdx []byte
	Value  []uint64
}

type ResultSet struct {
	RowId      []uint64
	ColumnType []byte
	Column     []Column
	Text       [][]byte
	Json       [][]byte
	Aggregate  []aggregator.Aggregator
	Order      []int
}
