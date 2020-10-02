// Code generated from gen/SSQL.g4 by ANTLR 4.8. DO NOT EDIT.

package gen // SSQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SSQLParser.
type SSQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SSQLParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by SSQLParser#selection.
	VisitSelection(ctx *SelectionContext) interface{}

	// Visit a parse tree produced by SSQLParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by SSQLParser#aggregate.
	VisitAggregate(ctx *AggregateContext) interface{}

	// Visit a parse tree produced by SSQLParser#percentile.
	VisitPercentile(ctx *PercentileContext) interface{}

	// Visit a parse tree produced by SSQLParser#groupBy.
	VisitGroupBy(ctx *GroupByContext) interface{}

	// Visit a parse tree produced by SSQLParser#partition.
	VisitPartition(ctx *PartitionContext) interface{}

	// Visit a parse tree produced by SSQLParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SSQLParser#tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by SSQLParser#vector.
	VisitVector(ctx *VectorContext) interface{}

	// Visit a parse tree produced by SSQLParser#or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by SSQLParser#and.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by SSQLParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by SSQLParser#eq.
	VisitEq(ctx *EqContext) interface{}

	// Visit a parse tree produced by SSQLParser#neq.
	VisitNeq(ctx *NeqContext) interface{}

	// Visit a parse tree produced by SSQLParser#gt.
	VisitGt(ctx *GtContext) interface{}

	// Visit a parse tree produced by SSQLParser#ge.
	VisitGe(ctx *GeContext) interface{}

	// Visit a parse tree produced by SSQLParser#lt.
	VisitLt(ctx *LtContext) interface{}

	// Visit a parse tree produced by SSQLParser#le.
	VisitLe(ctx *LeContext) interface{}

	// Visit a parse tree produced by SSQLParser#in.
	VisitIn(ctx *InContext) interface{}

	// Visit a parse tree produced by SSQLParser#between.
	VisitBetween(ctx *BetweenContext) interface{}

	// Visit a parse tree produced by SSQLParser#contain.
	VisitContain(ctx *ContainContext) interface{}

	// Visit a parse tree produced by SSQLParser#exist.
	VisitExist(ctx *ExistContext) interface{}

	// Visit a parse tree produced by SSQLParser#timeframe.
	VisitTimeframe(ctx *TimeframeContext) interface{}

	// Visit a parse tree produced by SSQLParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by SSQLParser#scalar.
	VisitScalar(ctx *ScalarContext) interface{}

	// Visit a parse tree produced by SSQLParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by SSQLParser#stringList.
	VisitStringList(ctx *StringListContext) interface{}

	// Visit a parse tree produced by SSQLParser#doubleList.
	VisitDoubleList(ctx *DoubleListContext) interface{}

	// Visit a parse tree produced by SSQLParser#intList.
	VisitIntList(ctx *IntListContext) interface{}

	// Visit a parse tree produced by SSQLParser#orderBy.
	VisitOrderBy(ctx *OrderByContext) interface{}

	// Visit a parse tree produced by SSQLParser#order.
	VisitOrder(ctx *OrderContext) interface{}

	// Visit a parse tree produced by SSQLParser#limit.
	VisitLimit(ctx *LimitContext) interface{}
}
