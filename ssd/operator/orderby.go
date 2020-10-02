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
	"bytes"
	"sort"

	"github.com/chronowave/chronowave/ssql"

	"github.com/chronowave/chronowave/ssd"
)

func OrderBy(rs *ssd.ResultSet, stmt *ssql.Statement) *ssd.ResultSet {
	if rs == nil || len(rs.RowId) == 0 {
		return rs
	}

	rows := make([]int, len(rs.RowId))
	for i := range rows {
		rows[i] = i
	}

	columns := make(map[string]int, len(stmt.Find))
	for i, f := range stmt.Find {
		columns[f.Name] = i
	}

	for i := len(stmt.OrderBy) - 1; i >= 0; i-- {
		orderby := stmt.OrderBy[i]
		idx, ok := columns[orderby.Name]
		if !ok {
			continue
		}
		v := getDoubleValue(stmt.Find[idx].Func)
		if len(rs.Column[idx].RowIdx) == 0 {
			continue
		}
		column, columnType := &rs.Column[idx], rs.ColumnType[idx]
		var less func(int, int) bool

		if orderby.Direction == ssql.OrderBy_ASC {
			// x < y, asc
			less = func(yi, xi int) bool {
				y, x := rows[yi], rows[xi]
				if column.RowIdx[x] == 0 {
					return false
				} else if column.RowIdx[y] == 0 {
					return true
				}

				switch columnType {
				case 0:
					return rs.Aggregate[column.Value[y]].DoneFloat(v) < rs.Aggregate[column.Value[x]].DoneFloat(v)
				case ssd.TEXT:
					// yv < xv
					return bytes.Compare(rs.Text[column.Value[y]], rs.Text[column.Value[x]]) < 0
				case ssd.INT64:
					return int64(column.Value[y]) < int64(column.Value[x])
				case ssd.NULL:
					return true
				}
				// data type doesn't match, could be null column
				return true
			}
		} else {
			// x > y, desc
			less = func(yi, xi int) bool {
				y, x := rows[yi], rows[xi]
				if column.RowIdx[x] == 0 {
					return true
				} else if column.RowIdx[y] == 0 {
					return false
				}

				switch columnType {
				case 0:
					return rs.Aggregate[column.Value[y]].DoneFloat(v) > rs.Aggregate[column.Value[x]].DoneFloat(v)
				case ssd.TEXT:
					// yv > xv
					return bytes.Compare(rs.Text[column.Value[y]], rs.Text[column.Value[x]]) > 0
				case ssd.INT64:
					return int64(column.Value[y]) > int64(column.Value[x])
				case ssd.NULL:
					return true
				}
				// data type doesn't match, could be null column
				return true
			}
		}

		sort.SliceStable(rows, less)
	}

	rs.Order = rows

	return rs
}

func getDoubleValue(f *ssql.Function) float64 {
	if f != nil {
		switch f.Name {
		case ssql.Function_PCTL:
			return f.GetDouble()
		}
	}
	return 0
}
