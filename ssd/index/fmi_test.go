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

package index

import (
	"reflect"
	"testing"

	hfmi "github.com/rleiwang/hfmi/ctor"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/codec"
)

func TestBuild(t *testing.T) {
	hfmi.SetSegmentCache(2038348)
	type args struct {
		json string
	}
	type result struct {
		entId    []uint32
		columnar ssd.Columnar
		hlt      ssd.HeaderLookupTable
		cda      []uint32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    result
	}{
		{"simple", args{`[{ "d": "c"}, {"a": "c"}, {"d" : "f"}]`}, false,
			result{
				entId: []uint32{0, 2, 4},
				columnar: ssd.Columnar{
					Float64: []float64{},
					Int64:   []int64{},
					Bool:    []bool{},
					Text:    []uint32{1, 2, 1},
				},
				hlt: ssd.HeaderLookupTable{
					Text: ssd.HeaderISA{
						Entity:    []uint16{0, 2, 1},
						Attribute: []uint16{0, 0, 0},
					},
				},
				cda: []uint32{0},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := codec.ParseJson([]byte(tt.args.json))
			if err != nil {
				t.Errorf("error to parse json %v", err)
			}

			got, err := Build(parsed)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Error("runtime error to build")
			}
			gotResult := result{
				entId:    got.EntityID,
				columnar: got.Columnar,
				hlt:      got.HLT,
				cda:      got.HeaderDA,
			}
			if !reflect.DeepEqual(gotResult, tt.want) {
				t.Errorf("Build() got = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
