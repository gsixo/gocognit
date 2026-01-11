package sonar

import (
	"go/ast"
	"go/token"

	"github.com/gsixo/gocognit/visitor"
)

type BinaryVisitor struct {
	parent ast.Visitor
	node   *ast.BinaryExpr

	calculatedExprs map[ast.Expr]bool
}

// func (v *BinaryVisitor) VisitBinaryOp() {
// 	if isBinaryLogicalOp(v.node.Op) && !v.isCalculated(v.node) {
// 		ops := v.collectBinaryOps(v.node)

// 		var lastOp token.Token
// 		for _, op := range ops {
// 			if lastOp != op {
// 				v.incComplexity(op.String(), n.OpPos)
// 				lastOp = op
// 			}
// 		}
// 	}

// 	return v
// }

func (v *BinaryVisitor) markCalculated(e ast.Expr) {
	if v.calculatedExprs == nil {
		v.calculatedExprs = make(map[ast.Expr]bool)
	}

	v.calculatedExprs[e] = true
}

func (v *BinaryVisitor) isCalculated(e ast.Expr) bool {
	if v.calculatedExprs == nil {
		return false
	}

	return v.calculatedExprs[e]
}

func (v *BinaryVisitor) collectBinaryOps(exp ast.Expr) []token.Token {
	v.markCalculated(exp)

	if exp, ok := exp.(*ast.BinaryExpr); ok {
		return mergeBinaryOps(v.collectBinaryOps(exp.X), exp.Op, v.collectBinaryOps(exp.Y))
	}
	return nil
}

func mergeBinaryOps(x []token.Token, op token.Token, y []token.Token) []token.Token {
	var out []token.Token
	out = append(out, x...)

	if isBinaryLogicalOp(op) {
		out = append(out, op)
	}

	out = append(out, y...)
	return out
}

func isBinaryLogicalOp(op token.Token) bool {
	return op == token.LAND || op == token.LOR
}

type BinaryVisitorWithCounters struct {
	visitor  visitor.BinaryExpressionVisitor
	counters visitor.VisitorCounters
}

func (v *BinaryVisitorWithCounters) Visit() (w ast.Visitor) {
	return nil
}
