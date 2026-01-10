package visitor

import "go/ast"

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

type IfVisitor interface {
	VisitInitCondition()
	VisitCondition()
	VisitBody()
	VisitElseBlockStatement()
	VisitElseIfStatement()
	IsElseNode() bool
	ElseNodeIsBlockStatement() bool
}

type VisitorWithCounters interface {
	IncComplexityCounterWithDelta(delta uint64)
	DecComplexityCounter()
	LoadComplexityCounter() uint64
	IncNestingCounterWithDelta(delta uint64)
	DecNestingCounter()
	LoadNestingCounter() uint64
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
