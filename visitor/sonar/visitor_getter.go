package sonar

import (
	"go/ast"

	"github.com/gsixo/gocognit/visitor"
)

type SonarVisitorsGetter struct{}

func (g *SonarVisitorsGetter) GetIfVisitor(parent ast.Visitor, node *ast.IfStmt) visitor.IfVisitor {
	return &IfVisitor{parent: parent, node: node}
}

func (g *SonarVisitorsGetter) GetSwitchVisitor(parent ast.Visitor, node *ast.SwitchStmt) visitor.SwitchVisitor {
	return &SwitchVisitor{parent: parent, node: node}
}

func (g *SonarVisitorsGetter) GetTypeSwitchVisitor(parent ast.Visitor, node *ast.TypeSwitchStmt) visitor.TypeSwitchVisitor {
	return &TypeSwitchVisitor{parent: parent, node: node}
}

func (g *SonarVisitorsGetter) GetSelectVisitor(parent ast.Visitor, node *ast.SelectStmt) visitor.SelectVisitor {
	return &SelectVisitor{parent: parent, node: node}
}

func (g *SonarVisitorsGetter) GetForVisitor(parent ast.Visitor, node *ast.ForStmt) visitor.ForVisitor {
	return &ForVisitor{parent: parent, node: node}
}

func (g *SonarVisitorsGetter) GetRangeVisitor(parent ast.Visitor, node *ast.RangeStmt) visitor.RangeVisitor {
	return &RangeVisitor{parent: parent, node: node}
}

func (g *SonarVisitorsGetter) GetFuncLiteralVisitor(parent ast.Visitor, node *ast.FuncLit) visitor.FuncLiteralVisitor {
	return &FuncLiteralVisitor{parent: parent, node: node}
}

func (g *SonarVisitorsGetter) GetBranchStatementVisitor(parent ast.Visitor, node *ast.BranchStmt) visitor.BranchStatementVisitor {
	return &BranchVisitor{parent: parent, node: node}
}

func (g *SonarVisitorsGetter) GetBinaryExpressionVisitor(parent ast.Visitor, node *ast.BinaryExpr) visitor.BinaryExpressionVisitor {
	return &BinaryVisitor{parent: parent, node: node}
}

func (g *SonarVisitorsGetter) GetCallExpressionVisitor(parent ast.Visitor, node *ast.CallExpr) visitor.CallExpressionVisitor {
	return &CallVisitor{parent: parent, node: node}
}
