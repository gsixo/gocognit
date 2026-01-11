package sonar

import (
	"go/ast"

	"github.com/gsixo/gocognit/visitor"
)

type SonarVisitor struct {
	visitors visitor.VisitorsWithCountersGetter
	counters visitor.VisitorCounters
}

func NewSonarVisitor() visitor.Visitor {
	return &SonarVisitor{
		visitors: NewSonarVisitorsWithCountersGetter(),
		counters: NewVisitorCounters(),
	}
}

func (v *SonarVisitor) GetComplexity() uint64 {
	return v.counters.LoadComplexityCounter()
}

func (v *SonarVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.IfStmt:
		return v.visitIf(n)
	case *ast.SwitchStmt:
		return v.visitSwitch(n)
	case *ast.TypeSwitchStmt:
		return v.visitTypeSwitchStmt(n)
	case *ast.SelectStmt:
		return v.visitSelectStmt(n)
	case *ast.ForStmt:
		return v.visitForStmt(n)
	case *ast.RangeStmt:
		return v.visitRangeStmt(n)
	case *ast.FuncLit:
		return v.visitFuncLit(n)
	case *ast.BranchStmt:
		return v.visitBranchStmt(n)
	case *ast.BinaryExpr:
		return v.visitBinaryExpr(n)
	case *ast.CallExpr:
		return v.visitCallExpr(n)
	}
	return nil
}

func (w *SonarVisitor) visitIf(node *ast.IfStmt) ast.Visitor {
	return w.visitors.GetIfVisitorWithCounters(w, node).Visit()
}

func (w *SonarVisitor) visitSwitch(node *ast.SwitchStmt) ast.Visitor {
	return w.visitors.GetSwitchVisitorWithCounters(w, node).Visit()
}

func (w *SonarVisitor) visitTypeSwitchStmt(node *ast.TypeSwitchStmt) ast.Visitor {
	return w.visitors.GetTypeSwitchVisitorWithCounters(w, node).Visit()
}

func (w *SonarVisitor) visitSelectStmt(node *ast.SelectStmt) ast.Visitor {
	return w.visitors.GetSelectVisitorWithCounters(w, node).Visit()
}

func (w *SonarVisitor) visitForStmt(node *ast.ForStmt) ast.Visitor {
	return w.visitors.GetForVisitorWithCounters(w, node).Visit()
}

func (w *SonarVisitor) visitRangeStmt(node *ast.RangeStmt) ast.Visitor {
	return w.visitors.GetRangeVisitorWithCounters(w, node).Visit()
}

func (w *SonarVisitor) visitFuncLit(node *ast.FuncLit) ast.Visitor {
	return w.visitors.GetFuncLiteralVisitorWithCounters(w, node).Visit()
}

func (w *SonarVisitor) visitBranchStmt(node *ast.BranchStmt) ast.Visitor {
	return w.visitors.GetBranchStatementVisitorWithCounters(w, node).Visit()
}

func (w *SonarVisitor) visitBinaryExpr(node *ast.BinaryExpr) ast.Visitor {
	return w.visitors.GetBinaryExpressionVisitorWithCounters(w, node).Visit()
}

func (w *SonarVisitor) visitCallExpr(node *ast.CallExpr) ast.Visitor {
	return w.visitors.GetCallExpressionVisitorWithCounters(w, node).Visit()
}
