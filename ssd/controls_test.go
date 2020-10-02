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

import "testing"

func TestIsControlCharacter(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sentinel", args{byte(0)}, true},
		{"ack", args{byte(6)}, true},
		{"bell", args{byte(7)}, false},
		{"bs", args{byte(8)}, false},
		{"ht", args{byte(9)}, false},
		{"nl", args{byte(10)}, false},
		{"vt", args{byte(11)}, false},
		{"ff", args{byte(12)}, false},
		{"cr", args{byte(13)}, false},
		{"shift out", args{byte(14)}, true},
		{"shift in", args{byte(15)}, true},
		{"space", args{byte(32)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsControlCharacter(tt.args.c); got != tt.want {
				t.Errorf("IsControlCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
