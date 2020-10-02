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
	"github.com/chronowave/chronowave/ssql"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/internal"
	"github.com/chronowave/chronowave/ssd/operator"
)

func Exec(index *ssd.IndexedBlock, stmt *ssql.Statement) *ssd.ResultSet {
	node := map[string][]byte{}
	entity := eval(index, node, stmt)
	if len(entity) == 0 {
		return &ssd.ResultSet{RowId: []uint64{}}
	}

	columns := make([]internal.Column, len(stmt.Find))
	for i, f := range stmt.Find {
		columns[i].Key, columns[i].Ok = node[f.Name]
		columns[i].Name = f.Name
		columns[i].Group = f.Group
		columns[i].Func = f.Func
	}

	columns = operator.Select(index, columns, entity)
	return operator.Consolidate(index.ID, entity, columns)
}
