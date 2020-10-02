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

	"github.com/chronowave/chronowave/ssd/internal"
)

func TestNested(t *testing.T) {
	type args struct {
		docs []internal.Document
	}
	tests := []struct {
		name  string
		args  args
		want  []uint16
		want1 []uint16
	}{
		{"single", args{[]internal.Document{
			{
				Entity:    []uint16{0, 1},
				Attribute: []uint16{12, 13},
			},
		}}, []uint16{0, 1}, []uint16{12, 13}},
		{"double", args{[]internal.Document{
			{
				Entity:    []uint16{0, 1},
				Attribute: []uint16{12, 13},
			},
			{
				Entity:    []uint16{0, 1},
				Attribute: []uint16{11, 13},
			},
		}}, []uint16{1}, []uint16{13}},
		{"triple", args{[]internal.Document{
			{
				Entity:    []uint16{1, 6, 12},
				Attribute: []uint16{14, 12, 13},
			},
			{
				Entity:    []uint16{0, 12, 14},
				Attribute: []uint16{11, 13, 28},
			},
			{
				Entity:    []uint16{0, 1, 12},
				Attribute: []uint16{11, 13, 13},
			},
		}}, []uint16{12}, []uint16{13}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Nested(nil, tt.args.docs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Nested() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Nested() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func BenchmarkNested(b *testing.B) {
	docs := make([]internal.Document, 3)
	for i := range docs {
		d := &docs[i]
		d.Entity, d.Attribute = make([]uint16, 1<<16), make([]uint16, 1<<16)
		for i := 0; i < 1<<16; i++ {
			d.Entity[i], d.Attribute[i] = uint16(i), uint16(i)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Nested(nil, docs)
	}
}
