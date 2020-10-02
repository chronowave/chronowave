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
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssd/codec"
	"github.com/chronowave/chronowave/ssd/index"
	"github.com/chronowave/chronowave/ssd/internal"
)

func TestSelectInt64(t *testing.T) {
	type args struct {
		json     string
		paths    []string
		entities []uint16
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"simple", args{
			json:     `[{"a": 3},{"b": 2},{"a": 1}]`,
			paths:    []string{"a"},
			entities: []uint16{2},
		},
			[]int64{int64(1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := codec.ParseJson([]byte(tt.args.json))
			if err != nil {
				t.Errorf("error to parse json %v", err)
			}

			index, err := index.Build(parsed)
			if err != nil {
				t.Errorf("Build() error = %v", err)
			}

			codes := make([][]byte, len(tt.args.paths))
			for i, p := range tt.args.paths {
				if code, ok := index.Meta.GetCode(bytes.Split([]byte(p), []byte("/"))); ok {
					codes[i] = code
				}
			}
			columns := make([]internal.Column, len(codes))
			for i := range columns {
				columns[i].Key = codes[i]
				columns[i].Ok = true
			}

			selected := Select(index, columns, tt.args.entities)
			int64s := selected[0].ColumnInt64
			var got []int64
			for i := range int64s.Rows {
				got = append(got, int64s.Columnar[int64s.Cols[i]])
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSelectInt64(b *testing.B) {
	entity := make([]uint16, benchSize)
	for i := range entity {
		entity[i] = uint16(i)
	}

	key, _ := benchIndex.Meta.GetCode(bytes.Split([]byte("a"), []byte{'/'}))
	columns := []internal.Column{
		{
			Key: key,
		},
	}

	//first, last := testData[0].OneInt64, testData[len(testData)-1].OneInt64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		selected := Select(benchIndex, columns, entity)
		if len(selected[0].ColumnInt64.Rows) != len(entity) {
			b.Error("selection error")
		} /* else if *selected[0][0].(*int64) != first || *selected[len(selected)-1][0].(*int64) != last {
			//b.Error("selection error")
		}*/
	}

}
