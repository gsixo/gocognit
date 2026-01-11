package sonar

import (
	"go/ast"

	"github.com/gsixo/gocognit/visitor"
)

type RangeVisitor struct {
	parent ast.Visitor
	node   *ast.RangeStmt
}

func (v *RangeVisitor) VisitKey() {
	if v.nodeHasKey() {
		v.walk(v.node.Key)
	}
}

func (v *RangeVisitor) walk(node ast.Node) {
	ast.Walk(v.parent, node)
}

func (v *RangeVisitor) nodeHasKey() bool {
	if n := v.node.Key; n != nil {
		return true
	}
	return false
}

func (v *RangeVisitor) VisitValue() {
	if v.nodeHasValue() {
		v.walk(v.node.Value)
	}
}

func (v *RangeVisitor) nodeHasValue() bool {
	if n := v.node.Value; n != nil {
		return true
	}
	return false
}

func (v *RangeVisitor) VisitX() {
	v.walk(v.node.X)
}

func (v *RangeVisitor) VisitBody() {
	v.walk(v.node.Body)
}

type RangeVisitorWithCounters struct {
	visitor  visitor.RangeVisitor
	counters visitor.VisitorCounters
}

func (v *RangeVisitorWithCounters) Visit() (w ast.Visitor) {
	v.counters.IncComplexityCounterWithPlusNestingCounterValue(1)
	v.visitor.VisitKey()
	v.visitor.VisitValue()
	v.visitor.VisitX()

	v.counters.IncDecNestingCounterWithFnBetween(
		v.visitor.VisitBody,
	)

	return nil
}
