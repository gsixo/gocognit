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
	visitor  IfVisitor
	counters visitor.VisitorWithCounters
}

func (v *IfVisitorWithCounters) Visit(node ast.Node) (w ast.Visitor) {
	if v.visitor.IsElseNode() {
		v.counters.IncComplexityCounterWithDelta(1)
	} else {
		w.incComplexityWithNesting()
	}

	v.VisitInitCondition()
	v.VisitCondition()
	w.wrapWithNesting(v.VisitBody)
	v.VisitElseBlockStatement()
	if v.ElseNodeIsBlockStatement() {
		w.incComplexity(1)
	} else {
		v.VisitElseIfStatement()
	}
}
