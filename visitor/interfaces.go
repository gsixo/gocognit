package visitor

import "go/ast"

type Visitor interface {
	Visit(node ast.Node) ast.Visitor
	GetComplexity() uint64
}

type VisitorsGetter interface {
	GetIfVisitor(parent ast.Visitor, node *ast.IfStmt) IfVisitor
	GetSwitchVisitor(parent ast.Visitor, node *ast.SwitchStmt) SwitchVisitor
	GetTypeSwitchVisitor(parent ast.Visitor, node *ast.TypeSwitchStmt) TypeSwitchVisitor
	GetSelectVisitor(parent ast.Visitor, node *ast.SelectStmt) SelectVisitor
	GetForVisitor(parent ast.Visitor, node *ast.ForStmt) ForVisitor
	GetRangeVisitor(parent ast.Visitor, node *ast.RangeStmt) RangeVisitor
	GetFuncLiteralVisitor(parent ast.Visitor, node *ast.FuncLit) FuncLiteralVisitor
	GetBranchStatementVisitor(parent ast.Visitor, node *ast.BranchStmt) BranchStatementVisitor
	GetBinaryExpressionVisitor(parent ast.Visitor, node *ast.BinaryExpr) BinaryExpressionVisitor
	GetCallExpressionVisitor(parent ast.Visitor, node *ast.CallExpr) CallExpressionVisitor
}

type VisitorsWithCountersGetter interface {
	GetIfVisitorWithCounters(parent ast.Visitor, node *ast.IfStmt) VisitorWithCounters
	GetSwitchVisitorWithCounters(parent ast.Visitor, node *ast.SwitchStmt) VisitorWithCounters
	GetTypeSwitchVisitorWithCounters(parent ast.Visitor, node *ast.TypeSwitchStmt) VisitorWithCounters
	GetSelectVisitorWithCounters(parent ast.Visitor, node *ast.SelectStmt) VisitorWithCounters
	GetForVisitorWithCounters(parent ast.Visitor, node *ast.ForStmt) VisitorWithCounters
	GetRangeVisitorWithCounters(parent ast.Visitor, node *ast.RangeStmt) VisitorWithCounters
	GetFuncLiteralVisitorWithCounters(parent ast.Visitor, node *ast.FuncLit) VisitorWithCounters
	GetBranchStatementVisitorWithCounters(parent ast.Visitor, node *ast.BranchStmt) VisitorWithCounters
	GetBinaryExpressionVisitorWithCounters(parent ast.Visitor, node *ast.BinaryExpr) VisitorWithCounters
	GetCallExpressionVisitorWithCounters(parent ast.Visitor, node *ast.CallExpr) VisitorWithCounters
}

type VisitorWithCounters interface {
	Visit() (w ast.Visitor)
}

type IfVisitor interface {
	VisitInitCondition()
	VisitCondition()
	VisitBody()
	VisitElseBlockStatement()
	VisitElseIfStatement()
	IsElseNode() bool
	ElseNodeIsBlockStatement() bool
}

type VisitorCounters interface {
	IncComplexityCounterWithDelta(delta uint64)
	DecComplexityCounter()
	LoadComplexityCounter() uint64
	IncComplexityCounterWithPlusNestingCounterValue(delta uint64)
	IncNestingCounterWithDelta(delta uint64)
	DecNestingCounter()
	LoadNestingCounter() uint64
	IncDecNestingCounterWithFnBetween(fn func())
}

type SwitchVisitor interface {
	VisitInitCondition()
	VisitTag()
	VisitBody()
}

type TypeSwitchVisitor interface {
	VisitInitCondition()
	VisitAssign()
	VisitBody()
}

type SelectVisitor interface {
	VisitBody()
}

type ForVisitor interface {
	VisitInitCondition()
	VisitCondition()
	VisitPost()
	VisitBody()
}

type RangeVisitor interface {
	VisitKey()
	VisitValue()
	VisitX()
	VisitBody()
}

type FuncLiteralVisitor interface {
	VisitType()
	VisitBody()
}

type BranchStatementVisitor interface {
}

type BinaryExpressionVisitor interface {
}

type CallExpressionVisitor interface {
	DetectRecursion() bool
}

type Counter interface {
	Inc(delta uint64)
	Dec()
	Load() uint64
}
