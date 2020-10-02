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
	"bytes"
	"reflect"
	"testing"
)

func TestTextColumnarAdd(t *testing.T) {
	type args struct {
		texts []string
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{"two texts", args{[]string{"abc", "bcd"}}, []uint32{1, 2}},
		{"three texts", args{[]string{"abc", "bcd", "abc"}}, []uint32{1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint32
			tc := NewTextColumnar()

			wantTexts := bytes.NewBuffer(make([]byte, 0, 1024*1024))
			exists := map[string]bool{}
			for _, t := range tt.args.texts {
				got = append(got, tc.Add(t))
				if _, ok := exists[t]; !ok {
					if wantTexts.Len() > 0 {
						wantTexts.WriteByte(SOH)
					}
					wantTexts.WriteString(t)
					exists[t] = true
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}

			gotTexts := tc.Value()
			if !reflect.DeepEqual(gotTexts, wantTexts.Bytes()) {
				t.Errorf("Add() = %v, want %v", gotTexts, wantTexts)
			}
		})
	}
}
