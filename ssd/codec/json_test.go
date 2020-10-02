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
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssd"
)

func TestParseEmpty(t *testing.T) {
	type args struct {
		json string
	}
	type result struct {
		meta []byte
	}
	tests := []struct {
		name string
		args args
		want result
	}{
		{"empty array", args{`[]`}, result{[]byte{}}},
		{"empty object", args{`{}`}, result{[]byte{}}},
		{"empty array object", args{`[{}]`}, result{[]byte{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJson([]byte(tt.args.json))
			if err != nil {
				t.Errorf("parsing error %v", err)
			}
			actual := result{
				meta: got.Entity.Bytes(),
			}
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("parse() = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestParseError(t *testing.T) {
	type args struct {
		json string
	}
	tests := []struct {
		name string
		args args
	}{
		{"error object", args{`{ key1: "adf"}`}},
		{"empty array", args{`[{ key1: "adf"}]`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseJson([]byte(tt.args.json))
			if err == nil {
				t.Errorf("didn't get expected parsing error")
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		json string
	}
	type result struct {
		count    uint32
		meta     []byte
		columnar ssd.Columnar
		content  []byte
	}
	tests := []struct {
		name string
		args args
		want result
	}{
		{"empty string object", args{`[{ "key1": ""}]`},
			result{
				count: 1,
				meta:  []byte{byte(32), ssd.TEXT},
				columnar: ssd.Columnar{
					Text: []uint32{0},
				},
				content: []byte{},
			},
		},
		{"string object", args{`[{ "key1": "val"}]`},
			result{
				count: 1,
				meta:  []byte{byte(32), ssd.TEXT},
				columnar: ssd.Columnar{
					Text: []uint32{1},
				},
				content: []byte("val"),
			},
		},
		{"int object", args{`[{"key1" : 1 } ]`},
			result{
				count: 1,
				meta:  []byte{byte(32), ssd.INT64},
				columnar: ssd.Columnar{
					Int64: []int64{1},
				},
				content: []byte{},
			},
		},
		{"float object", args{`[{"key1" : -0.1 } ]`},
			result{
				count: 1,
				meta:  []byte{byte(32), ssd.FLT64},
				columnar: ssd.Columnar{
					Float64: []float64{-0.1},
				},
				content: []byte{},
			},
		},
		{"float array value", args{`[{"key1" : -0.1, "key2": [1.5, 0.5] } ]`},
			result{
				count: 1,
				meta:  []byte{byte(32), ssd.FLT64, ssd.SOH, byte(33), ssd.SOA, ssd.SOH, byte(33), ssd.FLT64, ssd.SOH, ssd.AED, ssd.SOH, byte(33), ssd.FLT64, ssd.SOH, byte(33), ssd.EOA},
				columnar: ssd.Columnar{
					Float64: []float64{-0.1, 1.5, 0.5},
				},
				content: []byte{},
			},
		},
		{"nested object", args{`[{"key1" : { "key1": true, "key2": false} } ]`},
			result{
				count: 1,
				meta:  []byte{byte(32), byte(32), ssd.BOOL, ssd.SOH, byte(32), byte(33), ssd.BOOL},
				columnar: ssd.Columnar{
					Bool: []bool{true, false},
				},
				content: []byte{},
			},
		},
		{"null object", args{`[{"key1" : { "key1": null, "key2": null} } ]`},
			result{
				count:   1,
				meta:    []byte{byte(32), byte(32), ssd.NULL, ssd.SOH, byte(32), byte(33), ssd.NULL},
				content: []byte{},
			},
		},
		{"special key", args{`[{"k/ey1" : { "key1": null, "key2": null} } ]`},
			result{
				count:   1,
				meta:    []byte{byte(32), byte(33), ssd.NULL, ssd.SOH, byte(32), byte(34), ssd.NULL},
				content: []byte{},
			},
		},
		{"multiple object", args{`[{"key1" : 2},{ "key2": 3}]`},
			result{
				count: 2,
				meta:  []byte{byte(32), ssd.INT64, ssd.SOH, ssd.EOO, ssd.SOH, byte(33), ssd.INT64},
				columnar: ssd.Columnar{
					Int64: []int64{2, 3},
				},
				content: []byte{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJson([]byte(tt.args.json))
			if err != nil {
				t.Errorf("parsing error %v", err)
			}
			actual := result{
				count:    got.Count,
				meta:     got.Entity.Bytes(),
				columnar: got.Columnar,
				content:  got.Content.Value(),
			}
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("parse() = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestFragmentLongText(t *testing.T) {
	type args struct {
		value []byte
		fsz   int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"short", args{[]byte("abc"), 4}, []byte("abc")},
		{"long", args{[]byte("abc"), 2}, []byte{'a', 'b', ssd.FRAG, 'c'}},
		{"same", args{[]byte("cba"), 3}, []byte("cba")},
		{"same 2", args{[]byte("cbanba"), 3}, append(append([]byte("cba"), ssd.FRAG), []byte("nba")...)},
		{"control short", args{[]byte{'a', ssd.TEXT, 'c'}, 4}, []byte{'a', ' ', 'c'}},
		{"control long", args{[]byte{ssd.INT64, 'b', ssd.INT64, 'c'}, 2}, []byte{' ', 'b', ssd.FRAG, ' ', 'c'}},
		{"control same", args{[]byte{'c', ssd.TEXT, 'a'}, 3}, []byte{'c', ' ', 'a'}},
		{"control same 2", args{[]byte{'c', ssd.TEXT, 'a', 'a', ssd.INT16, 'c'}, 3}, []byte{'c', ' ', 'a', ssd.FRAG, 'a', ' ', 'c'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fragmentLongText(tt.args.value, tt.args.fsz); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fragmentLongText() = %v, want %v", got, tt.want)
			}
		})
	}
}
