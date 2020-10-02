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

package parser

import (
	"strconv"

	"github.com/chronowave/chronowave/ssql"
	"github.com/chronowave/chronowave/ssql/parser/gen"
)

func (p *parser) VisitTimeframe(ctx *gen.TimeframeContext) interface{} {
	var (
		first  int64
		second int64
		err    error
	)
	operands := ctx.AllINTEGER()
	if len(operands) == 0 {
		token := ctx.GetStart()
		p.errors = append(p.errors, Error{
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Message: "missing integer values",
		})
	} else {
		first, err = strconv.ParseInt(operands[0].GetText(), 10, 64)
		if err != nil {
			token := operands[0].GetSymbol()
			p.errors = append(p.errors, Error{
				Line:    token.GetLine(),
				Column:  token.GetColumn(),
				Message: err.Error(),
			})
		}

		second, err = strconv.ParseInt(operands[1].GetText(), 10, 64)
		if err != nil {
			token := operands[1].GetSymbol()
			p.errors = append(p.errors, Error{
				Line:    token.GetLine(),
				Column:  token.GetColumn(),
				Message: err.Error(),
			})
		}

		if first > second {
			token := ctx.GetStart()
			p.errors = append(p.errors, Error{
				Line:    token.GetLine(),
				Column:  token.GetColumn(),
				Message: "timeframe second operand must be greater",
			})
		}
	}

	return &ssql.Tuple{
		Predicate: &ssql.Tuple_Timeframe{
			Timeframe: &ssql.Binary{
				First: &ssql.Operand{
					Value: &ssql.Operand_Int{
						Int: first,
					},
				},
				Second: &ssql.Operand{
					Value: &ssql.Operand_Int{
						Int: second,
					},
				},
			},
		},
	}
}

func (p *parser) VisitKey(ctx *gen.KeyContext) interface{} {
	if ctx.INTEGER() != nil {
		v, err := strconv.ParseInt(ctx.INTEGER().GetText(), 10, 64)
		if err != nil {
			token := ctx.GetStart()
			p.errors = append(p.errors, Error{
				Line:    token.GetLine(),
				Column:  token.GetColumn(),
				Message: err.Error(),
			})
		}

		return &ssql.Tuple{
			Predicate: &ssql.Tuple_Key{
				Key: &ssql.Unary{
					First: &ssql.Operand{
						Value: &ssql.Operand_Int{Int: v},
					},
				},
			},
		}
	}

	var text string
	if ctx.STRING() == nil {
		token := ctx.GetStart()
		p.errors = append(p.errors, Error{
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Message: "missing value",
		})
	} else {
		text = stripQuote(ctx.STRING().GetText())
	}

	return &ssql.Tuple{
		Predicate: &ssql.Tuple_Key{
			Key: &ssql.Unary{
				First: &ssql.Operand{
					Value: &ssql.Operand_Text{
						Text: text,
					},
				},
			},
		},
	}
}
