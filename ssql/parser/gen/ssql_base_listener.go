// Code generated from gen/SSQL.g4 by ANTLR 4.8. DO NOT EDIT.

package gen // SSQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSSQLListener is a complete listener for a parse tree produced by SSQLParser.
type BaseSSQLListener struct{}

var _ SSQLListener = &BaseSSQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSSQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSSQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSSQLListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSSQLListener) ExitStart(ctx *StartContext) {}

// EnterSelection is called when production selection is entered.
func (s *BaseSSQLListener) EnterSelection(ctx *SelectionContext) {}

// ExitSelection is called when production selection is exited.
func (s *BaseSSQLListener) ExitSelection(ctx *SelectionContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseSSQLListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseSSQLListener) ExitAttribute(ctx *AttributeContext) {}

// EnterAggregate is called when production aggregate is entered.
func (s *BaseSSQLListener) EnterAggregate(ctx *AggregateContext) {}

// ExitAggregate is called when production aggregate is exited.
func (s *BaseSSQLListener) ExitAggregate(ctx *AggregateContext) {}

// EnterPercentile is called when production percentile is entered.
func (s *BaseSSQLListener) EnterPercentile(ctx *PercentileContext) {}

// ExitPercentile is called when production percentile is exited.
func (s *BaseSSQLListener) ExitPercentile(ctx *PercentileContext) {}

// EnterGroupBy is called when production groupBy is entered.
func (s *BaseSSQLListener) EnterGroupBy(ctx *GroupByContext) {}

// ExitGroupBy is called when production groupBy is exited.
func (s *BaseSSQLListener) ExitGroupBy(ctx *GroupByContext) {}

// EnterPartition is called when production partition is entered.
func (s *BaseSSQLListener) EnterPartition(ctx *PartitionContext) {}

// ExitPartition is called when production partition is exited.
func (s *BaseSSQLListener) ExitPartition(ctx *PartitionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSSQLListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSSQLListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTuple is called when production tuple is entered.
func (s *BaseSSQLListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production tuple is exited.
func (s *BaseSSQLListener) ExitTuple(ctx *TupleContext) {}

// EnterVector is called when production vector is entered.
func (s *BaseSSQLListener) EnterVector(ctx *VectorContext) {}

// ExitVector is called when production vector is exited.
func (s *BaseSSQLListener) ExitVector(ctx *VectorContext) {}

// EnterOr is called when production or is entered.
func (s *BaseSSQLListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BaseSSQLListener) ExitOr(ctx *OrContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseSSQLListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseSSQLListener) ExitAnd(ctx *AndContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseSSQLListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseSSQLListener) ExitPredicate(ctx *PredicateContext) {}

// EnterEq is called when production eq is entered.
func (s *BaseSSQLListener) EnterEq(ctx *EqContext) {}

// ExitEq is called when production eq is exited.
func (s *BaseSSQLListener) ExitEq(ctx *EqContext) {}

// EnterNeq is called when production neq is entered.
func (s *BaseSSQLListener) EnterNeq(ctx *NeqContext) {}

// ExitNeq is called when production neq is exited.
func (s *BaseSSQLListener) ExitNeq(ctx *NeqContext) {}

// EnterGt is called when production gt is entered.
func (s *BaseSSQLListener) EnterGt(ctx *GtContext) {}

// ExitGt is called when production gt is exited.
func (s *BaseSSQLListener) ExitGt(ctx *GtContext) {}

// EnterGe is called when production ge is entered.
func (s *BaseSSQLListener) EnterGe(ctx *GeContext) {}

// ExitGe is called when production ge is exited.
func (s *BaseSSQLListener) ExitGe(ctx *GeContext) {}

// EnterLt is called when production lt is entered.
func (s *BaseSSQLListener) EnterLt(ctx *LtContext) {}

// ExitLt is called when production lt is exited.
func (s *BaseSSQLListener) ExitLt(ctx *LtContext) {}

// EnterLe is called when production le is entered.
func (s *BaseSSQLListener) EnterLe(ctx *LeContext) {}

// ExitLe is called when production le is exited.
func (s *BaseSSQLListener) ExitLe(ctx *LeContext) {}

// EnterIn is called when production in is entered.
func (s *BaseSSQLListener) EnterIn(ctx *InContext) {}

// ExitIn is called when production in is exited.
func (s *BaseSSQLListener) ExitIn(ctx *InContext) {}

// EnterBetween is called when production between is entered.
func (s *BaseSSQLListener) EnterBetween(ctx *BetweenContext) {}

// ExitBetween is called when production between is exited.
func (s *BaseSSQLListener) ExitBetween(ctx *BetweenContext) {}

// EnterContain is called when production contain is entered.
func (s *BaseSSQLListener) EnterContain(ctx *ContainContext) {}

// ExitContain is called when production contain is exited.
func (s *BaseSSQLListener) ExitContain(ctx *ContainContext) {}

// EnterExist is called when production exist is entered.
func (s *BaseSSQLListener) EnterExist(ctx *ExistContext) {}

// ExitExist is called when production exist is exited.
func (s *BaseSSQLListener) ExitExist(ctx *ExistContext) {}

// EnterTimeframe is called when production timeframe is entered.
func (s *BaseSSQLListener) EnterTimeframe(ctx *TimeframeContext) {}

// ExitTimeframe is called when production timeframe is exited.
func (s *BaseSSQLListener) ExitTimeframe(ctx *TimeframeContext) {}

// EnterKey is called when production key is entered.
func (s *BaseSSQLListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseSSQLListener) ExitKey(ctx *KeyContext) {}

// EnterScalar is called when production scalar is entered.
func (s *BaseSSQLListener) EnterScalar(ctx *ScalarContext) {}

// ExitScalar is called when production scalar is exited.
func (s *BaseSSQLListener) ExitScalar(ctx *ScalarContext) {}

// EnterList is called when production list is entered.
func (s *BaseSSQLListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseSSQLListener) ExitList(ctx *ListContext) {}

// EnterStringList is called when production stringList is entered.
func (s *BaseSSQLListener) EnterStringList(ctx *StringListContext) {}

// ExitStringList is called when production stringList is exited.
func (s *BaseSSQLListener) ExitStringList(ctx *StringListContext) {}

// EnterDoubleList is called when production doubleList is entered.
func (s *BaseSSQLListener) EnterDoubleList(ctx *DoubleListContext) {}

// ExitDoubleList is called when production doubleList is exited.
func (s *BaseSSQLListener) ExitDoubleList(ctx *DoubleListContext) {}

// EnterIntList is called when production intList is entered.
func (s *BaseSSQLListener) EnterIntList(ctx *IntListContext) {}

// ExitIntList is called when production intList is exited.
func (s *BaseSSQLListener) ExitIntList(ctx *IntListContext) {}

// EnterOrderBy is called when production orderBy is entered.
func (s *BaseSSQLListener) EnterOrderBy(ctx *OrderByContext) {}

// ExitOrderBy is called when production orderBy is exited.
func (s *BaseSSQLListener) ExitOrderBy(ctx *OrderByContext) {}

// EnterOrder is called when production order is entered.
func (s *BaseSSQLListener) EnterOrder(ctx *OrderContext) {}

// ExitOrder is called when production order is exited.
func (s *BaseSSQLListener) ExitOrder(ctx *OrderContext) {}

// EnterLimit is called when production limit is entered.
func (s *BaseSSQLListener) EnterLimit(ctx *LimitContext) {}

// ExitLimit is called when production limit is exited.
func (s *BaseSSQLListener) ExitLimit(ctx *LimitContext) {}
