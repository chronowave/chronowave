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

import "bytes"

type TextColumnar struct {
	texts  *bytes.Buffer
	offset map[string]int
}

func NewTextColumnar() *TextColumnar {
	return &TextColumnar{
		texts:  bytes.NewBuffer(make([]byte, 0, 1024*1024)),
		offset: map[string]int{},
	}
}

// Add add text t to text columnar, returns offset in the column
func (tc *TextColumnar) Add(t string) uint32 {
	i, ok := tc.offset[t]
	if !ok {
		i = len(tc.offset)
		if i > 0 {
			tc.texts.WriteByte(SOH)
		}
		tc.texts.WriteString(t)
		tc.offset[t] = i
	}

	return uint32(i) + 1
}

func (tc *TextColumnar) Value() []byte {
	return tc.texts.Bytes()
}
