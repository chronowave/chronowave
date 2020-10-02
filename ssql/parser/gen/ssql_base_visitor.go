// Code generated from gen/SSQL.g4 by ANTLR 4.8. DO NOT EDIT.

package gen // SSQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSSQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSSQLVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitSelection(ctx *SelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitAggregate(ctx *AggregateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitPercentile(ctx *PercentileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitGroupBy(ctx *GroupByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitPartition(ctx *PartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitVector(ctx *VectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitOr(ctx *OrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitAnd(ctx *AndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitEq(ctx *EqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitNeq(ctx *NeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitGt(ctx *GtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitGe(ctx *GeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitLt(ctx *LtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitLe(ctx *LeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitIn(ctx *InContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitBetween(ctx *BetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitContain(ctx *ContainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitExist(ctx *ExistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitTimeframe(ctx *TimeframeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitScalar(ctx *ScalarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitStringList(ctx *StringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitDoubleList(ctx *DoubleListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitIntList(ctx *IntListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitOrderBy(ctx *OrderByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitOrder(ctx *OrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSSQLVisitor) VisitLimit(ctx *LimitContext) interface{} {
	return v.VisitChildren(ctx)
}
