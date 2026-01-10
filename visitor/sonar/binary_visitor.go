package sonar

import "go/ast"

type BinaryVisitor struct {
	parent ast.Visitor
	node   *ast.BinaryExpr
}
