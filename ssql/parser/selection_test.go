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
	"reflect"
	"testing"

	"github.com/chronowave/chronowave/ssql"
)

func TestFind(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"group-by", args{"find group-by($b) where [$b /adf/adf]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b", Group: true}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
					}},
				}}}, false,
		},
		{"avg", args{"find group-by($b), avg($d) where [$b /adf/adf] [$d /df]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{
					{Name: "b", Group: true},
					{Name: "d", Func: &ssql.Function{Name: ssql.Function_AVG}},
				},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"max", args{"find group-by($b), max($d) where [$b /adf/adf] [$d /df]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{
					{Name: "b", Group: true},
					{Name: "d", Func: &ssql.Function{Name: ssql.Function_MAX}},
				},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"min", args{"find group-by($b), min($d) where [$b /adf/adf] [$d /df]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{
					{Name: "b", Group: true},
					{Name: "d", Func: &ssql.Function{Name: ssql.Function_MIN}},
				},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"count", args{"find group-by($b), count($d) where [$b /adf/adf] [$d /df]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{
					{Name: "b", Group: true},
					{Name: "d", Func: &ssql.Function{Name: ssql.Function_COUNT}},
				},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"pctl", args{"find group-by($b), pctl($d, 0.6) where [$b /adf/adf] [$d /df]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{
					{Name: "b", Group: true},
					{Name: "d", Func: &ssql.Function{
						Name:      ssql.Function_PCTL,
						Parameter: &ssql.Function_Double{Double: 0.6},
					}},
				},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"part", args{"find group-by(part($b, 20)), pctl($d, 0.6) where [$b /adf/adf] [$d /df]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{
					{Name: "b", Group: true, Func: &ssql.Function{
						Name:      ssql.Function_PART,
						Parameter: &ssql.Function_Int{Int: 20},
					}},
					{Name: "d", Func: &ssql.Function{
						Name:      ssql.Function_PCTL,
						Parameter: &ssql.Function_Double{Double: 0.6},
					}},
				},
				Where: []*ssql.Expr{
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
