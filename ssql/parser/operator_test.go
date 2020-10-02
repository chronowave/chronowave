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

func TestEqual(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"eq_int", args{"find $b where [$b /adf/adf eq(2)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Eq{
							Eq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 2}}}}}},
				}}}, false,
		},
		{"eq_double", args{"find $b where [$b /adf/adf eq(2.5)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Eq{
							Eq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Double{Double: 2.5}}}}}},
				}}}, false,
		},
		{"eq_string", args{"find $b where [$b /adf/adf eq('abc')]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Eq{
							Eq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 0}}}}}},
				}}}, true,
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

func TestNotEqual(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"neq_int", args{"find $b where [$b /adf/adf neq(2)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Neq{
							Neq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 2}}}}}},
				}}}, false,
		},
		{"neq_double", args{"find $b where [$b /adf/adf neq(2.5)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Neq{
							Neq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Double{Double: 2.5}}}}}},
				}}}, false,
		},
		{"neq_string", args{"find $b where [$b /adf/adf neq('abc')]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Neq{
							Neq: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 0}}}}}},
				}}}, true,
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

func TestGT(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"gt_int", args{"find $b where [$b /adf/adf gt(2)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Gt{
							Gt: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 2}}}}}},
				}}}, false,
		},
		{"gt_double", args{"find $b where [$b /adf/adf gt(2.5)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Gt{
							Gt: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Double{Double: 2.5}}}}}},
				}}}, false,
		},
		{"gt_string", args{"find $b where [$b /adf/adf gt('abc')]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Gt{
							Gt: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 0}}}}}},
				}}}, true,
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

func TestGE(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"ge_int", args{"find $b where [$b /adf/adf ge(2)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Ge{
							Ge: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 2}}}}}},
				}}}, false,
		},
		{"ge_double", args{"find $b where [$b /adf/adf ge(2.5)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Ge{
							Ge: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Double{Double: 2.5}}}}}},
				}}}, false,
		},
		{"ge_string", args{"find $b where [$b /adf/adf ge('abc')]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Ge{
							Ge: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 0}}}}}},
				}}}, true,
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

func TestLT(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"lt_int", args{"find $b where [$b /adf/adf lt(2)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Lt{
							Lt: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 2}}}}}},
				}}}, false,
		},
		{"lt_double", args{"find $b where [$b /adf/adf lt(2.5)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Lt{
							Lt: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Double{Double: 2.5}}}}}},
				}}}, false,
		},
		{"lt_string", args{"find $b where [$b /adf/adf lt('abc')]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Lt{
							Lt: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 0}}}}}},
				}}}, true,
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

func TestLE(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"le_int", args{"find $b where [$b /adf/adf le(2)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Le{
							Le: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 2}}}}}},
				}}}, false,
		},
		{"le_double", args{"find $b where [$b /adf/adf le(2.5)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Le{
							Le: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Double{Double: 2.5}}}}}},
				}}}, false,
		},
		{"le_string", args{"find $b where [$b /adf/adf le('abc')]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Le{
							Le: &ssql.Unary{First: &ssql.Operand{Value: &ssql.Operand_Int{Int: 0}}}}}},
				}}}, true,
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

func TestIN(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"in_int", args{"find $b where [$b /adf/adf in(2, 2)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_In{
							In: &ssql.Unary{First: &ssql.Operand{
								Value: &ssql.Operand_List{List: &ssql.List{Int: []int64{2, 2}}}}}}}},
				}}}, false,
		},
		{"in_double", args{"find $b where [$b /adf/adf in(2.5, 2.0)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_In{
							In: &ssql.Unary{First: &ssql.Operand{
								Value: &ssql.Operand_List{List: &ssql.List{Double: []float64{2.5, 2.0}}}}}}}},
				}}}, false,
		},
		{"in_string", args{"find $b where [$b /adf/adf in('abc', 'adf')]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_In{
							In: &ssql.Unary{First: &ssql.Operand{
								Value: &ssql.Operand_List{List: &ssql.List{Text: []string{"abc", "adf"}}}}}}}},
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

func TestBetween(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"between_int", args{"find $b where [$b /adf/adf between(2, 3)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Between{
							Between: &ssql.Binary{
								First:  &ssql.Operand{Value: &ssql.Operand_Int{Int: 2}},
								Second: &ssql.Operand{Value: &ssql.Operand_Int{Int: 3}}}}}},
				}}}, false,
		},
		{"between_double", args{"find $b where [$b /adf/adf between(2.5, 5.0)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Between{
							Between: &ssql.Binary{
								First:  &ssql.Operand{Value: &ssql.Operand_Double{Double: 2.5}},
								Second: &ssql.Operand{Value: &ssql.Operand_Double{Double: 5.0}}}}}},
				}}}, false,
		},
		{"between_string", args{"find $b where [$b /adf/adf between('abc', 'adf')]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Between{
							Between: &ssql.Binary{
								First:  &ssql.Operand{Value: &ssql.Operand_Int{Int: 0}},
								Second: &ssql.Operand{Value: &ssql.Operand_Int{Int: 0}}}}}},
				}}}, true,
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

func TestContain(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"contain_int", args{"find $b where [$b /adf/adf contain(2)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Contain{
							Contain: &ssql.Unary{
								First: &ssql.Operand{Value: &ssql.Operand_Text{Text: ""}}}}}},
				}}}, true,
		},
		{"contain_double", args{"find $b where [$b /adf/adf contain(5.0)]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Contain{
							Contain: &ssql.Unary{
								First: &ssql.Operand{Value: &ssql.Operand_Text{Text: ""}}}}}},
				}}}, true,
		},
		{"contain_string", args{"find $b where [$b /adf/adf contain('abc')]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name: "b",
						Path: "/adf/adf",
						Predicate: &ssql.Tuple_Contain{
							Contain: &ssql.Unary{
								First: &ssql.Operand{Value: &ssql.Operand_Text{Text: "abc"}}}}}},
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

func TestExist(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssql.Statement
		wantErr bool
	}{
		{"exist ()", args{"find $b where [$b /adf/adf exist()]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name:      "b",
						Path:      "/adf/adf",
						Predicate: &ssql.Tuple_Exist{}}},
				}}}, false,
		},
		{"exist", args{"find $b where [$b /adf/adf exist]"},
			&ssql.Statement{
				Find: []*ssql.Attribute{{Name: "b"}},
				Where: []*ssql.Expr{{
					Field: &ssql.Expr_Tuple{Tuple: &ssql.Tuple{
						Name:      "b",
						Path:      "/adf/adf",
						Predicate: &ssql.Tuple_Exist{}}},
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
