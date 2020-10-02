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
	"strings"

	"github.com/chronowave/chronowave/ssql"
	"github.com/chronowave/chronowave/ssql/parser/gen"
)

func (p *parser) VisitOrderBy(ctx *gen.OrderByContext) interface{} {
	for _, order := range ctx.AllOrder() {
		p.stmt.OrderBy = append(p.stmt.OrderBy, p.VisitOrder(order.(*gen.OrderContext)).(*ssql.OrderBy))
	}
	return nil
}

func (p *parser) VisitOrder(ctx *gen.OrderContext) interface{} {
	orderby := &ssql.OrderBy{
		Name: extractVariableName(ctx.IDENTIFIER().GetText()),
	}

	token := ctx.GetDir()
	if token != nil {
		if strings.EqualFold(token.GetText(), "desc") {
			orderby.Direction = ssql.OrderBy_DESC
		} else {
			orderby.Direction = ssql.OrderBy_ASC
		}
	} else {
		orderby.Direction = ssql.OrderBy_ASC
	}

	return orderby
}

func (p *parser) VisitLimit(ctx *gen.LimitContext) interface{} {
	var t string
	if len(ctx.AllINTEGER()) > 0 {
		t = ctx.AllINTEGER()[0].GetText()
	}
	if v, err := strconv.ParseUint(t, 10, 32); err == nil {
		p.stmt.Limit = uint32(v)
	} else {
		token := ctx.GetStart()
		p.errors = append(p.errors, Error{
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Message: err.Error(),
		})
	}

	return nil
}
