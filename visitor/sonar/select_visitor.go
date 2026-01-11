package sonar

import (
	"go/ast"

	"github.com/gsixo/gocognit/visitor"
)

type SelectVisitor struct {
	parent ast.Visitor
	node   *ast.SelectStmt
}

func (v *SelectVisitor) VisitBody() {
	v.walk(v.node.Body)
}

func (v *SelectVisitor) walk(node ast.Node) {
	ast.Walk(v.parent, node)
}

type SelectVisitorWithCounters struct {
	visitor  visitor.SelectVisitor
	counters visitor.VisitorCounters
}

func (v *SelectVisitorWithCounters) Visit() (w ast.Visitor) {
	v.counters.IncComplexityCounterWithPlusNestingCounterValue(1)

	v.counters.IncDecNestingCounterWithFnBetween(
		v.visitor.VisitBody,
	)

	return nil
}
