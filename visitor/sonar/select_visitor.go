package sonar

import "go/ast"

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
