package sonar

import (
	"go/ast"

	"github.com/gsixo/gocognit/visitor"
)

type SwitchVisitor struct {
	parent ast.Visitor
	node   *ast.SwitchStmt
}

func (v *SwitchVisitor) VisitInitCondition() {
	if v.nodeHasInitCondition() {
		v.walk(v.node.Init)
	}
}

func (v *SwitchVisitor) walk(node ast.Node) {
	ast.Walk(v.parent, node)
}

func (v *SwitchVisitor) nodeHasInitCondition() bool {
	if n := v.node.Init; n != nil {
		return true
	}
	return false
}

func (v *SwitchVisitor) VisitTag() {
	if v.nodeHasTag() {
		v.walk(v.node.Tag)
	}
}

func (v *SwitchVisitor) nodeHasTag() bool {
	if n := v.node.Tag; n != nil {
		return true
	}
	return false
}

func (v *SwitchVisitor) VisitBody() {
	v.walk(v.node.Body)
}

type SwitchVisitorWithCounters struct {
	visitor  visitor.SwitchVisitor
	counters visitor.VisitorCounters
}

func (v *SwitchVisitorWithCounters) Visit() (w ast.Visitor) {
	v.counters.IncComplexityCounterWithPlusNestingCounterValue(1)
	v.visitor.VisitInitCondition()
	v.visitor.VisitTag()

	v.counters.IncDecNestingCounterWithFnBetween(
		v.visitor.VisitBody,
	)

	return nil
}
