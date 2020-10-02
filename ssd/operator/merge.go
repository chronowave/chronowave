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
	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/aggregator"
)

const (
	rii = iota
	cti
	txi
	jsi
	agi
	odi
)

func Merge(rss []*ssd.ResultSet) *ssd.ResultSet {
	var (
		merged = &ssd.ResultSet{}
		plain  = true
		sz     = [6]int{}
	)
	j := 0
	for _, rs := range rss {
		if len(rs.RowId) == 0 {
			continue
		}
		rss[j] = rs
		j++

		sz[rii] += len(rs.RowId)
		sz[cti] = len(rs.ColumnType)
		sz[txi] += len(rs.Text)
		sz[jsi] += len(rs.Json)
		sz[agi] += len(rs.Aggregate)
		sz[odi] += len(rs.Order)
	}
	merged.RowId = make([]uint64, sz[rii])
	merged.ColumnType = make([]byte, sz[cti])
	if sz[cti] > 0 {
		copy(merged.ColumnType, rss[j-1].ColumnType)
	}
	merged.Column = make([]ssd.Column, sz[cti])
	bufs, uint64s := make([]byte, sz[rii]*sz[cti]), make([]uint64, sz[rii]*sz[cti])
	for i := range merged.Column {
		merged.Column[i].RowIdx, bufs = bufs[:sz[rii]], bufs[sz[rii]:]
		merged.Column[i].Value, uint64s = uint64s[:sz[rii]], uint64s[sz[rii]:]
	}
	merged.Text = make([][]byte, sz[txi])
	merged.Json = make([][]byte, sz[jsi])
	merged.Aggregate = make([]aggregator.Aggregator, sz[agi])
	merged.Order = make([]int, sz[odi])

	for i := 0; i < len(sz); i++ {
		sz[i] = 0
	}

	rss = rss[:j]

	for _, rs := range rss {
		for i, c := range rs.ColumnType {
			switch c {
			case 0:
				plain = false
				// aggregate
				if sz[agi] > 0 {
					for j := range rs.Column[i].Value {
						rs.Column[i].Value[j] += uint64(sz[agi])
					}
				}
			case ssd.TEXT:
				// text
				if sz[txi] > 0 {
					for j := range rs.Column[i].Value {
						rs.Column[i].Value[j] += uint64(sz[txi])
					}
				}
			case ssd.JSON:
				// json
				if sz[jsi] > 0 {
					for j := range rs.Column[i].Value {
						rs.Column[i].Value[j] += uint64(sz[jsi])
					}
				}
			}

			copy(merged.Column[i].RowIdx[sz[rii]:], rs.Column[i].RowIdx)
			copy(merged.Column[i].Value[sz[rii]:], rs.Column[i].Value)
			merged.Column[i].Name = rs.Column[i].Name
			if merged.ColumnType[i] == ssd.NULL {
				merged.ColumnType[i] = c
			}
		}

		sz[rii] += copy(merged.RowId[sz[rii]:], rs.RowId)
		sz[txi] += copy(merged.Text[sz[txi]:], rs.Text)
		sz[jsi] += copy(merged.Json[sz[jsi]:], rs.Json)
		sz[agi] += copy(merged.Aggregate[sz[agi]:], rs.Aggregate)
		sz[odi] += copy(merged.Order[sz[odi]:], rs.Order)
	}

	if plain {
		return merged
	}

	var (
		rows = make(map[uint64]int, len(merged.RowId))
		cnt  uint64
	)

	for i, rid := range merged.RowId {
		if idx, ok := rows[rid]; ok {
			for j, ct := range merged.ColumnType {
				if ct == 0 && merged.Column[j].RowIdx[i] > 0 {
					// aggregate
					if merged.Column[j].RowIdx[idx] == 0 {
						merged.Column[j].RowIdx[idx] = 1
						merged.Aggregate[cnt] = merged.Aggregate[merged.Column[j].Value[i]]
						merged.Column[j].Value[idx] = cnt
						cnt++
					} else {
						y, x := merged.Column[j].Value[i], merged.Column[j].Value[idx]
						merged.Aggregate[x].Merge(merged.Aggregate[y])
					}
				}
			}
		} else {
			idx = len(rows)
			rows[rid] = idx
			if i > idx {
				merged.RowId[idx] = rid
				for j, ct := range merged.ColumnType {
					merged.Column[j].RowIdx[idx] = merged.Column[j].RowIdx[i]
					if ct == 0 {
						// aggregate
						merged.Aggregate[cnt] = merged.Aggregate[merged.Column[j].Value[i]]
						merged.Column[j].Value[idx] = cnt
						cnt++
					} else {
						merged.Column[j].Value[idx] = merged.Column[j].Value[i]
					}
				}
			} else {
				for j, ct := range merged.ColumnType {
					if ct == 0 {
						// aggregate
						v := merged.Column[j].Value[i]
						if cnt < v {
							merged.Aggregate[cnt] = merged.Aggregate[v]
							merged.Column[j].Value[idx] = cnt
						}

						cnt++
					}
				}
			}
		}
	}

	merged.Aggregate = merged.Aggregate[:cnt]
	cnt = uint64(len(rows))
	merged.RowId = merged.RowId[:cnt]
	for i := range merged.Column {
		merged.Column[i].RowIdx = merged.Column[i].RowIdx[:cnt]
		merged.Column[i].Value = merged.Column[i].Value[:cnt]
	}

	return merged
}
