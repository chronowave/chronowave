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
	"bytes"
	"sync"

	ext "github.com/chronowave/ext/operator"

	"github.com/chronowave/chronowave/ssql"

	"github.com/rs/zerolog/log"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/internal"
	"github.com/chronowave/chronowave/ssd/operator"
)

var (
	entityPool = &sync.Pool{
		New: func() interface{} {
			return make([]uint16, 1<<16)
		},
	}

	iePool = &sync.Pool{
		New: func() interface{} {
			return make([]uint32, 1<<16)
		},
	}

	empty []uint16
)

func eval(index *ssd.IndexedBlock, node map[string][]byte, stmt *ssql.Statement) []uint16 {
	entity, ie := entityPool.Get().([]uint16), iePool.Get().([]uint32)
	refEntity, refIE := entity, ie
	defer func() {
		entityPool.Put(refEntity)
		iePool.Put(refIE)
	}()

	entity, ie = entity[:len(index.EntityID)], ie[:len(index.EntityID)]
	for i := range entity {
		entity[i] = uint16(i)
	}

	entity = evalAnd(index, node, stmt.Where, entity, ie)

	clone := make([]uint16, len(entity))
	copy(clone, entity)

	return clone
}

func getKey(index *ssd.IndexedBlock, node map[string][]byte, tuple *ssql.Tuple) ([]byte, bool) {
	path := bytes.Split([]byte(tuple.Path), []byte{'/'})
	if len(tuple.Name) == 0 {
		return index.Meta.GetCode(path)
	}

	key, ok := node[tuple.Name]
	if !ok {
		// key code
		if key, ok = index.Meta.GetCode(path); ok {
			node[tuple.Name] = key
		} else {
			// TODO: add trace to keep track state
			log.Info().Msgf("path %v doesn't exist in block %v", tuple.Path, index.ID)
		}
	}

	return key, ok
}

func evalAnd(index *ssd.IndexedBlock, node map[string][]byte, and []*ssql.Expr, entity []uint16, ie []uint32) []uint16 {
	// default is AND
	for _, expr := range and {
		var selected []uint16
		switch expr.Field.(type) {
		case *ssql.Expr_Tuple:
			tuple := expr.Field.(*ssql.Expr_Tuple).Tuple
			key, ok := getKey(index, node, tuple)
			if tuple.Predicate == nil {
				// no predicate, select only
				continue
			} else if ok {
				selected = evalTuple(index, node, key, expr.Field.(*ssql.Expr_Tuple).Tuple)
			} else {
				selected = empty
			}
		case *ssql.Expr_Or:
			selected = evalOr(index, node, expr.Field.(*ssql.Expr_Or).Or.Expr)
		}
		cnt := ext.IntersectXUint16(entity, selected, ie)
		if cnt == 0 {
			return empty
		}

		for i, e := range ie[:cnt] {
			entity[i] = entity[e]
		}
		entity = entity[:cnt]
	}

	return entity
}

func evalOr(index *ssd.IndexedBlock, node map[string][]byte, or []*ssql.Expr) []uint16 {
	entity := make([]bool, len(index.EntityID))
	for _, expr := range or {
		var selected []uint16
		switch expr.Field.(type) {
		case *ssql.Expr_Tuple:
			tuple := expr.Field.(*ssql.Expr_Tuple).Tuple
			key, ok := getKey(index, node, tuple)
			if tuple.Predicate == nil {
				// no predicate, select only
				continue
			} else if ok {
				selected = evalTuple(index, node, key, expr.Field.(*ssql.Expr_Tuple).Tuple)
			}
		case *ssql.Expr_Or:
			selected = evalOr(index, node, expr.Field.(*ssql.Expr_Or).Or.Expr)
		}

		for _, x := range selected {
			entity[x] = true
		}
	}

	rs := make([]uint16, len(index.EntityID))

	j := 0
	for i, b := range entity {
		if b {
			rs[j] = uint16(i)
			j++
		}
	}

	return rs[:j]
}

func evalTuple(index *ssd.IndexedBlock, node map[string][]byte, key []byte, expr *ssql.Tuple) []uint16 {
	entity := empty
	switch expr.Predicate.(type) {
	case *ssql.Tuple_Nested:
		return evalNested(index, node, expr, expr.Predicate.(*ssql.Tuple_Nested).Nested.Expr)
	case *ssql.Tuple_Between:
		entity, _ = evalBetween(index, key, expr.Predicate.(*ssql.Tuple_Between))
	case *ssql.Tuple_Timeframe:
		entity, _ = evalTimeframe(index, key, expr.Predicate.(*ssql.Tuple_Timeframe))
	case *ssql.Tuple_Key:
		entity, _ = evalKey(index, key, expr.Predicate.(*ssql.Tuple_Key))
	case *ssql.Tuple_Contain:
		entity, _ = evalContain(index, key, expr.Predicate.(*ssql.Tuple_Contain))
	case *ssql.Tuple_Eq:
		entity, _ = evalEqual(index, key, expr.Predicate.(*ssql.Tuple_Eq))
	case *ssql.Tuple_Neq:
		entity, _ = evalNotEqual(index, key, expr.Predicate.(*ssql.Tuple_Neq))
	case *ssql.Tuple_Gt:
		entity, _ = evalGreaterThan(index, key, expr.Predicate.(*ssql.Tuple_Gt))
	case *ssql.Tuple_Ge:
		entity, _ = evalGreaterEqual(index, key, expr.Predicate.(*ssql.Tuple_Ge))
	case *ssql.Tuple_Lt:
		entity, _ = evalLessThan(index, key, expr.Predicate.(*ssql.Tuple_Lt))
	case *ssql.Tuple_Le:
		entity, _ = evalLessEqual(index, key, expr.Predicate.(*ssql.Tuple_Le))
	case *ssql.Tuple_In:
		entity, _ = evalIn(index, key, expr.Predicate.(*ssql.Tuple_In))
	case *ssql.Tuple_Exist:
		entity, _ = evalExist(index, key, expr.Predicate.(*ssql.Tuple_Exist))
	}

	return entity
}

func evalNested(index *ssd.IndexedBlock, node map[string][]byte, parent *ssql.Tuple, nested []*ssql.Expr) []uint16 {
	docs := make([]internal.Document, len(nested))
	for i, expr := range nested {
		docs[i].Entity, docs[i].Attribute = empty, empty
		switch expr.Field.(type) {
		case *ssql.Expr_Tuple:
			tuple := expr.Field.(*ssql.Expr_Tuple).Tuple
			if tuple.Predicate == nil {
				// no predicate, select only
				return empty
			}
			tuple = &ssql.Tuple{
				Name:      tuple.Name,
				Path:      parent.Path + "/" + tuple.Path,
				Predicate: tuple.Predicate,
			}

			if key, ok := getKey(index, node, tuple); ok {
				switch tuple.Predicate.(type) {
				case *ssql.Tuple_Between:
					docs[i].Entity, docs[i].Attribute = evalBetween(index, key, tuple.Predicate.(*ssql.Tuple_Between))
				case *ssql.Tuple_Contain:
					docs[i].Entity, docs[i].Attribute = evalContain(index, key, tuple.Predicate.(*ssql.Tuple_Contain))
				case *ssql.Tuple_Eq:
					docs[i].Entity, docs[i].Attribute = evalEqual(index, key, tuple.Predicate.(*ssql.Tuple_Eq))
				case *ssql.Tuple_Neq:
					docs[i].Entity, docs[i].Attribute = evalNotEqual(index, key, tuple.Predicate.(*ssql.Tuple_Neq))
				case *ssql.Tuple_Gt:
					docs[i].Entity, docs[i].Attribute = evalGreaterThan(index, key, tuple.Predicate.(*ssql.Tuple_Gt))
				case *ssql.Tuple_Ge:
					docs[i].Entity, docs[i].Attribute = evalGreaterEqual(index, key, tuple.Predicate.(*ssql.Tuple_Ge))
				case *ssql.Tuple_Lt:
					docs[i].Entity, docs[i].Attribute = evalLessThan(index, key, tuple.Predicate.(*ssql.Tuple_Lt))
				case *ssql.Tuple_Le:
					docs[i].Entity, docs[i].Attribute = evalLessEqual(index, key, tuple.Predicate.(*ssql.Tuple_Le))
				case *ssql.Tuple_In:
					docs[i].Entity, docs[i].Attribute = evalIn(index, key, tuple.Predicate.(*ssql.Tuple_In))
				case *ssql.Tuple_Exist:
					docs[i].Entity, docs[i].Attribute = evalExist(index, key, tuple.Predicate.(*ssql.Tuple_Exist))
				}
			}
		}
	}

	entity, _ := operator.Nested(index, docs)
	return entity
}

func evalBetween(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_Between) ([]uint16, []uint16) {
	x, y := op.Between.First, op.Between.Second
	switch x.Value.(type) {
	case *ssql.Operand_Int:
		if yv, ok := y.Value.(*ssql.Operand_Int); ok {
			return operator.BinaryInt64(index, key, x.Value.(*ssql.Operand_Int).Int, yv.Int, ext.BetweenInt64)
		} else if yv, ok := y.Value.(*ssql.Operand_Double); ok {
			return operator.BinaryFloat64(index, key, float64(x.Value.(*ssql.Operand_Int).Int), yv.Double, ext.BetweenFloat64)
		}
	case *ssql.Operand_Double:
		if yv, ok := y.Value.(*ssql.Operand_Double); ok {
			return operator.BinaryFloat64(index, key, x.Value.(*ssql.Operand_Double).Double, yv.Double, ext.BetweenFloat64)
		} else if yv, ok := y.Value.(*ssql.Operand_Int); ok {
			return operator.BinaryFloat64(index, key, x.Value.(*ssql.Operand_Double).Double, float64(yv.Int), ext.BetweenFloat64)
		}
	}

	return empty, empty
}

func evalTimeframe(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_Timeframe) ([]uint16, []uint16) {
	return operator.BinaryInt64(index, key, op.Timeframe.First.Value.(*ssql.Operand_Int).Int,
		op.Timeframe.Second.Value.(*ssql.Operand_Int).Int, ext.BetweenInt64)
}

func evalContain(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_Contain) ([]uint16, []uint16) {
	if text, ok := op.Contain.First.Value.(*ssql.Operand_Text); ok {
		return operator.Contain(index, key, []byte(text.Text))
	}

	return empty, empty
}

func evalKey(index *ssd.IndexedBlock, path []byte, op *ssql.Tuple_Key) ([]uint16, []uint16) {
	text := op.Key.First.Value.(*ssql.Operand_Text).Text
	key := make([]byte, len(text)+2)
	copy(key[1:], text)
	key[0], key[len(key)-1] = '^', '$'
	return operator.Contain(index, path, key)
}

func evalEqual(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_Eq) ([]uint16, []uint16) {
	v := op.Eq.First.Value
	switch v.(type) {
	case *ssql.Operand_Int:
		return operator.UnaryInt64(index, key, v.(*ssql.Operand_Int).Int, ext.EqualInt64)
	case *ssql.Operand_Double:
		return operator.UnaryFloat64(index, key, v.(*ssql.Operand_Double).Double, ext.EqualFloat64)
	}

	return empty, empty
}

func evalNotEqual(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_Neq) ([]uint16, []uint16) {
	v := op.Neq.First.Value
	switch v.(type) {
	case *ssql.Operand_Int:
		return operator.UnaryInt64(index, key, v.(*ssql.Operand_Int).Int, ext.NotEqualInt64)
	case *ssql.Operand_Double:
		return operator.UnaryFloat64(index, key, v.(*ssql.Operand_Double).Double, ext.NotEqualFloat64)
	}

	return empty, empty
}

func evalGreaterThan(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_Gt) ([]uint16, []uint16) {
	v := op.Gt.First.Value
	switch v.(type) {
	case *ssql.Operand_Int:
		return operator.UnaryInt64(index, key, v.(*ssql.Operand_Int).Int, ext.GreaterThanInt64)
	case *ssql.Operand_Double:
		return operator.UnaryFloat64(index, key, v.(*ssql.Operand_Double).Double, ext.GreaterThanFloat64)
	}

	return empty, empty
}

func evalGreaterEqual(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_Ge) ([]uint16, []uint16) {
	v := op.Ge.First.Value
	switch v.(type) {
	case *ssql.Operand_Int:
		return operator.UnaryInt64(index, key, v.(*ssql.Operand_Int).Int, ext.GreaterEqualInt64)
	case *ssql.Operand_Double:
		return operator.UnaryFloat64(index, key, v.(*ssql.Operand_Double).Double, ext.GreaterEqualFloat64)
	}

	return empty, empty
}

func evalLessThan(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_Lt) ([]uint16, []uint16) {
	v := op.Lt.First.Value
	switch v.(type) {
	case *ssql.Operand_Int:
		return operator.UnaryInt64(index, key, v.(*ssql.Operand_Int).Int, ext.LessThanInt64)
	case *ssql.Operand_Double:
		return operator.UnaryFloat64(index, key, v.(*ssql.Operand_Double).Double, ext.LessThanFloat64)
	}

	return empty, empty
}

func evalLessEqual(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_Le) ([]uint16, []uint16) {
	v := op.Le.First.Value
	switch v.(type) {
	case *ssql.Operand_Int:
		return operator.UnaryInt64(index, key, v.(*ssql.Operand_Int).Int, ext.LessEqualInt64)
	case *ssql.Operand_Double:
		return operator.UnaryFloat64(index, key, v.(*ssql.Operand_Double).Double, ext.LessEqualFloat64)
	}

	return empty, empty
}

func evalIn(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_In) ([]uint16, []uint16) {
	if list, ok := op.In.First.Value.(*ssql.Operand_List); ok {
		if len(list.List.Text) > 0 {
			return operator.InText(index, key, list.List.Text)
		} else if len(list.List.Int) > 0 {
			return operator.InInt(index, key, list.List.Int)
		} else if len(list.List.Double) > 0 {
			return operator.InFloat(index, key, list.List.Double)
		}
	}

	return empty, empty
}

func evalExist(index *ssd.IndexedBlock, key []byte, op *ssql.Tuple_Exist) ([]uint16, []uint16) {
	return operator.Exist(index, key)
}
