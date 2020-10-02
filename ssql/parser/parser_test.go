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

//go:generate protoc --go_out=../../ --proto_path=../ proto/ssql.proto
//go:generate ./gen/antlr4/gen -Dlanguage=Go -visitor -long-messages -package gen gen/SSQL.g4

import (
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssql"
)

func TestParse(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want *ssql.Statement
	}{
		{
			"complete",
			args{`find group-by(part($c,10),$b), avg($g)
                              where [$b /adf/adf eq(10)]
                                    [$g /dkf/adf]
                                    [/adf/dkf [/df between(15, 20)]
                                              [/adf eq(2.5)]
                                    ]`},
			&ssql.Statement{
				Find: []*ssql.Attribute{
					{
						Name:  "c",
						Group: true,
						Func:  &ssql.Function{Name: ssql.Function_PART, Parameter: &ssql.Function_Int{Int: 10}},
					},
					{
						Name:  "b",
						Group: true,
					},
					{
						Name: "g",
						Func: &ssql.Function{Name: ssql.Function_AVG},
					},
				},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &ssql.Tuple_Eq{
								Eq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 10}}}}}},
					},
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "g",
							Path: "/dkf/adf"}},
					},
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Path: "/adf/dkf",
							Predicate: &ssql.Tuple_Nested{
								Nested: &ssql.Nested{
									Expr: []*ssql.Expr{
										{
											Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
												Path: "/df",
												Predicate: &ssql.Tuple_Between{
													Between: &ssql.Binary{
														First:  &ssql.Operand{Value: &ssql.Operand_Int{Int: 15}},
														Second: &ssql.Operand{Value: &ssql.Operand_Int{Int: 20}}},
												}},
											},
										},
										{
											Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
												Path: "/adf",
												Predicate: &ssql.Tuple_Eq{
													Eq: &ssql.Unary{First: &ssql.Operand{
														Value: &ssql.Operand_Double{Double: 2.5}},
													}}},
											},
										},
									}},
							},
						},
						},
					}}},
		},
		{
			"zero_int",
			args{`find $g where [$g /adf/adf eq(0)]`},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "g"}},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "g",
							Path: "/adf/adf",
							Predicate: &ssql.Tuple_Eq{
								Eq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 0}}}}}},
					}},
			},
		},
		{
			"zeros_int",
			args{`find $g where [$g /adf/adf eq(000)]`},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "g"}},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "g",
							Path: "/adf/adf",
							Predicate: &ssql.Tuple_Eq{
								Eq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 0}}}}}},
					}},
			},
		},
		{
			"zero_double",
			args{`find $g where [$g /adf/adf eq(0.0)]`},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "g"}},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "g",
							Path: "/adf/adf",
							Predicate: &ssql.Tuple_Eq{
								Eq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Double{Double: 0}}}}}},
					}},
			},
		},
		{
			"nested",
			args{`find $g where [$r / [/adf/adf eq(0.0)]]`},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "g"}},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "r",
							Path: "/",
							Predicate: &ssql.Tuple_Nested{
								Nested: &ssql.Nested{
									Expr: []*ssql.Expr{{
										Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
											Path: "/adf/adf",
											Predicate: &ssql.Tuple_Eq{
												Eq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Double{Double: 0}}}}}},
									}},
								},
							}}},
					},
				},
			},
		},
		{
			"or",
			args{`find $g where {[/adf/adf eq(0.0)] [/d lt(1)]}`},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "g"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Or{Or: &ssql.OR{
						Expr: []*ssql.Expr{
							{
								Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
									Path: "/adf/adf",
									Predicate: &ssql.Tuple_Eq{
										Eq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Double{Double: 0}}},
									},
								}},
							},
							{
								Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
									Path: "/d",
									Predicate: &ssql.Tuple_Lt{
										Lt: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 1}}},
									},
								}},
							},
						}},
					},
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if len(errs) > 0 {
				t.Errorf("parsing errors %v", errs)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
