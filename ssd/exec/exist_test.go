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

package exec

import (
	jsonenc "encoding/json"
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssql/parser"

	"github.com/chronowave/chronowave/ssd/codec"
)

func TestExist(t *testing.T) {
	type args struct {
		json string
		stmt string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"root", args{
			json: `[{"a": [{"b": 1, "c": 3}, {"b": 2, "c": 1}], "c": "dd"},
                    {"a": [{"b": 2, "c": 2}, {"b": 3, "c": 2}], "d": "a"},
                    {"a": [{"b": 3, "c": 2}, {"b": 3, "c": 3}]}]`,
			stmt: `find $val where [$val /] [/d exist]`},
			[]byte(`[{"val":{"a":[{"b":2,"c":2},{"b":3,"c":2}],"d":"a"}}]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := codec.ParseJson([]byte(tt.args.json))
			if err != nil {
				t.Errorf("error to parse json %v", err)
			}

			indexed, err := buildTestIndex(parsed)
			if err != nil {
				t.Errorf("index Build() error = %v", err)
			}

			stmt, e := parser.Parse(tt.args.stmt)
			if len(e) > 0 {
				t.Errorf("ssql Parse() error = %v", e)
			}
			rs := Exec(indexed, stmt)

			got := codec.MarshalResultSet(rs, 0)

			var data []map[string]interface{}
			err = jsonenc.Unmarshal(got, &data)

			if err != nil {
				t.Errorf("invalid json response %v", string(got))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
