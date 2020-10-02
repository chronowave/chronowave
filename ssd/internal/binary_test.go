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

package internal

import (
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssd"
)

func TestColumnInt64Bytes(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		from := &ColumnInt64{
			Columnar: []int64{40, 20, 10, 30},
			Rows:     []uint32{3, 6, 8, 10},
			Cols:     []uint32{0, 1, 2, 3},
		}
		buf := from.Bytes()
		copied := make([]byte, len(buf))
		copy(copied, buf)
		to := &ColumnInt64{}
		to.FromBytes(copied)

		if !reflect.DeepEqual(to.Columnar, from.Columnar) ||
			!reflect.DeepEqual(to.Rows, from.Rows) {
			t.Errorf("Bytes() = %v, want %v", to, from)
		}
	})
}

func TestColumnTextBytes(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		from := &ColumnText{
			Columnar: map[uint32][]byte{
				12: []byte("this is test"),
				16: []byte("this is test 2"),
			},
			Rows: []uint32{3, 6, 8, 10},
			Cols: []uint32{0, 12, 12, 16},
		}
		buf := from.Bytes()
		copied := make([]byte, len(buf))
		copy(copied, buf)
		to := &ColumnText{}
		to.FromBytes(copied)

		if !reflect.DeepEqual(to.Columnar, from.Columnar) ||
			!reflect.DeepEqual(to.Rows, from.Rows) {
			t.Errorf("Bytes() = %v, want %v", to, from)
		}
	})
}

func TestResultSet(t *testing.T) {
	type args struct {
		rs *ssd.ResultSet
	}
	tests := []struct {
		name string
		args args
	}{
		{"empty", args{&ssd.ResultSet{}}},
		{"simple", args{&ssd.ResultSet{
			RowId:      []uint64{15, 68, 79},
			ColumnType: []byte{ssd.INT64, 0},
			Column: []ssd.Column{
				{
					RowIdx: []byte{1, 1, 1},
					Value:  []uint64{25, 28, 29},
				},
				{
					RowIdx: []byte{1, 1, 1},
					Value:  []uint64{0, 1, 2},
				},
			},
			Text: [][]byte{[]byte("test"), []byte("test2")},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serialized := SerializeResultSet(tt.args.rs)
			clone := make([]byte, len(serialized))
			copy(clone, serialized)
			got := DeserializeResultSet(clone)

			if !reflect.DeepEqual(got, tt.args.rs) {
				t.Errorf("SerializeResultSet() = %v, want %v", got, tt.args.rs)
			}
		})
	}
}
