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
	"encoding/json"
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssd"
)

func TestToJSON(t *testing.T) {
	type args struct {
		rs    *ssd.ResultSet
		limit uint32
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"empty", args{
			rs: &ssd.ResultSet{},
		}, []byte("[]")},
		{"text", args{
			rs: &ssd.ResultSet{
				RowId: []uint64{2, 4, 5},
				Column: []ssd.Column{
					{
						Name:   "a",
						RowIdx: []byte{1, 0, 1},
						Value:  []uint64{1, 3, 0},
					},
					{
						Name:   "b",
						RowIdx: []byte{0, 1, 1},
						Value:  []uint64{12, 4, 5},
					},
				},
				ColumnType: []byte{ssd.TEXT, ssd.INT64},
				Text:       [][]byte{[]byte("text 1"), []byte("text 2")},
			},
		}, []byte(`[{"a":"text 2","b":null},{"a":null,"b":4},{"a":"text 1","b":5}]`)},
		{"limit", args{
			rs: &ssd.ResultSet{
				RowId: []uint64{2, 4, 5},
				Column: []ssd.Column{
					{
						Name:   "a",
						RowIdx: []byte{1, 0, 1},
						Value:  []uint64{1, 3, 0},
					},
					{
						Name:   "b",
						RowIdx: []byte{0, 1, 1},
						Value:  []uint64{12, 4, 5},
					},
				},
				ColumnType: []byte{ssd.TEXT, ssd.INT64},
				Text:       [][]byte{[]byte("text 1"), []byte("text 2")},
			},
			limit: 1,
		}, []byte(`[{"a":"text 2","b":null}]`)},
		{"bool", args{
			rs: &ssd.ResultSet{
				RowId: []uint64{2, 4, 5},
				Column: []ssd.Column{
					{
						Name:   "a",
						RowIdx: []byte{1, 0, 1},
						Value:  []uint64{4609434218613702656, 0, 4609434218613702656},
					},
					{
						Name:   "b",
						RowIdx: []byte{0, 1, 1},
						Value:  []uint64{1, 0, 1},
					},
				},
				ColumnType: []byte{ssd.FLT64, ssd.BOOL},
			},
		}, []byte(`[{"a":1.5,"b":null},{"a":null,"b":false},{"a":1.5,"b":true}]`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MarshalResultSet(tt.args.rs, tt.args.limit)

			var data []interface{}
			err := json.Unmarshal(got, &data)
			if err != nil {
				t.Errorf("generated invalidate json err = %v, json = %v\n", err, string(got))
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalResultSet() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
