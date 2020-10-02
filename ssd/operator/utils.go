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

// intersectUint32 returns index in X if X[i] exists in Y[]
func intersectUint32(x, y []uint32) []uint32 {
	idx, cnt := make([]uint32, len(x)), 0
	mapY := map[uint32]bool{}
	for _, v := range y {
		mapY[v] = true
	}
	for i, v := range x {
		if _, ok := mapY[v]; ok {
			idx[cnt] = uint32(i)
			cnt++
		}
	}

	return idx[:cnt]
}
