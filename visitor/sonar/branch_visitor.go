package sonar

import "go/ast"

type BranchVisitor struct {
	parent ast.Visitor
	node   *ast.BranchStmt
}
