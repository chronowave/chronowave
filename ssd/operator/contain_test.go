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
	jsonenc "encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssd/codec"
	"github.com/chronowave/chronowave/ssd/index"
	"github.com/chronowave/chronowave/ssd/internal"
)

func TestContain(t *testing.T) {
	type args struct {
		json    string
		path    string
		pattern string
	}
	tests := []struct {
		name          string
		args          args
		wantEntity    []uint16
		wantAttribute []uint16
	}{
		{"simple", args{
			json:    `[{ "d": "c"}, {"a": "c"}, {"d" : "f"}]`,
			path:    `d`,
			pattern: `f`,
		}, []uint16{2}, []uint16{0}},
		{"fragment single first", args{
			json:    `[{ "d": "ccddcc"}, {"a": "c"}]`,
			path:    `d`,
			pattern: `cc`,
		}, []uint16{0}, []uint16{0}},
		{"fragment single last", args{
			json:    `[{"a": "c"}, {"d" : "fccf"}]`,
			path:    `d`,
			pattern: `cc`,
		}, []uint16{1}, []uint16{0}},
		{"fragment double", args{
			json:    `[{ "d": "ccddcc"}, {"a": "c"}, {"d" : "fccf"}]`,
			path:    `d`,
			pattern: `cc`,
		}, []uint16{0, 2}, []uint16{0, 0}},
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

			key, ok := index.Meta.GetCode(bytes.Split([]byte(tt.args.path), []byte{'/'}))
			if !ok {
				t.Errorf("unable to locate the path")
			}
			gotEntity, gotAttribute := Contain(index, key, []byte(tt.args.pattern))
			if !reflect.DeepEqual(gotEntity, tt.wantEntity) {
				t.Errorf("Contain() gotEntity = %v, wantEntity %v", gotEntity, tt.wantEntity)
			}
			if !reflect.DeepEqual(gotAttribute, tt.wantAttribute) {
				t.Errorf("Contain() gotAttribute = %v, wantAttribute %v", gotAttribute, tt.wantAttribute)
			}
		})
	}
}

func BenchmarkContain(b *testing.B) {
	// 4ms 128, 1.3ms 32
	internal.FragmentSize = 128
	data, err := ioutil.ReadFile("server.log_2015-09-25T16-22-16")
	if err != nil {
		b.Errorf("error load test data %v", err)
	}

	data, err = jsonenc.Marshal(map[string]interface{}{
		"aa": string(data),
	})
	if err != nil {
		b.Errorf("error json marshall test data %v", err)
	}

	parsed, err := codec.ParseJson(data)
	if err != nil {
		b.Errorf("error to parse json %v", err)
	}

	indexed, err := index.Build(parsed)
	if err != nil {
		b.Errorf("Build() error = %v", err)
	}

	key, _ := indexed.Meta.GetCode([][]byte{[]byte("aa")})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gotEntity, _ := Contain(indexed, key, []byte("Exception"))
		if len(gotEntity) == 0 {
			b.Errorf("Contain() error = %v", err)
		}
	}
}
