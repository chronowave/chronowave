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

package operator

import (
	ext "github.com/chronowave/ext/operator"
	"github.com/rleiwang/hfmi"

	"github.com/rs/zerolog/log"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/internal"
)

// Select returns entity attribute value in result set
func Select(index *ssd.IndexedBlock, columns []internal.Column, entity []uint16) []internal.Column {
	_, sohEND, ok := index.Entity.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}
	sohEND++

	var rank []uint
	for i := range columns {
		if !columns[i].Ok {
			columns[i].Type = ssd.NULL
			continue
		}

		if len(columns[i].Key) == 0 {
			columns[i].Type = ssd.JSON
			columns[i].ColumnJson = selectJson(index, columns[i].Key, columns[i].Func == nil, entity)
			continue
		}

		// rank is [s, e)
		columns[i].Type, rank = findKeyTypeRank(index.Entity, sohEND, columns[i].Key)
		if len(rank) == 0 {
			columns[i].Type = ssd.NULL
			continue
		}

		// ssd.VALUES = []byte{TEXT, FLT32, FLT64, INT8, INT16, INT32, INT64, BTS, BFS, NULL}
		switch columns[i].Type {
		case ssd.TEXT:
			columns[i].ColumnText = selectText(index.Content, rank, &index.HLT.Text, index.Columnar.Text, entity)
		case ssd.FLT32:
			// ignore
		case ssd.FLT64:
			columns[i].ColumnFloat64 = selectFloat64(rank, &index.HLT.Float64, index.Columnar.Float64, entity)
		case ssd.INT8:
			// ignore
		case ssd.INT16:
			// ignore
		case ssd.INT32:
			// ignore
		case ssd.INT64:
			columns[i].ColumnInt64 = selectInt64(rank, &index.HLT.Int64, index.Columnar.Int64, entity)
		case ssd.BOOL:
			columns[i].ColumnBool = selectBool(rank, &index.HLT.Bool, index.Columnar.Bool, entity)
		case ssd.NULL:
			columns[i].ColumnNull = selectNull(rank, &index.HLT.Null, entity)
		default:
			columns[i].Type = ssd.JSON
			columns[i].ColumnJson = selectJson(index, columns[i].Key, columns[i].Func == nil, entity)
		}
	}

	return columns
}

// findKeyTypeRank returns key data type and rank
func findKeyTypeRank(fmi hfmi.FMI, sohEND uint, key []byte) (byte, []uint) {
	// in SOH block, note: bound is (s, e]
	bound := findEndRange(fmi, key[0], sohEND)
	if len(bound) == 0 {
		return ssd.NULL, bound
	}

	for _, k := range key[1:] {
		if bound = findBound(fmi, k, bound); len(bound) == 0 {
			return ssd.NULL, bound
		}
	}

	// bound is (s, e]
	cc, _, ok := fmi.Access(bound[end])
	if !ok {
		log.Error().Msgf("failed to access entity end of bound %v", bound)
		return ssd.NULL, nil
	}

	s, _, ok := fmi.GetBound(cc)
	if !ok {
		log.Error().Msgf("entity control char [%v] is not found", cc)
		return ssd.NULL, nil
	}

	if bound = findBound(fmi, cc, bound); len(bound) == 0 {
		return ssd.NULL, bound
	}

	return cc, []uint{bound[beg] - s, bound[end] - s}
}

func selectText(fmi hfmi.FMI, rank []uint, header *ssd.HeaderISA, columnar []uint32, entity []uint16) *internal.ColumnText {
	sz := rank[end] - rank[beg]
	if sz > uint(len(entity)) {
		sz = uint(len(entity))
	}

	flat := make([]uint32, 2*sz)
	ix, iy := flat[:sz], flat[sz:]
	// IntersectSortedUint16(x, y, ix, iy) returns intersection size
	cnt := ext.IntersectSortedUint16(header.Entity[rank[beg]:rank[end]], entity, ix, iy)

	column := &internal.ColumnText{
		Columnar: make(map[uint32][]byte, cnt),
		Rows:     iy[:cnt],
		Cols:     ix[:cnt],
	}

	offset := uint32(rank[beg])
	for i := cnt - 1; i >= 0; i-- {
		column.Columnar[columnar[offset+ix[i]]] = nil
		ix[i] = columnar[offset+ix[i]]
	}

	extractContent(fmi, column.Columnar)
	return column
}

// rank [s, e)
func selectInt64(rank []uint, header *ssd.HeaderISA, columnar []int64, entity []uint16) *internal.ColumnInt64 {
	sz := rank[end] - rank[beg]
	if sz < uint(len(entity)) {
		sz = uint(len(entity))
	}
	flat := make([]uint32, 2*sz)
	ix, iy := flat[:sz], flat[sz:2*sz]
	// IntersectSortedUint16(x, y, ix, iy) returns intersection size
	cnt := ext.IntersectSortedUint16(header.Entity[rank[beg]:rank[end]], entity, ix, iy)

	return &internal.ColumnInt64{
		Columnar: columnar[rank[beg]:rank[end]],
		Rows:     iy[:cnt],
		Cols:     ix[:cnt],
	}
}

// rank [s, e)
func selectBool(rank []uint, header *ssd.HeaderISA, columnar []bool, entity []uint16) *internal.ColumnBool {
	sz := rank[end] - rank[beg]
	if sz < uint(len(entity)) {
		sz = uint(len(entity))
	}
	flat := make([]uint32, 2*sz)
	ix, iy := flat[:sz], flat[sz:2*sz]
	// IntersectSortedUint16(x, y, ix, iy) returns intersection size
	cnt := ext.IntersectSortedUint16(header.Entity[rank[beg]:rank[end]], entity, ix, iy)

	return &internal.ColumnBool{
		Columnar: columnar[rank[beg]:rank[end]],
		Rows:     iy[:cnt],
		Cols:     ix[:cnt],
	}
}

// rank [s, e)
func selectNull(rank []uint, header *ssd.HeaderISA, entity []uint16) *internal.ColumnNull {
	sz := rank[end] - rank[beg]
	if sz < uint(len(entity)) {
		sz = uint(len(entity))
	}
	flat := make([]uint32, 2*sz)
	ix, iy := flat[:sz], flat[sz:2*sz]
	// IntersectSortedUint16(x, y, ix, iy) returns intersection size
	cnt := ext.IntersectSortedUint16(header.Entity[rank[beg]:rank[end]], entity, ix, iy)

	return &internal.ColumnNull{
		Rows: iy[:cnt],
		Cols: ix[:cnt],
	}
}

// rank [s, e)
func selectFloat64(rank []uint, header *ssd.HeaderISA, columnar []float64, entity []uint16) *internal.ColumnFloat64 {
	sz := rank[end] - rank[beg]
	if sz < uint(len(entity)) {
		sz = uint(len(entity))
	}
	flat := make([]uint32, 2*sz)
	ix, iy := flat[:sz], flat[sz:2*sz]
	// IntersectSortedUint16(x, y, ix, iy) returns intersection size
	cnt := ext.IntersectSortedUint16(header.Entity[rank[beg]:rank[end]], entity, ix, iy)

	return &internal.ColumnFloat64{
		Columnar: columnar[rank[beg]:rank[end]],
		Rows:     iy[:cnt],
		Cols:     ix[:cnt],
	}
}

func selectJson(index *ssd.IndexedBlock, key []byte, extract bool, entity []uint16) *internal.ColumnJson {
	_, end, ok := index.Entity.GetBound(ssd.SOH)
	if !ok {
		end = 0
	}

	rows := make([]uint32, len(entity))
	cols := make([][]byte, len(entity))
	m := len(entity)

	var (
		bound   []uint
		padding = uint(1)
	)

	for j := len(entity) - 1; j >= 0; j-- {
		if int(entity[j]) < len(index.EntityID)-1 {
			// including EOO
			end = uint(index.EntityID[entity[j]+1]) - 2
		}

		if entity[j] == 0 {
			if len(key) == 0 {
				padding, bound = 0, []uint{0, end + 1}
			} else {
				bound = findKeyBoundFromSOH(index.Entity, end+1, key)
			}
		} else {
			bound = findKeyBound(index.Entity, []uint{uint(index.EntityID[entity[j]]) - 1, end}, key)
		}
		if bound != nil {
			m--
			rows[m] = uint32(j)
			if extract {
				cols[m] = extractJson(index, key, bound, padding)
			} else {
				cols[m] = []byte{}
			}
		}
	}

	return &internal.ColumnJson{
		Rows: rows[m:],
		Cols: cols[m:],
	}
}
