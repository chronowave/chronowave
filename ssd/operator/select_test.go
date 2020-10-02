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

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/codec"
	"github.com/chronowave/chronowave/ssd/index"
	"github.com/chronowave/chronowave/ssd/internal"
)

func TestSelectBool(t *testing.T) {
	type args struct {
		json     string
		paths    []string
		entities []uint16
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"simple", args{
			json:     `[{"a": false},{"b": false},{"a": true}]`,
			paths:    []string{"a"},
			entities: []uint16{2},
		},
			[]bool{true},
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
			booleans := selected[0].ColumnBool
			var got []bool
			for i := range booleans.Rows {
				got = append(got, booleans.Columnar[booleans.Cols[i]])
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestSelectNull(t *testing.T) {
	type args struct {
		json     string
		paths    []string
		entities []uint16
	}
	tests := []struct {
		name string
		args args
		want []internal.Column
	}{
		{"simple", args{
			json:     `[{"a": null},{"b": false},{"a": true}]`,
			paths:    []string{"a"},
			entities: []uint16{0},
		}, []internal.Column{{
			Key:  []byte{32},
			Ok:   true,
			Name: "a",
			Type: ssd.BOOL,
		}}},
		{"no path", args{
			json:     `[{"a": null},{"b": false},{"a": true}]`,
			paths:    []string{"k"},
			entities: []uint16{0},
		}, []internal.Column{{
			Key:  nil,
			Ok:   false,
			Name: "k",
			Type: ssd.NULL,
		}}},
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

			columns := make([]internal.Column, len(tt.args.paths))
			for i, p := range tt.args.paths {
				columns[i].Name = p
				columns[i].Key, columns[i].Ok = index.Meta.GetCode(bytes.Split([]byte(p), []byte("/")))
			}

			got := Select(index, columns, tt.args.entities)
			got[0].ColumnBool = nil
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
