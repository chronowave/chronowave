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

package codec

import (
	"bytes"
	"math"
	"strconv"

	"github.com/chronowave/chronowave/ssd"
)

func MarshalResultSet(rs *ssd.ResultSet, limit uint32) []byte {
	if rs == nil || len(rs.RowId) == 0 {
		return []byte("[]")
	} else if limit == 0 || limit > uint32(len(rs.RowId)) {
		limit = uint32(len(rs.RowId))
	}

	order := rs.Order
	if len(order) == 0 {
		order = make([]int, limit)
		for i := range rs.RowId[:limit] {
			order[i] = i
		}
	}

	w := bytes.NewBuffer(make([]byte, 0, 2048))
	w.WriteByte('[')

	var (
		comma = []byte{','}
		row   []byte
	)
	for _, i := range order {
		w.Write(row)
		w.WriteByte('{')
		var col []byte
		for j, c := range rs.Column {
			w.Write(col)
			w.WriteByte('"')
			w.WriteString(c.Name)
			w.WriteByte('"')
			w.WriteByte(':')

			if len(c.RowIdx) == 0 || c.RowIdx[i] == 0 {
				w.WriteString("null")
			} else {
				switch rs.ColumnType[j] {
				case 0:
					v := math.Round(rs.Aggregate[c.Value[i]].DoneFloat(.5)*10000) / 10000
					if math.IsNaN(v) {
						w.WriteString("null")
					} else {
						w.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
					}
				case ssd.TEXT:
					w.WriteByte('"')
					w.Write(rs.Text[c.Value[i]])
					w.WriteByte('"')
				case ssd.FLT64:
					w.WriteString(strconv.FormatFloat(math.Float64frombits(c.Value[i]), 'f', -1, 64))
				case ssd.INT64:
					w.WriteString(strconv.FormatInt(int64(c.Value[i]), 10))
				case ssd.JSON:
					w.Write(rs.Json[c.Value[i]])
				case ssd.BOOL:
					if c.Value[i] == 0 {
						w.WriteString("false")
					} else {
						w.WriteString("true")
					}
				}
			}
			col = comma
		}
		w.WriteByte('}')
		row = comma
	}

	w.WriteByte(']')
	return w.Bytes()
}
