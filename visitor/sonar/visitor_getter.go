package sonar

import (
	"go/ast"

	"github.com/gsixo/gocognit/visitor"
)

type SonarVisitorsGetter struct{}

func NewSonarVisitorsGetter() visitor.VisitorsGetter {
	return &SonarVisitorsGetter{}
}

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

type SonarVisitorsWithCountersGetter struct {
	visitorsGetter visitor.VisitorsGetter
	counters       visitor.VisitorCounters
}

func NewSonarVisitorsWithCountersGetter() visitor.VisitorsWithCountersGetter {
	return &SonarVisitorsWithCountersGetter{
		visitorsGetter: NewSonarVisitorsGetter(),
		counters:       NewVisitorCounters(),
	}
}

func (g *SonarVisitorsWithCountersGetter) GetIfVisitorWithCounters(parent ast.Visitor, node *ast.IfStmt) visitor.VisitorWithCounters {
	visitor := g.visitorsGetter.GetIfVisitor(parent, node)
	return &IfVisitorWithCounters{
		visitor:  visitor,
		counters: g.counters,
	}
}

func (g *SonarVisitorsWithCountersGetter) GetSwitchVisitorWithCounters(parent ast.Visitor, node *ast.SwitchStmt) visitor.VisitorWithCounters {
	visitor := g.visitorsGetter.GetSwitchVisitor(parent, node)
	return &SwitchVisitorWithCounters{
		visitor:  visitor,
		counters: g.counters,
	}
}

func (g *SonarVisitorsWithCountersGetter) GetTypeSwitchVisitorWithCounters(parent ast.Visitor, node *ast.TypeSwitchStmt) visitor.VisitorWithCounters {
	visitor := g.visitorsGetter.GetTypeSwitchVisitor(parent, node)
	return &TypeSwitchVisitorWithCounters{
		visitor:  visitor,
		counters: g.counters,
	}
}

func (g *SonarVisitorsWithCountersGetter) GetSelectVisitorWithCounters(parent ast.Visitor, node *ast.SelectStmt) visitor.VisitorWithCounters {
	visitor := g.visitorsGetter.GetSelectVisitor(parent, node)
	return &SelectVisitorWithCounters{
		visitor:  visitor,
		counters: g.counters,
	}
}

func (g *SonarVisitorsWithCountersGetter) GetForVisitorWithCounters(parent ast.Visitor, node *ast.ForStmt) visitor.VisitorWithCounters {
	visitor := g.visitorsGetter.GetForVisitor(parent, node)
	return &ForVisitorWithCounters{
		visitor:  visitor,
		counters: g.counters,
	}
}

func (g *SonarVisitorsWithCountersGetter) GetRangeVisitorWithCounters(parent ast.Visitor, node *ast.RangeStmt) visitor.VisitorWithCounters {
	visitor := g.visitorsGetter.GetRangeVisitor(parent, node)
	return &RangeVisitorWithCounters{
		visitor:  visitor,
		counters: g.counters,
	}
}

func (g *SonarVisitorsWithCountersGetter) GetFuncLiteralVisitorWithCounters(parent ast.Visitor, node *ast.FuncLit) visitor.VisitorWithCounters {
	visitor := g.visitorsGetter.GetFuncLiteralVisitor(parent, node)
	return &FuncLiteralVisitorWithCounters{
		visitor:  visitor,
		counters: g.counters,
	}
}

func (g *SonarVisitorsWithCountersGetter) GetBranchStatementVisitorWithCounters(parent ast.Visitor, node *ast.BranchStmt) visitor.VisitorWithCounters {
	visitor := g.visitorsGetter.GetBranchStatementVisitor(parent, node)
	return &BranchVisitorWithCounters{
		visitor:  visitor,
		counters: g.counters,
	}
}

func (g *SonarVisitorsWithCountersGetter) GetBinaryExpressionVisitorWithCounters(parent ast.Visitor, node *ast.BinaryExpr) visitor.VisitorWithCounters {
	visitor := g.visitorsGetter.GetBinaryExpressionVisitor(parent, node)
	return &BinaryVisitorWithCounters{
		visitor:  visitor,
		counters: g.counters,
	}
}

func (g *SonarVisitorsWithCountersGetter) GetCallExpressionVisitorWithCounters(parent ast.Visitor, node *ast.CallExpr) visitor.VisitorWithCounters {
	visitor := g.visitorsGetter.GetCallExpressionVisitor(parent, node)
	return &BinaryVisitorWithCounters{
		visitor:  visitor,
		counters: g.counters,
	}
}
