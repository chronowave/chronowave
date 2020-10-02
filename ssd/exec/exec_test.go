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

package exec

import (
	"math"
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssql/parser"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/codec"
)

func TestExec(t *testing.T) {
	v, f := int64(-1), 0.5
	type args struct {
		json string
		stmt string
	}
	tests := []struct {
		name string
		args args
		want *ssd.ResultSet
	}{
		{"text", args{
			json: `{"a": "bb", "c": "dd"}`,
			stmt: `find $a where [$a /a contain("b")]`},
			&ssd.ResultSet{
				RowId:      []uint64{0},
				ColumnType: []byte{ssd.TEXT},
				Column: []ssd.Column{{
					Name:   "a",
					RowIdx: []byte{1},
					Value:  []uint64{0},
				}},
				Text: [][]byte{[]byte("bb")},
			},
		},
		{"int64", args{
			json: `{"a": -1, "c": "dd"}`,
			stmt: `find $a where [$a /a le(1)]`},
			&ssd.ResultSet{
				RowId:      []uint64{0},
				ColumnType: []byte{ssd.INT64},
				Column: []ssd.Column{{
					Name:   "a",
					RowIdx: []byte{1},
					Value:  []uint64{uint64(v)},
				}},
				Text: [][]byte{},
			},
		},
		{"float64", args{
			json: `{"a": 0.5, "c": "dd"}`,
			stmt: `find $a where [$a /a le(2.0)]`},
			&ssd.ResultSet{
				RowId:      []uint64{0},
				ColumnType: []byte{ssd.FLT64},
				Column: []ssd.Column{{
					Name:   "a",
					RowIdx: []byte{1},
					Value:  []uint64{math.Float64bits(f)},
				}},
				Text: [][]byte{},
			},
		},
		{"empty", args{
			json: `{"a": 0.5, "c": "dd"}`,
			stmt: `find $a where [$a /a eq(2.0)]`},
			&ssd.ResultSet{
				RowId: []uint64{},
			},
		},
		{"null", args{
			json: `{"a": 0.5, "c": "dd"} {"a": 0.5, "d": 2}`,
			stmt: `find $b where [$b /abc][/a eq(0.5)]`},
			&ssd.ResultSet{
				RowId:      []uint64{0, 1},
				ColumnType: []byte{ssd.NULL},
				Column: []ssd.Column{{
					Name: "b",
				}},
				Text: [][]byte{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := codec.ParseJson([]byte(tt.args.json))
			if err != nil {
				t.Errorf("error to parse json %v", err)
			}

			indexed, err := buildTestIndex(parsed)
			if err != nil {
				t.Errorf("index Build() error = %v", err)
			}

			stmt, e := parser.Parse(tt.args.stmt)
			if len(e) > 0 {
				t.Errorf("ssql Parse() error = %v", e)
			}

			if got := Exec(indexed, stmt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
