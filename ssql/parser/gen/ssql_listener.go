// Code generated from gen/SSQL.g4 by ANTLR 4.8. DO NOT EDIT.

package gen // SSQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SSQLListener is a complete listener for a parse tree produced by SSQLParser.
type SSQLListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterSelection is called when entering the selection production.
	EnterSelection(c *SelectionContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterAggregate is called when entering the aggregate production.
	EnterAggregate(c *AggregateContext)

	// EnterPercentile is called when entering the percentile production.
	EnterPercentile(c *PercentileContext)

	// EnterGroupBy is called when entering the groupBy production.
	EnterGroupBy(c *GroupByContext)

	// EnterPartition is called when entering the partition production.
	EnterPartition(c *PartitionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTuple is called when entering the tuple production.
	EnterTuple(c *TupleContext)

	// EnterVector is called when entering the vector production.
	EnterVector(c *VectorContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterEq is called when entering the eq production.
	EnterEq(c *EqContext)

	// EnterNeq is called when entering the neq production.
	EnterNeq(c *NeqContext)

	// EnterGt is called when entering the gt production.
	EnterGt(c *GtContext)

	// EnterGe is called when entering the ge production.
	EnterGe(c *GeContext)

	// EnterLt is called when entering the lt production.
	EnterLt(c *LtContext)

	// EnterLe is called when entering the le production.
	EnterLe(c *LeContext)

	// EnterIn is called when entering the in production.
	EnterIn(c *InContext)

	// EnterBetween is called when entering the between production.
	EnterBetween(c *BetweenContext)

	// EnterContain is called when entering the contain production.
	EnterContain(c *ContainContext)

	// EnterExist is called when entering the exist production.
	EnterExist(c *ExistContext)

	// EnterTimeframe is called when entering the timeframe production.
	EnterTimeframe(c *TimeframeContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterScalar is called when entering the scalar production.
	EnterScalar(c *ScalarContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterStringList is called when entering the stringList production.
	EnterStringList(c *StringListContext)

	// EnterDoubleList is called when entering the doubleList production.
	EnterDoubleList(c *DoubleListContext)

	// EnterIntList is called when entering the intList production.
	EnterIntList(c *IntListContext)

	// EnterOrderBy is called when entering the orderBy production.
	EnterOrderBy(c *OrderByContext)

	// EnterOrder is called when entering the order production.
	EnterOrder(c *OrderContext)

	// EnterLimit is called when entering the limit production.
	EnterLimit(c *LimitContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitSelection is called when exiting the selection production.
	ExitSelection(c *SelectionContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitAggregate is called when exiting the aggregate production.
	ExitAggregate(c *AggregateContext)

	// ExitPercentile is called when exiting the percentile production.
	ExitPercentile(c *PercentileContext)

	// ExitGroupBy is called when exiting the groupBy production.
	ExitGroupBy(c *GroupByContext)

	// ExitPartition is called when exiting the partition production.
	ExitPartition(c *PartitionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTuple is called when exiting the tuple production.
	ExitTuple(c *TupleContext)

	// ExitVector is called when exiting the vector production.
	ExitVector(c *VectorContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitEq is called when exiting the eq production.
	ExitEq(c *EqContext)

	// ExitNeq is called when exiting the neq production.
	ExitNeq(c *NeqContext)

	// ExitGt is called when exiting the gt production.
	ExitGt(c *GtContext)

	// ExitGe is called when exiting the ge production.
	ExitGe(c *GeContext)

	// ExitLt is called when exiting the lt production.
	ExitLt(c *LtContext)

	// ExitLe is called when exiting the le production.
	ExitLe(c *LeContext)

	// ExitIn is called when exiting the in production.
	ExitIn(c *InContext)

	// ExitBetween is called when exiting the between production.
	ExitBetween(c *BetweenContext)

	// ExitContain is called when exiting the contain production.
	ExitContain(c *ContainContext)

	// ExitExist is called when exiting the exist production.
	ExitExist(c *ExistContext)

	// ExitTimeframe is called when exiting the timeframe production.
	ExitTimeframe(c *TimeframeContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitScalar is called when exiting the scalar production.
	ExitScalar(c *ScalarContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitStringList is called when exiting the stringList production.
	ExitStringList(c *StringListContext)

	// ExitDoubleList is called when exiting the doubleList production.
	ExitDoubleList(c *DoubleListContext)

	// ExitIntList is called when exiting the intList production.
	ExitIntList(c *IntListContext)

	// ExitOrderBy is called when exiting the orderBy production.
	ExitOrderBy(c *OrderByContext)

	// ExitOrder is called when exiting the order production.
	ExitOrder(c *OrderContext)

	// ExitLimit is called when exiting the limit production.
	ExitLimit(c *LimitContext)
}
