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

package ssd

import (
	"reflect"
	"testing"
)

func TestEntityMetaGenerateCode(t *testing.T) {
	type args struct {
		path [][]byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"single", args{[][]byte{[]byte("abc")}}, []byte{byte(32)}},
		{"nested single", args{[][]byte{[]byte("abc"), []byte("abc")}}, []byte{byte(32), byte(32)}},
		{"nested two", args{[][]byte{[]byte("abc"), []byte("bcd")}}, []byte{byte(32), byte(33)}},
	}
	s := NewEntityMeta()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.GenerateCode(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntityMetaGetCode(t *testing.T) {
	s := NewEntityMeta()
	s.GenerateCode([][]byte{[]byte("abc"), []byte("abc")})
	s.GenerateCode([][]byte{[]byte("bcd"), []byte("abc")})
	s.GenerateCode([][]byte{[]byte("abc"), []byte("bcd")})

	type args struct {
		path [][]byte
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 bool
	}{
		{"single true", args{[][]byte{[]byte("abc")}}, []byte{byte(32)}, true},
		{"single false", args{[][]byte{[]byte("abd")}}, nil, false},
		{"double true", args{[][]byte{[]byte("bcd"), []byte("abc")}}, []byte{byte(33), byte(32)}, true},
		{"double empty true", args{[][]byte{{}, []byte("bcd"), {}, []byte("abc")}}, []byte{byte(33), byte(32)}, true},
		{"double false", args{[][]byte{[]byte("abc"), []byte("abcd")}}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := s.GetCode(tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetCode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEntityMetaGetPath(t *testing.T) {
	s := NewEntityMeta()
	s.GenerateCode([][]byte{[]byte("abc"), []byte("abc")})
	s.GenerateCode([][]byte{[]byte("bcd"), []byte("abc")})
	s.GenerateCode([][]byte{[]byte("abc"), []byte("bcd")})

	data, err := s.Bytes()
	if err != nil {
		t.Errorf("failed to serialize %v", err)
	}

	type args struct {
		code []byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{"single", args{[]byte{byte(33)}}, [][]byte{[]byte("bcd")}},
		{"double", args{[]byte{byte(33), byte(32)}}, [][]byte{[]byte("bcd"), []byte("abc")}},
		{"not found", args{[]byte{byte(33), byte(33)}}, nil},
	}
	for _, tt := range tests {
		s, err = EntityMetaFromBytes(data)
		if err != nil {
			t.Errorf("failed to deserialize %v", err)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := s.GetPath(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
