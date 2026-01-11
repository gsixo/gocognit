package sonar

import (
	"go/ast"

	"github.com/gsixo/gocognit/visitor"
)

type IfVisitor struct {
	parent    ast.Visitor
	node      *ast.IfStmt
	elseNodes map[*ast.IfStmt]bool
}

func (v *IfVisitor) VisitInitCondition() {
	if v.nodeHasInitCondition() {
		v.walk(v.node.Init)
	}
}

func (v *IfVisitor) walk(node ast.Node) {
	ast.Walk(v.parent, node)
}

func (v *IfVisitor) nodeHasInitCondition() bool {
	if n := v.node.Init; n != nil {
		return true
	}
	return false
}

func (v *IfVisitor) VisitCondition() {
	v.walk(v.node.Cond)
}

func (v *IfVisitor) VisitBody() {
	v.walk(v.node.Body)
}

func (v *IfVisitor) VisitElseBlockStatement() {
	switch v.node.Else.(type) {
	case *ast.BlockStmt:
		v.walk(v.node.Else)
	}
}

func (v *IfVisitor) VisitElseIfStatement() {
	switch v.node.Else.(type) {
	case *ast.IfStmt:
		v.markAsElseNode(v.node)
		v.walk(v.node.Else)
	}
}

func (v *IfVisitor) VisitElseBlock() {
	switch v.node.Else.(type) {
	case *ast.BlockStmt:
		v.walk(v.node.Else)
	}
}

func (v *IfVisitor) IsElseNode() bool {
	if v.elseNodes == nil {
		return false
	}

	return v.elseNodes[v.node]
}

func (v *IfVisitor) markAsElseNode(n *ast.IfStmt) {
	if v.elseNodes == nil {
		v.elseNodes = make(map[*ast.IfStmt]bool)
	}

	v.elseNodes[n] = true
}

func (v *IfVisitor) ElseNodeIsBlockStatement() bool {
	switch v.node.Else.(type) {
	case *ast.BlockStmt:
		return true
	default:
		return false
	}
}

type IfVisitorWithCounters struct {
	visitor  visitor.IfVisitor
	counters visitor.VisitorCounters
}

func (v *IfVisitorWithCounters) Visit() (w ast.Visitor) {
	if v.visitor.IsElseNode() {
		v.counters.IncComplexityCounterWithDelta(1)
	} else {
		v.counters.IncComplexityCounterWithPlusNestingCounterValue(1)
	}

	v.visitor.VisitInitCondition()
	v.visitor.VisitCondition()
	v.counters.IncDecNestingCounterWithFnBetween(
		v.visitor.VisitBody,
	)
	v.visitor.VisitElseBlockStatement()
	if v.visitor.ElseNodeIsBlockStatement() {
		v.counters.IncComplexityCounterWithDelta(1)
	} else {
		v.visitor.VisitElseIfStatement()
	}

	return nil
}
