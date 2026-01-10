package sonar

import "go/ast"

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
