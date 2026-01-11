package sonar

import (
	"go/ast"

	"github.com/gsixo/gocognit/visitor"
)

type ForVisitor struct {
	parent ast.Visitor
	node   *ast.ForStmt
}

func (v *ForVisitor) VisitInitCondition() {
	if v.hasInitCondition() {
		v.walk(v.node.Init)
	}
}

func (v *ForVisitor) walk(node ast.Node) {
	ast.Walk(v.parent, node)
}

func (v *ForVisitor) hasInitCondition() bool {
	if n := v.node.Init; n != nil {
		return true
	}
	return false
}

func (v *ForVisitor) VisitCondition() {
	v.walk(v.node.Cond)
}

func (v *ForVisitor) VisitPost() {
	if v.hasPost() {
		v.walk(v.node.Post)
	}
}

func (v *ForVisitor) hasPost() bool {
	if n := v.node.Post; n != nil {
		return true
	}
	return false
}

func (v *ForVisitor) VisitBody() {
	v.walk(v.node.Body)
}

type ForVisitorWithCounters struct {
	visitor  visitor.ForVisitor
	counters visitor.VisitorCounters
}

func (v *ForVisitorWithCounters) Visit() (w ast.Visitor) {
	v.counters.IncComplexityCounterWithPlusNestingCounterValue(1)

	v.visitor.VisitInitCondition()
	v.visitor.VisitCondition()
	v.visitor.VisitPost()

	v.counters.IncDecNestingCounterWithFnBetween(
		v.visitor.VisitBody,
	)

	return nil
}
