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

	hfmi "github.com/rleiwang/hfmi/ctor"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/index"
	"github.com/chronowave/chronowave/ssd/internal"
)

func TestMarshalToJson(t *testing.T) {
	hfmi.SetSegmentCache(2038348)
	type args struct {
		json []byte
		eav  []internal.Attribute
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"simple", args{
			json: []byte(`{ "a": [{"b": 2}, {"b": 3}, { "b": 4}], "b": { "c": 12 }, "c": { "a": 23 } }`),
			eav: []internal.Attribute{
				{
					Code:      []byte{32},
					Offset:    0,
					ValueType: ssd.SOA,
				},
				{
					Code:      []byte{32, 33},
					Offset:    2,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32, 33},
					Offset:    3,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32},
					Offset:    0,
					ValueType: ssd.EOA,
				},
				{
					Code:      []byte{33, 34},
					Offset:    0,
					ValueType: ssd.INT64,
					Value:     10,
				},
				{
					Code:      []byte{34, 32},
					Offset:    0,
					ValueType: ssd.INT64,
					Value:     10,
				},
			},
		}, []byte(`{"a":[{"b":2},{"b":2}],"b":{"c":10},"c":{"a":10}}`)},
		{"simple", args{
			json: []byte(`{ "a": [2, 3, 4], "b": { "c": 12 }, "c": { "a": 23 } }`),
			eav: []internal.Attribute{
				{
					Code:      []byte{32},
					Offset:    0,
					ValueType: ssd.SOA,
				},
				{
					Code:      []byte{32},
					Offset:    2,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32},
					Offset:    3,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32},
					Offset:    0,
					ValueType: ssd.EOA,
				},
				{
					Code:      []byte{33, 34},
					Offset:    0,
					ValueType: ssd.INT64,
					Value:     10,
				},
				{
					Code:      []byte{34, 32},
					Offset:    0,
					ValueType: ssd.INT64,
					Value:     10,
				},
			},
		}, []byte(`{"a":[2,2],"b":{"c":10},"c":{"a":10}}`)},
		{"nested", args{
			json: []byte(`{ "a": [{"b": [2, 3]}, {"b": [2, 3]}, {"b": [3, 4]}], "b": { "c": 12 }, "c": { "a": 23 } }`),
			eav: []internal.Attribute{
				{
					Code:      []byte{32},
					Offset:    0,
					ValueType: ssd.SOA,
				},
				{
					Code:      []byte{32, 33},
					Offset:    0,
					ValueType: ssd.SOA,
					Value:     2,
				},
				{
					Code:      []byte{32, 33},
					Offset:    1,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32, 33},
					Offset:    2,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32, 33},
					Offset:    0,
					ValueType: ssd.EOA,
					Value:     2,
				},
				{
					Code:      []byte{32, 33},
					Offset:    0,
					ValueType: ssd.SOA,
					Value:     2,
				},
				{
					Code:      []byte{32, 33},
					Offset:    3,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32, 33},
					Offset:    4,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32, 33},
					Offset:    0,
					ValueType: ssd.EOA,
					Value:     2,
				},
				{
					Code:      []byte{32},
					Offset:    0,
					ValueType: ssd.EOA,
				},
				{
					Code:      []byte{33, 34},
					Offset:    0,
					ValueType: ssd.INT64,
					Value:     10,
				},
				{
					Code:      []byte{34, 32},
					Offset:    0,
					ValueType: ssd.INT64,
					Value:     12,
				},
			},
		}, []byte(`{"a":[{"b":[2,2]},{"b":[2,2]}],"b":{"c":10},"c":{"a":12}}`)},
		{"array nested object", args{
			json: []byte(`{ "a": [{"b": {"c":  2}}, {"b": {"c": 3}}, { "b": { "c": 4}}], "b": { "c": 12 }, "c": { "a": 23 } }`),
			eav: []internal.Attribute{
				{
					Code:      []byte{32},
					Offset:    0,
					ValueType: ssd.SOA,
				},
				{
					Code:      []byte{32, 33, 34},
					Offset:    1,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32, 33, 34},
					Offset:    2,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32, 33, 34},
					Offset:    3,
					ValueType: ssd.INT64,
					Value:     2,
				},
				{
					Code:      []byte{32},
					Offset:    0,
					ValueType: ssd.EOA,
				},
				{
					Code:      []byte{33, 34},
					Offset:    0,
					ValueType: ssd.INT64,
					Value:     12,
				},
				{
					Code:      []byte{34, 32},
					Offset:    0,
					ValueType: ssd.INT64,
					Value:     12,
				},
			},
		}, []byte(`{"a":[{"b":{"c":2}},{"b":{"c":2}},{"b":{"c":2}}],"b":{"c":12},"c":{"a":12}}`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := ParseJson(tt.args.json)
			if err != nil {
				t.Errorf("error to parse json %v", err)
			}

			indexed, err := index.Build(parsed)
			if err != nil {
				t.Errorf("index Build() error = %v", err)
			}

			holder := map[string]interface{}{}
			err = json.Unmarshal(tt.want, &holder)
			if err != nil {
				t.Errorf("json Unmarshall() error = %v", err)
			}
			if got := MarshalToJson(indexed, nil, tt.args.eav, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Restore() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
