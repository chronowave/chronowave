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
	"io/ioutil"
	"math"
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssql"
	"github.com/chronowave/chronowave/ssql/parser"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/aggregator"
	"github.com/chronowave/chronowave/ssd/codec"
	"github.com/chronowave/chronowave/ssd/index"
	"github.com/chronowave/chronowave/ssd/internal"
)

func TestMerge(t *testing.T) {
	f := 2.5
	type args struct {
		rss []*ssd.ResultSet
	}
	tests := []struct {
		name string
		args args
		want *ssd.ResultSet
	}{
		{"two", args{[]*ssd.ResultSet{
			{
				RowId:      []uint64{0},
				ColumnType: []byte{ssd.FLT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						Name:   "a",
						RowIdx: []byte{1},
						Value:  []uint64{math.Float64bits(f)},
					},
					{
						Name:   "b",
						RowIdx: []byte{1},
						Value:  []uint64{1},
					},
				},
				Text: [][]byte{[]byte(""), []byte("test1")},
			},
			{
				RowId:      []uint64{2},
				ColumnType: []byte{ssd.FLT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						Name:   "a",
						RowIdx: []byte{1},
						Value:  []uint64{math.Float64bits(f)},
					},
					{
						Name:   "b",
						RowIdx: []byte{1},
						Value:  []uint64{1},
					},
				},
				Text: [][]byte{[]byte(""), []byte("test2")},
			},
		}}, &ssd.ResultSet{
			RowId:      []uint64{0, 2},
			ColumnType: []byte{ssd.FLT64, ssd.TEXT},
			Column: []ssd.Column{
				{
					Name:   "a",
					RowIdx: []byte{1, 1},
					Value:  []uint64{math.Float64bits(f), math.Float64bits(f)},
				},
				{
					Name:   "b",
					RowIdx: []byte{1, 1},
					Value:  []uint64{1, 3},
				},
			},
			Text:      [][]byte{[]byte(""), []byte("test1"), []byte(""), []byte("test2")},
			Json:      [][]byte{},
			Aggregate: []aggregator.Aggregator{},
			Order:     []int{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.rss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeNull(t *testing.T) {
	f := 2.5
	type args struct {
		rss []*ssd.ResultSet
	}
	tests := []struct {
		name string
		args args
		want *ssd.ResultSet
	}{
		{"two", args{[]*ssd.ResultSet{
			{
				RowId:      []uint64{0},
				ColumnType: []byte{ssd.FLT64, ssd.TEXT},
				Column: []ssd.Column{
					{
						Name:   "a",
						RowIdx: []byte{1},
						Value:  []uint64{math.Float64bits(f)},
					},
					{
						Name:   "b",
						RowIdx: []byte{1},
						Value:  []uint64{1},
					},
				},
				Text: [][]byte{[]byte(""), []byte("test1")},
			},
			{
				RowId:      []uint64{2},
				ColumnType: []byte{ssd.NULL, ssd.TEXT},
				Column: []ssd.Column{
					{
						Name: "a",
					},
					{
						Name:   "b",
						RowIdx: []byte{1},
						Value:  []uint64{1},
					},
				},
				Text: [][]byte{[]byte(""), []byte("test2")},
			},
		}}, &ssd.ResultSet{
			RowId:      []uint64{0, 2},
			ColumnType: []byte{ssd.FLT64, ssd.TEXT},
			Column: []ssd.Column{
				{
					Name:   "a",
					RowIdx: []byte{1, 0},
					Value:  []uint64{math.Float64bits(f), 0},
				},
				{
					Name:   "b",
					RowIdx: []byte{1, 1},
					Value:  []uint64{1, 3},
				},
			},
			Text:      [][]byte{[]byte(""), []byte("test1"), []byte(""), []byte("test2")},
			Json:      [][]byte{},
			Aggregate: []aggregator.Aggregator{},
			Order:     []int{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.rss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeAggregate(t *testing.T) {
	type args struct {
		input []string
		stmt  string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"merge", args{
			input: []string{"test_1.json", "test_2.json"},
			stmt:  `find group-by($a), max($b), avg($c) where [$a /url] [$b /hit] [$c /duration]`,
		}, []byte(`[{"a":"/sample","b":50,"c":1.5},{"a":"/info","b":20,"c":1.5},{"a":"/checkout","b":8,"c":6.5},{"a":"/logout","b":50,"c":0.5},{"a":"/intro","b":200,"c":2.5}]`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt, e := parser.Parse(tt.args.stmt)
			if len(e) > 0 {
				t.Errorf("ssql Parse() error = %v", e)
			}
			rss := make([]*ssd.ResultSet, len(tt.args.input))
			for i := range rss {
				rss[i] = exec(t, uint64(i), tt.args.input[i], stmt)
			}

			merged := Merge(rss)

			if got := codec.MarshalResultSet(merged, 0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func exec(t *testing.T, iid uint64, file string, stmt *ssql.Statement) *ssd.ResultSet {
	data, err := ioutil.ReadFile("testdata/" + file)
	if err != nil {
		t.Errorf("error loading test file %v %v", file, err)
	}
	parsed, err := codec.ParseJson(data)
	if err != nil {
		t.Errorf("error to parse json %v", err)
	}

	indexed, err := index.Build(parsed)
	if err != nil {
		t.Errorf("index Build() error = %v", err)
	}

	node := map[string][]byte{}
	for _, expr := range stmt.Where {
		node[expr.GetTuple().Name], _ = indexed.Meta.GetCode(bytes.Split([]byte(expr.GetTuple().Path), []byte{'/'}))
	}

	columns := make([]internal.Column, len(stmt.Find))
	for i, f := range stmt.Find {
		columns[i].Key, columns[i].Ok = node[f.Name]
		columns[i].Name = f.Name
		columns[i].Group = f.Group
		columns[i].Func = f.Func
	}

	entity := make([]uint16, len(indexed.EntityID))
	for i := range entity {
		entity[i] = uint16(i)
	}

	columns = Select(indexed, columns, entity)
	return Consolidate(iid, entity, columns)
}
