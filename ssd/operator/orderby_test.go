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
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssql/parser"

	"github.com/chronowave/chronowave/ssd"
)

func TestOrderByText(t *testing.T) {
	type args struct {
		rs   *ssd.ResultSet
		stmt string
	}
	tests := []struct {
		name string
		args args
		want *ssd.ResultSet
	}{
		{"desc", args{
			rs: &ssd.ResultSet{
				RowId:      []uint64{12, 28, 32, 48},
				ColumnType: []byte{ssd.INT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						Name:   "a",
						RowIdx: []byte{1, 1, 1, 1},
						Value:  []uint64{7, 4, 2, 3},
					},
					{
						Name:   "b",
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{1, 0, 0, 2},
					},
				},
				Text: [][]byte{[]byte("order 1"), []byte("order 2"), []byte("order 3")},
			},
			stmt: "find $a, $b WHERE [$a /][$b /] order-by $b desc",
		},
			&ssd.ResultSet{
				Order: []int{3, 0, 2, 1},
			},
		},
		{"asc", args{
			rs: &ssd.ResultSet{
				RowId:      []uint64{12, 28, 32, 48},
				ColumnType: []byte{ssd.INT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						RowIdx: []byte{1, 1, 1, 1},
						Value:  []uint64{7, 4, 2, 3},
					},
					{
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{1, 0, 0, 2},
					},
				},
				Text: [][]byte{[]byte("order 1"), []byte("order 2"), []byte("order 3")},
			},
			stmt: "find $a, $b WHERE [$a /][$b /] order-by $b asc",
		},
			&ssd.ResultSet{
				Order: []int{1, 2, 0, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt, err := parser.Parse(tt.args.stmt)
			if err != nil {
				t.Errorf("syntax error %v", tt.args.stmt)
			}
			if got := OrderBy(tt.args.rs, stmt); !reflect.DeepEqual(got.Order, tt.want.Order) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderByINT64(t *testing.T) {
	type args struct {
		rs   *ssd.ResultSet
		stmt string
	}
	tests := []struct {
		name string
		args args
		want *ssd.ResultSet
	}{
		{"desc", args{
			rs: &ssd.ResultSet{
				RowId:      []uint64{12, 28, 32, 48},
				ColumnType: []byte{ssd.INT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						Name:   "a",
						RowIdx: []byte{1, 1, 1, 1},
						Value:  []uint64{4, 7, 2, 3},
					},
					{
						Name:   "b",
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{1, 0, 0, 2},
					},
				},
				Text: [][]byte{[]byte("order 1"), []byte("order 2"), []byte("order 3")},
			},
			stmt: "find $a, $b WHERE [$a /][$b /] order-by $a desc", /*[]orderby{
				{
					col: 0,
					v:   0,
					dir: ssql.OrderBy_DESC,
				},
			},
			*/
		},
			&ssd.ResultSet{
				Order: []int{1, 0, 3, 2},
			},
		},
		{"asc", args{
			rs: &ssd.ResultSet{
				RowId:      []uint64{12, 28, 32, 48},
				ColumnType: []byte{ssd.INT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						RowIdx: []byte{1, 1, 1, 1},
						Value:  []uint64{4, 7, 2, 3},
					},
					{
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{1, 0, 0, 2},
					},
				},
				Text: [][]byte{[]byte("order 1"), []byte("order 2"), []byte("order 3")},
			},
			stmt: "find $a, $b WHERE [$a /][$b /] order-by $a asc",
		},
			&ssd.ResultSet{
				Order: []int{2, 3, 0, 1},
			},
		},
		{"desc null", args{
			rs: &ssd.ResultSet{
				RowId:      []uint64{12, 28, 32, 48},
				ColumnType: []byte{ssd.INT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						RowIdx: []byte{0, 1, 1, 1},
						Value:  []uint64{4, 7, 2, 2},
					},
					{
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{1, 0, 0, 2},
					},
				},
				Text: [][]byte{[]byte("order 1"), []byte("order 2"), []byte("order 3")},
			},
			stmt: "find $a, $b WHERE [$a /][$b /] order-by $a desc",
		},
			&ssd.ResultSet{
				Order: []int{1, 2, 3, 0},
			},
		},
		{"asc null", args{
			rs: &ssd.ResultSet{
				RowId:      []uint64{12, 28, 32, 48},
				ColumnType: []byte{ssd.INT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{4, 7, 2, 2},
					},
					{
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{1, 0, 0, 2},
					},
				},
				Text: [][]byte{[]byte("order 1"), []byte("order 2"), []byte("order 3")},
			},
			stmt: "find $a, $b WHERE [$a /][$b /] order-by $a asc",
		},
			&ssd.ResultSet{
				Order: []int{1, 2, 3, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt, err := parser.Parse(tt.args.stmt)
			if err != nil {
				t.Errorf("error parsing %v", err)
			}
			if got := OrderBy(tt.args.rs, stmt); !reflect.DeepEqual(got.Order, tt.want.Order) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderBy(t *testing.T) {
	type args struct {
		rs   *ssd.ResultSet
		stmt string
	}
	tests := []struct {
		name string
		args args
		want *ssd.ResultSet
	}{
		{"desc desc", args{
			rs: &ssd.ResultSet{
				RowId:      []uint64{12, 28, 32, 48},
				ColumnType: []byte{ssd.INT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						Name:   "a",
						RowIdx: []byte{1, 1, 1, 1},
						Value:  []uint64{2, 7, 3, 4},
					},
					{
						Name:   "b",
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{1, 0, 2, 2},
					},
				},
				Text: [][]byte{[]byte("order 1"), []byte("order 2"), []byte("order 3")},
			},
			stmt: "find $a, $b WHERE [$a /][$b /] order-by $b desc, $a desc",
		},
			&ssd.ResultSet{
				Order: []int{3, 2, 0, 1},
			},
		},
		{"desc asc", args{
			rs: &ssd.ResultSet{
				RowId:      []uint64{12, 28, 32, 48},
				ColumnType: []byte{ssd.INT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						Name:   "a",
						RowIdx: []byte{1, 1, 1, 1},
						Value:  []uint64{2, 7, 3, 4},
					},
					{
						Name:   "b",
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{1, 0, 2, 2},
					},
				},
				Text: [][]byte{[]byte("order 1"), []byte("order 2"), []byte("order 3")},
			},
			stmt: "find $a, $b WHERE [$a /][$b /] order-by $b desc, $a limit 2",
		},
			&ssd.ResultSet{
				Order: []int{2, 3, 0, 1},
			},
		},
		{"asc asc", args{
			rs: &ssd.ResultSet{
				RowId:      []uint64{12, 28, 32, 48},
				ColumnType: []byte{ssd.INT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						RowIdx: []byte{1, 1, 1, 1},
						Value:  []uint64{7, 4, 2, 3},
					},
					{
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{1, 0, 0, 2},
					},
				},
				Text: [][]byte{[]byte("order 1"), []byte("order 2"), []byte("order 3")},
			},
			stmt: "find $a, $b WHERE [$a /][$b /] order-by $b asc, $a",
		},
			&ssd.ResultSet{
				Order: []int{1, 2, 0, 3},
			},
		},
		{"asc desc", args{
			rs: &ssd.ResultSet{
				RowId:      []uint64{12, 28, 32, 48},
				ColumnType: []byte{ssd.INT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						RowIdx: []byte{1, 1, 1, 1},
						Value:  []uint64{7, 2, 4, 3},
					},
					{
						RowIdx: []byte{1, 0, 0, 1},
						Value:  []uint64{1, 0, 0, 2},
					},
				},
				Text: [][]byte{[]byte("order 1"), []byte("order 2"), []byte("order 3")},
			},
			stmt: "find $a, $b WHERE [$a /][$b /] order-by $b, $a desc",
		},
			&ssd.ResultSet{
				Order: []int{2, 1, 0, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt, err := parser.Parse(tt.args.stmt)
			if err != nil {
				t.Errorf("error parsing %v", err)
			}
			if got := OrderBy(tt.args.rs, stmt); !reflect.DeepEqual(got.Order, tt.want.Order) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
