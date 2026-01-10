package sonar

import "go/ast"

type TypeSwitchVisitor struct {
	parent ast.Visitor
	node   *ast.TypeSwitchStmt
}

func (v *TypeSwitchVisitor) VisitInitCondition() {
	if v.nodeHasInitCondition() {
		v.walk(v.node.Init)
	}
}

func (v *TypeSwitchVisitor) walk(node ast.Node) {
	ast.Walk(v.parent, node)
}

func (v *TypeSwitchVisitor) nodeHasInitCondition() bool {
	if n := v.node.Init; n != nil {
		return true
	}
	return false
}

func (v *TypeSwitchVisitor) VisitAssign() {
	if v.nodeHasAssign() {
		v.walk(v.node.Assign)
	}
}

func (v *TypeSwitchVisitor) nodeHasAssign() bool {
	if n := v.node.Assign; n != nil {
		return true
	}
	return false
}

func (v *TypeSwitchVisitor) VisitBody() {
	v.walk(v.node.Body)
}
