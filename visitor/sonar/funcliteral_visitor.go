package sonar

import (
	"go/ast"

	"github.com/gsixo/gocognit/visitor"
)

type FuncLiteralVisitor struct {
	parent ast.Visitor
	node   *ast.FuncLit
}

func (v *FuncLiteralVisitor) VisitType() {
	v.walk(v.node.Type)
}

func (v *FuncLiteralVisitor) walk(node ast.Node) {
	ast.Walk(v.parent, node)
}

func (v *FuncLiteralVisitor) VisitBody() {
	v.walk(v.node.Body)
}

type FuncLiteralVisitorWithCounters struct {
	visitor  visitor.FuncLiteralVisitor
	counters visitor.VisitorCounters
}

func (v *FuncLiteralVisitorWithCounters) Visit() (w ast.Visitor) {
	v.visitor.VisitType()

	v.counters.IncDecNestingCounterWithFnBetween(
		v.visitor.VisitBody,
	)

	return nil
}
