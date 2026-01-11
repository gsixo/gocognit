package sonar

import (
	"go/ast"

	"github.com/gsixo/gocognit/visitor"
)

type BranchVisitor struct {
	parent ast.Visitor
	node   *ast.BranchStmt
}

type BranchVisitorWithCounters struct {
	visitor  visitor.BranchStatementVisitor
	counters visitor.VisitorCounters
}

func (v *BranchVisitorWithCounters) Visit() (w ast.Visitor) {
	return nil
}
