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

func (p *parser) VisitEq(ctx *gen.EqContext) interface{} {
	return &ssql.Tuple{
		Predicate: &ssql.Tuple_Eq{
			Eq: &ssql.Unary{
				First: p.VisitScalar(ctx.Scalar().(*gen.ScalarContext)).(*ssql.Operand),
			},
		},
	}
}

func (p *parser) VisitNeq(ctx *gen.NeqContext) interface{} {
	return &ssql.Tuple{
		Predicate: &ssql.Tuple_Neq{
			Neq: &ssql.Unary{
				First: p.VisitScalar(ctx.Scalar().(*gen.ScalarContext)).(*ssql.Operand),
			},
		},
	}
}

func (p *parser) VisitGt(ctx *gen.GtContext) interface{} {
	return &ssql.Tuple{
		Predicate: &ssql.Tuple_Gt{
			Gt: &ssql.Unary{
				First: p.VisitScalar(ctx.Scalar().(*gen.ScalarContext)).(*ssql.Operand),
			},
		},
	}
}

func (p *parser) VisitGe(ctx *gen.GeContext) interface{} {
	return &ssql.Tuple{
		Predicate: &ssql.Tuple_Ge{
			Ge: &ssql.Unary{
				First: p.VisitScalar(ctx.Scalar().(*gen.ScalarContext)).(*ssql.Operand),
			},
		},
	}
}

func (p *parser) VisitLt(ctx *gen.LtContext) interface{} {
	return &ssql.Tuple{
		Predicate: &ssql.Tuple_Lt{
			Lt: &ssql.Unary{
				First: p.VisitScalar(ctx.Scalar().(*gen.ScalarContext)).(*ssql.Operand),
			},
		},
	}
}

func (p *parser) VisitLe(ctx *gen.LeContext) interface{} {
	return &ssql.Tuple{
		Predicate: &ssql.Tuple_Le{
			Le: &ssql.Unary{
				First: p.VisitScalar(ctx.Scalar().(*gen.ScalarContext)).(*ssql.Operand),
			},
		},
	}
}

func (p *parser) VisitIn(ctx *gen.InContext) interface{} {
	list := ctx.List()
	if list == nil {
		return &ssql.Tuple{
			Predicate: &ssql.Tuple_In{},
		}
	}

	return &ssql.Tuple{
		Predicate: &ssql.Tuple_In{
			In: &ssql.Unary{
				First: p.VisitList(list.(*gen.ListContext)).(*ssql.Operand),
			},
		},
	}
}

func (p *parser) VisitBetween(ctx *gen.BetweenContext) interface{} {
	operands := ctx.AllINTEGER()
	if len(operands) > 0 {
		first, err := strconv.ParseInt(operands[0].GetText(), 10, 64)
		if err != nil {
			// syntax error
		}

		second, err := strconv.ParseInt(operands[1].GetText(), 10, 64)
		if err != nil {
			// TODO: syntax error
		}

		if first >= second {
			// TODO: semantic error
		}

		return &ssql.Tuple{
			Predicate: &ssql.Tuple_Between{
				Between: &ssql.Binary{
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

	operands = ctx.AllREAL_NUMBER()
	if len(operands) > 0 {
		first, err := strconv.ParseFloat(operands[0].GetText(), 64)
		if err != nil {
			// syntax error
		}

		second, err := strconv.ParseFloat(operands[1].GetText(), 64)
		if err != nil {
			// TODO: syntax error
		}

		if first >= second {
			// TODO: semantic error
		}

		return &ssql.Tuple{
			Predicate: &ssql.Tuple_Between{
				Between: &ssql.Binary{
					First: &ssql.Operand{
						Value: &ssql.Operand_Double{
							Double: first,
						},
					},
					Second: &ssql.Operand{
						Value: &ssql.Operand_Double{
							Double: second,
						},
					},
				},
			},
		}
	}

	return &ssql.Tuple{
		Predicate: &ssql.Tuple_Between{
			Between: &ssql.Binary{
				First: &ssql.Operand{
					Value: &ssql.Operand_Int{
						Int: 0,
					},
				},
				Second: &ssql.Operand{
					Value: &ssql.Operand_Int{
						Int: 0,
					},
				},
			},
		},
	}
}

// Visit a parse tree produced by SSQLParser#contain.
func (p *parser) VisitContain(ctx *gen.ContainContext) interface{} {
	var text string
	if ctx.STRING() != nil {
		text = stripQuote(ctx.STRING().GetText())
	}

	return &ssql.Tuple{
		Predicate: &ssql.Tuple_Contain{
			Contain: &ssql.Unary{
				First: &ssql.Operand{
					Value: &ssql.Operand_Text{
						Text: text,
					},
				},
			},
		},
	}
}

// Visit a parse tree produced by SSQLParser#exist.
func (p *parser) VisitExist(ctx *gen.ExistContext) interface{} {
	return &ssql.Tuple{Predicate: &ssql.Tuple_Exist{}}
}

func (p *parser) VisitScalar(ctx *gen.ScalarContext) interface{} {
	if r := ctx.REAL_NUMBER(); r != nil {
		v, err := strconv.ParseFloat(r.GetText(), 64)
		if err != nil {
			token := ctx.GetStart()
			p.errors = append(p.errors, Error{
				Line:    token.GetLine(),
				Column:  token.GetColumn(),
				Message: err.Error(),
			})
		}

		return &ssql.Operand{
			Value: &ssql.Operand_Double{Double: v},
		}
	}

	v, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		token := ctx.GetStart()
		p.errors = append(p.errors, Error{
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Message: err.Error(),
		})
	}

	return &ssql.Operand{
		Value: &ssql.Operand_Int{Int: v},
	}
}

func (p *parser) VisitList(ctx *gen.ListContext) interface{} {
	str := ctx.StringList()
	if str != nil {
		return p.VisitStringList(str.(*gen.StringListContext))
	}

	real := ctx.DoubleList()
	if real != nil {
		return p.VisitDoubleList(real.(*gen.DoubleListContext))
	}

	ints := ctx.IntList()
	if ints != nil {
		return p.VisitIntList(ints.(*gen.IntListContext))
	}

	token := ctx.GetStart()
	p.errors = append(p.errors, Error{
		Line:    token.GetLine(),
		Column:  token.GetColumn(),
		Message: "list must be string, double or integer",
	})

	return nil
}

func (p *parser) VisitStringList(ctx *gen.StringListContext) interface{} {
	texts := make([]string, 0, ctx.GetChildCount())
	for _, n := range ctx.AllSTRING() {
		texts = append(texts, stripQuote(n.GetText()))
	}

	return &ssql.Operand{
		Value: &ssql.Operand_List{
			List: &ssql.List{
				Text: texts,
			},
		},
	}
}

func (p *parser) VisitDoubleList(ctx *gen.DoubleListContext) interface{} {
	doubles := make([]float64, 0, ctx.GetChildCount())
	for _, n := range ctx.AllREAL_NUMBER() {
		v, err := strconv.ParseFloat(n.GetText(), 64)
		if err != nil {
			token := ctx.GetStart()
			p.errors = append(p.errors, Error{
				Line:    token.GetLine(),
				Column:  token.GetColumn(),
				Message: err.Error(),
			})
		}
		doubles = append(doubles, v)
	}

	return &ssql.Operand{
		Value: &ssql.Operand_List{
			List: &ssql.List{
				Double: doubles,
			},
		},
	}
}

func (p *parser) VisitIntList(ctx *gen.IntListContext) interface{} {
	ints := make([]int64, 0, ctx.GetChildCount())
	for _, n := range ctx.AllINTEGER() {
		v, err := strconv.ParseInt(n.GetText(), 10, 64)
		if err != nil {
			token := ctx.GetStart()
			p.errors = append(p.errors, Error{
				Line:    token.GetLine(),
				Column:  token.GetColumn(),
				Message: err.Error(),
			})
		}
		ints = append(ints, v)
	}

	return &ssql.Operand{
		Value: &ssql.Operand_List{
			List: &ssql.List{
				Int: ints,
			},
		},
	}
}

func stripQuote(text string) string {
	if len(text) < 2 {
		return text
	}
	return text[1 : len(text)-1]
}
