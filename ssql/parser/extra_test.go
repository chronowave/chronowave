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

func TestTimeframe(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"timeframe", args{"find $b where [$b /adf/adf timeframe(2, 3)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Timeframe{
							Timeframe: &ssql.Binary{
								First:  &ssql.Operand{Value: &ssql.Operand_Int{Int: 2}},
								Second: &ssql.Operand{Value: &ssql.Operand_Int{Int: 3}}}}}},
				}}}, false,
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

func TestKey(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"key_int", args{"find $b where [$b /adf/adf key(2)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Key{
							Key: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 2}}}}}},
				}},
			}, false,
		},
		{"key_string", args{"find $b where [$b /adf/adf key('abc')]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{
						Tuple: &ssql.Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &ssql.Tuple_Key{
								Key: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Text{Text: "abc"}}},
							},
						},
					},
				}},
			}, false,
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
