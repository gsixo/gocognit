package sonar

import "go/ast"

type CallVisitor struct {
	parent ast.Visitor
	node   *ast.CallExpr
	name   *ast.Ident
}

func (v *CallVisitor) DetectRecursion() bool {
	callIdentificator := v.getFunIdentificator()

	obj, name := callIdentificator.Obj, callIdentificator.Name

	if obj == v.name.Obj && name == v.name.Name {
		return true
	}

	return false
}

func (v *CallVisitor) nodeFunIsIdentificator() bool {
	if _, ok := v.node.Fun.(*ast.Ident); ok {
		return true
	}
	return false
}

func (v *CallVisitor) getFunIdentificator() *ast.Ident {
	if v.nodeFunIsIdentificator() {
		return v.node.Fun.(*ast.Ident)
	}
	return nil
}
