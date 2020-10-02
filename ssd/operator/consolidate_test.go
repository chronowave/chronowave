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
	"math"
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssql"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/codec"
	"github.com/chronowave/chronowave/ssd/internal"
)

func BenchmarkConsolidate(b *testing.B) {
	entity := make([]uint16, benchSize)
	for i := range entity {
		entity[i] = uint16(i)
	}

	key, _ := benchIndex.Meta.GetCode(bytes.Split([]byte("a"), []byte{'/'}))
	columns := []internal.Column{
		{
			Ok:  true,
			Key: key,
		},
	}

	//first, last := testData[0].OneInt64, testData[len(testData)-1].OneInt64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		selected := Select(benchIndex, columns, entity)
		Consolidate(0, entity, selected)
	}
}

func TestConsolidateSelect(t *testing.T) {
	type args struct {
		entity  []uint16
		columns []internal.Column
	}
	tests := []struct {
		name string
		args args
		want *ssd.ResultSet
	}{
		{"select", args{
			entity: []uint16{3, 6, 8, 10},
			columns: []internal.Column{
				{
					Ok:   true,
					Type: ssd.TEXT,
					ColumnText: &internal.ColumnText{
						Columnar: map[uint32][]byte{
							0:  {},
							12: []byte("this is test"),
							16: []byte("this is test 2"),
						},
						Rows: []uint32{0, 2, 3},
						Cols: []uint32{0, 12, 16},
					},
				},
				{
					Ok:   true,
					Type: ssd.INT64,
					ColumnInt64: &internal.ColumnInt64{
						Columnar: []int64{40, 20, 10, 30},
						Rows:     []uint32{0, 1, 2, 3},
						Cols:     []uint32{0, 1, 2, 3},
					},
				},
				{
					Ok:   true,
					Type: ssd.FLT64,
					ColumnFloat64: &internal.ColumnFloat64{
						Columnar: []float64{9.2, 8.1, 6.5, 7.2},
						Rows:     []uint32{0, 1, 3},
						Cols:     []uint32{0, 1, 3},
					},
				},
			},
		},
			&ssd.ResultSet{
				RowId:      []uint64{3, 6, 8, 10},
				ColumnType: []byte{ssd.TEXT, ssd.INT64, ssd.FLT64},
				Column: []ssd.Column{
					{
						RowIdx: []byte{1, 0, 1, 1},
						Value:  []uint64{0, 0, 1, 2},
					},
					{
						RowIdx: []byte{1, 1, 1, 1},
						Value:  []uint64{40, 20, 10, 30},
					},
					{
						RowIdx: []byte{1, 1, 0, 1},
						Value:  []uint64{math.Float64bits(9.2), math.Float64bits(8.1), 0, math.Float64bits(7.2)},
					},
				},
				// TODO: sorting map key before assert
				Text: [][]byte{{}, []byte("this is test"), []byte("this is test 2")},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Consolidate(0, tt.args.entity, tt.args.columns)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("group() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConsolidateGroup(t *testing.T) {
	type args struct {
		entity  []uint16
		columns []internal.Column
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"group", args{
			entity: []uint16{3, 6, 8, 10},
			columns: []internal.Column{
				{
					Ok:    true,
					Name:  "t",
					Type:  ssd.TEXT,
					Group: true,
					ColumnText: &internal.ColumnText{
						Columnar: map[uint32][]byte{
							0:  {},
							12: []byte("test 1"),
							16: []byte("test 2"),
						},
						Rows: []uint32{0, 1, 2, 3},
						Cols: []uint32{0, 12, 16, 12},
					},
				},
				{
					Ok:    true,
					Name:  "int1",
					Type:  ssd.INT64,
					Group: true,
					ColumnInt64: &internal.ColumnInt64{
						Columnar: []int64{40, 20, 10, 30},
						Rows:     []uint32{0, 1, 2, 3},
						Cols:     []uint32{0, 1, 2, 1},
					},
				},
				{
					Ok:   true,
					Name: "int2",
					Type: ssd.INT64,
					Func: &ssql.Function{Name: ssql.Function_AVG},
					ColumnInt64: &internal.ColumnInt64{
						Columnar: []int64{90, 80, 60, 70},
						Rows:     []uint32{0, 1, 3},
						Cols:     []uint32{0, 1, 3},
					},
				},
			},
		}, []byte(`[{"t":"","int1":40,"int2":90},{"t":"test 1","int1":20,"int2":75},{"t":"test 2","int1":10,"int2":null}]`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Consolidate(0, tt.args.entity, tt.args.columns)
			json := codec.MarshalResultSet(got, 0)
			if !reflect.DeepEqual(json, tt.want) {
				t.Errorf("group() = %v, want %v", string(json), tt.want)
			}
		})
	}
}

func TestConsolidateNull(t *testing.T) {
	type args struct {
		entity  []uint16
		columns []internal.Column
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"no path", args{
			entity: []uint16{3, 6, 8, 10},
			columns: []internal.Column{
				{
					Ok:    true,
					Name:  "t",
					Type:  ssd.TEXT,
					Group: true,
					ColumnText: &internal.ColumnText{
						Columnar: map[uint32][]byte{
							0:  {},
							12: []byte("test 1"),
							16: []byte("test 2"),
						},
						Rows: []uint32{0, 1, 2, 3},
						Cols: []uint32{0, 12, 16, 12},
					},
				},
				{
					Ok:    true,
					Name:  "int1",
					Type:  ssd.INT64,
					Group: true,
					ColumnInt64: &internal.ColumnInt64{
						Columnar: []int64{40, 20, 10, 30},
						Rows:     []uint32{0, 1, 2, 3},
						Cols:     []uint32{0, 1, 2, 1},
					},
				},
				{
					Ok:   false,
					Name: "int2",
					Type: ssd.NULL,
				},
			},
		}, []byte(`[{"t":"","int1":40,"int2":null},{"t":"test 1","int1":20,"int2":null},{"t":"test 2","int1":10,"int2":null}]`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Consolidate(0, tt.args.entity, tt.args.columns)
			json := codec.MarshalResultSet(got, 0)
			if !reflect.DeepEqual(json, tt.want) {
				t.Errorf("group() = %v, want %v", string(json), tt.want)
			}
		})
	}
}
