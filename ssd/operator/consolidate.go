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
	"encoding/binary"
	"math"
	"unsafe"

	"github.com/chronowave/chronowave/ssql"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/aggregator"
	"github.com/chronowave/chronowave/ssd/internal"

	"github.com/cespare/xxhash/v2"
)

func Consolidate(iid uint64, rows []uint16, columns []internal.Column) *ssd.ResultSet {
	nor, noc, iid := len(rows), len(columns), iid*ssd.MaxDoc

	var (
		grpIdx      map[uint64]int
		digest      *xxhash.Digest
		texts       = make([][]byte, 0, 128)
		textIdx     = make(map[uint32]int, 128)
		groups      = make([]int, 0, noc)
		aggregates  = make([]int, 0, noc)
		allocRows   = make([]byte, nor*noc)
		allocValues = make([]uint64, nor*noc)
	)

	rs := &ssd.ResultSet{
		RowId:      make([]uint64, nor),
		Column:     make([]ssd.Column, noc),
		ColumnType: make([]byte, noc),
	}

	for i := range columns {
		rs.Column[i].Name = columns[i].Name
		if columns[i].Group {
			if digest == nil {
				digest = xxhash.New()
				grpIdx = make(map[uint64]int, nor)
			}
			if columns[i].Ok {
				groups = append(groups, i)
			}
		}

		rs.ColumnType[i] = columns[i].Type
		if !columns[i].Ok {
			continue
		}

		if columns[i].Func != nil {
			// rs.ColumnType[i] == 0 for aggregate
			rs.ColumnType[i] = 0
			aggregates = append(aggregates, i)
		}

		rs.Column[i].RowIdx, allocRows = allocRows[:nor], allocRows[nor:]
		rs.Column[i].Value, allocValues = allocValues[:nor], allocValues[nor:]

		if columns[i].Type == ssd.TEXT {
			for k, v := range columns[i].ColumnText.Columnar {
				if _, ok := textIdx[k]; !ok {
					textIdx[k] = len(texts)
					texts = append(texts, v)
				}
			}
		}
	}

	rs.Text = texts

	if digest == nil && len(aggregates) == 0 {
		for e, rid := range rows {
			rs.RowId[e] = uint64(rid) + iid
		}
		for i := range columns {
			if !columns[i].Ok {
				continue
			}
			switch columns[i].Type {
			case ssd.TEXT:
				columnar := columns[i].ColumnText
				for j, r := range columnar.Rows {
					v := textIdx[columnar.Cols[j]]
					rs.Column[i].RowIdx[r] = 1
					rs.Column[i].Value[r] = uint64(v)
				}
			case ssd.FLT64:
				columnar := columns[i].ColumnFloat64
				for j, r := range columnar.Rows {
					v := columnar.Columnar[columnar.Cols[j]]
					rs.Column[i].RowIdx[r] = 1
					rs.Column[i].Value[r] = math.Float64bits(v)
				}
			case ssd.INT64:
				columnar := columns[i].ColumnInt64
				for j, r := range columnar.Rows {
					v := columnar.Columnar[columnar.Cols[j]]
					rs.Column[i].RowIdx[r] = 1
					rs.Column[i].Value[r] = uint64(v)
				}
			case ssd.BOOL:
				columnar := columns[i].ColumnBool
				for j, r := range columnar.Rows {
					v := columnar.Columnar[columnar.Cols[j]]
					rs.Column[i].RowIdx[r] = 1
					rs.Column[i].Value[r] = uint64(*(*byte)(unsafe.Pointer(&v)) & 1)
				}
			case ssd.JSON:
				columnar := columns[i].ColumnJson
				for j, r := range columnar.Rows {
					rs.Column[i].RowIdx[r] = 1
					rs.Column[i].Value[r] = uint64(len(rs.Json))
					rs.Json = append(rs.Json, columnar.Cols[j])
				}
			}
		}

		return rs
	}

	var (
		aidx = 0
		ridx = 0
		cidx = make([]int, len(columns))
		buf  = [8]byte{}
	)

	rs.Aggregate = make([]aggregator.Aggregator, nor*len(aggregates))
	if len(aggregates) > 0 && digest == nil {
		// aggregates without group by
		for _, f := range aggregates {
			// init functions
			switch columns[f].Func.Name {
			case ssql.Function_AVG:
				rs.Aggregate[aidx] = aggregator.NewAverage()
			case ssql.Function_SUM:
				rs.Aggregate[aidx] = aggregator.NewSum()
			case ssql.Function_MIN:
				rs.Aggregate[aidx] = aggregator.NewMin()
			case ssql.Function_MAX:
				rs.Aggregate[aidx] = aggregator.NewMax()
			case ssql.Function_COUNT:
				rs.Aggregate[aidx] = aggregator.NewCount()
			case ssql.Function_PCTL:
				rs.Aggregate[aidx] = aggregator.NewPercentile()
			}
			rs.Column[f].RowIdx[0] = 0
			rs.Column[f].Value[0] = uint64(aidx)
			aidx++
		}
	}

	for i, rid := range rows {
		rs.RowId[i], ridx = uint64(rid)+iid, i

		if digest != nil {
			digest.Reset()

			// calculate group by hash value
			for _, g := range groups {
				switch columns[g].Type {
				case ssd.TEXT:
					j, columnar := cidx[g], columns[g].ColumnText
					if columnar.Rows[j] == uint32(i) {
						v := textIdx[columnar.Cols[j]]
						digest.Write(rs.Text[v])
						cidx[g]++

						rs.Column[g].RowIdx[i] = 1
						rs.Column[g].Value[i] = uint64(v)
					} else {
						rs.Column[g].RowIdx[i] = 0
					}
				case ssd.FLT64:
					j, columnar := cidx[g], columns[g].ColumnFloat64
					if columnar.Rows[j] == uint32(i) {
						v := math.Float64bits(columnar.Columnar[columnar.Cols[j]])
						binary.LittleEndian.PutUint64(buf[:8], v)
						digest.Write(buf[:8])
						cidx[g]++

						rs.Column[g].RowIdx[i] = 1
						rs.Column[g].Value[i] = v
					} else {
						rs.Column[g].RowIdx[i] = 0
					}
				case ssd.INT64:
					j, columnar := cidx[g], columns[g].ColumnInt64
					if columnar.Rows[j] == uint32(i) {
						v := columnar.Columnar[columnar.Cols[j]]
						binary.LittleEndian.PutUint64(buf[:8], uint64(v))
						digest.Write(buf[:8])
						cidx[g]++

						rs.Column[g].RowIdx[i] = 1
						rs.Column[g].Value[i] = uint64(v)
					} else {
						rs.Column[g].RowIdx[i] = 0
					}
				case ssd.BOOL:
					j, columnar := cidx[g], columns[g].ColumnBool
					if columnar.Rows[j] == uint32(i) {
						v := columnar.Columnar[columnar.Cols[j]]
						buf[0] = *(*byte)(unsafe.Pointer(&v)) & 1
						digest.Write(buf[:1])
						cidx[g]++

						rs.Column[g].RowIdx[i] = 1
						rs.Column[g].Value[i] = uint64(buf[0])
					} else {
						rs.Column[g].RowIdx[i] = 0
					}
				case ssd.JSON:
					j, columnar := cidx[g], columns[g].ColumnJson
					if columnar.Rows[j] == uint32(i) {
						digest.Write(columnar.Cols[j])
						cidx[g]++

						rs.Column[g].RowIdx[i] = 1
						rs.Column[g].Value[i] = uint64(len(rs.Json))
						rs.Json = append(rs.Json, columnar.Cols[j])
					} else {
						rs.Column[g].RowIdx[i] = 0
					}
				}
			}

			rid := digest.Sum64()
			if idx, ok := grpIdx[rid]; ok {
				// use existing index
				ridx = idx
			} else {
				ridx = len(grpIdx)
				grpIdx[rid] = ridx
				rs.RowId[ridx] = rid

				for _, g := range groups {
					rs.Column[g].RowIdx[ridx] = rs.Column[g].RowIdx[i]
					rs.Column[g].Value[ridx] = rs.Column[g].Value[i]
				}

				// init functions
				for _, f := range aggregates {
					switch columns[f].Func.Name {
					case ssql.Function_AVG:
						rs.Aggregate[aidx] = aggregator.NewAverage()
					case ssql.Function_SUM:
						rs.Aggregate[aidx] = aggregator.NewSum()
					case ssql.Function_MIN:
						rs.Aggregate[aidx] = aggregator.NewMin()
					case ssql.Function_MAX:
						rs.Aggregate[aidx] = aggregator.NewMax()
					case ssql.Function_COUNT:
						rs.Aggregate[aidx] = aggregator.NewCount()
					case ssql.Function_PCTL:
						rs.Aggregate[aidx] = aggregator.NewPercentile()
					}
					rs.Column[f].RowIdx[ridx] = 0
					rs.Column[f].Value[ridx] = uint64(aidx)
					aidx++
				}
			}
		} else {
			ridx = 0
		}

		for _, f := range aggregates {
			column := &columns[f]
			switch column.Type {
			case ssd.TEXT:
				j, columnar := cidx[f], column.ColumnText
				if columnar.Rows[j] == uint32(i) {
					v := uint32(textIdx[columnar.Cols[j]])
					rs.Aggregate[rs.Column[f].Value[ridx]].StepInt64(int64(v))
					rs.Column[f].RowIdx[ridx] = 1
					cidx[f]++
				}
			case ssd.FLT64:
				j, columnar := cidx[f], column.ColumnFloat64
				if columnar.Rows[j] == uint32(i) {
					v := columnar.Columnar[columnar.Cols[j]]
					rs.Aggregate[rs.Column[f].Value[ridx]].StepFloat64(v)
					rs.Column[f].RowIdx[ridx] = 1
					cidx[f]++
				}
			case ssd.INT64:
				j, columnar := cidx[f], column.ColumnInt64
				if columnar.Rows[j] == uint32(i) {
					v := columnar.Columnar[columnar.Cols[j]]
					rs.Aggregate[rs.Column[f].Value[ridx]].StepInt64(v)
					rs.Column[f].RowIdx[ridx] = 1
					cidx[f]++
				}
			case ssd.BOOL:
				j, columnar := cidx[f], column.ColumnBool
				if columnar.Rows[j] == uint32(i) {
					rs.Aggregate[rs.Column[f].Value[ridx]].StepInt64(0)
					rs.Column[f].RowIdx[ridx] = 1
					cidx[f]++
				}
			case ssd.NULL:
				j, columnar := cidx[f], column.ColumnNull
				if columnar.Rows[j] == uint32(i) {
					rs.Aggregate[rs.Column[f].Value[ridx]].StepInt64(0)
					rs.Column[f].RowIdx[ridx] = 1
					cidx[f]++
				}
			case ssd.JSON:
				j, columnar := cidx[f], column.ColumnJson
				if columnar.Rows[j] == uint32(i) {
					rs.Aggregate[rs.Column[f].Value[ridx]].StepInt64(0)
					rs.Column[f].RowIdx[ridx] = 1
					cidx[f]++
				}
			}
		}
	}

	rs.Aggregate = rs.Aggregate[:aidx]
	cnt := 1
	if len(grpIdx) > 0 {
		cnt = len(grpIdx)
	}
	rs.RowId = rs.RowId[:cnt]
	return rs
}
