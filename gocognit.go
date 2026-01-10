package gocognit

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"

	"github.com/gsixo/gocognit/visitor/sonar"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Stat is statistic of the complexity.
type Stat struct {
	PkgName     string
	FuncName    string
	Complexity  int
	Pos         token.Position
	Diagnostics []Diagnostic `json:",omitempty"`
}

// Diagnostic contains information how the complexity increase.
type Diagnostic struct {
	Inc     int
	Nesting int `json:",omitempty"`
	Text    string
	Pos     DiagnosticPosition
}

// DiagnosticPosition is the position of the diagnostic.
type DiagnosticPosition struct {
	Offset int // offset, starting at 0
	Line   int // line number, starting at 1
	Column int // column number, starting at 1 (byte count)
}

func (pos DiagnosticPosition) isValid() bool {
	return pos.Line > 0
}

func (pos DiagnosticPosition) String() string {
	var s string
	if pos.isValid() {
		if s != "" {
			s += ":"
		}

		s += strconv.Itoa(pos.Line)
		if pos.Column != 0 {
			s += fmt.Sprintf(":%d", pos.Column)
		}
	}

	if s == "" {
		s = "-"
	}

	return s
}

func (d Diagnostic) String() string {
	if d.Nesting == 0 {
		return fmt.Sprintf("+%d", d.Inc)
	}

	return fmt.Sprintf("+%d (nesting=%d)", d.Inc, d.Nesting)
}

func (s Stat) String() string {
	return fmt.Sprintf("%d %s %s %s", s.Complexity, s.PkgName, s.FuncName, s.Pos)
}

// ComplexityStats builds the complexity statistics.
func ComplexityStats(f *ast.File, fset *token.FileSet, stats []Stat) []Stat {
	return ComplexityStatsWithDiagnostic(f, fset, stats, false)
}

// ComplexityStatsWithDiagnostic builds the complexity statistics with diagnostic.
func ComplexityStatsWithDiagnostic(f *ast.File, fset *token.FileSet, stats []Stat, enableDiagnostics bool) []Stat {
	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			d := parseDirective(fn.Doc)
			if d.Ignore {
				continue
			}

			res := ScanComplexity(fn, enableDiagnostics)

			stats = append(stats, Stat{
				PkgName:     f.Name.Name,
				FuncName:    funcName(fn),
				Complexity:  res.Complexity,
				Diagnostics: generateDiagnostics(fset, res.Diagnostics),
				Pos:         fset.Position(fn.Pos()),
			})
		}
	}

	return stats
}

func generateDiagnostics(fset *token.FileSet, diags []diagnostic) []Diagnostic {
	out := make([]Diagnostic, 0, len(diags))

	for _, diag := range diags {
		pos := fset.Position(diag.Pos)
		diagPos := DiagnosticPosition{
			Offset: pos.Offset,
			Line:   pos.Line,
			Column: pos.Column,
		}

		out = append(out, Diagnostic{
			Inc:     diag.Inc,
			Nesting: diag.Nesting,
			Text:    diag.Text,
			Pos:     diagPos,
		})
	}

	return out
}

type directive struct {
	Ignore bool
}

func parseDirective(doc *ast.CommentGroup) directive {
	if doc == nil {
		return directive{}
	}

	for _, c := range doc.List {
		if c.Text == "//gocognit:ignore" {
			return directive{Ignore: true}
		}
	}

	return directive{}
}

// funcName returns the name representation of a function or method:
// "(Type).Name" for methods or simply "Name" for functions.
func funcName(fn *ast.FuncDecl) string {
	if fn.Recv != nil {
		if fn.Recv.NumFields() > 0 {
			typ := fn.Recv.List[0].Type

			return fmt.Sprintf("(%s).%s", recvString(typ), fn.Name)
		}
	}

	return fn.Name.Name
}

// Complexity calculates the cognitive complexity of a function.
func Complexity(fn *ast.FuncDecl) int {
	res := ScanComplexity(fn, false)

	return res.Complexity
}

// ScanComplexity scans the function declaration.
func ScanComplexity(fn *ast.FuncDecl, includeDiagnostics bool) ScanResult {
	v := sonar.SonarVisitor{}
	ast.Walk(&v, fn)

	return ScanResult{
		Diagnostics: v.diagnostics,
		Complexity:  v.,
	}
}

type ScanResult struct {
	Diagnostics []diagnostic
	Complexity  int
}

type diagnostic struct {
	Inc     int
	Nesting int
	Text    string
	Pos     token.Pos
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

const Doc = `Find complex function using cognitive complexity calculation.

The gocognit analysis reports functions or methods which the complexity is over 
than the specified limit.`

// Analyzer reports a diagnostic for every function or method which is
// too complex specified by its -over flag.
var Analyzer = &analysis.Analyzer{
	Name:     "gocognit",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var (
	over int // -over flag
)

func init() {
	Analyzer.Flags.IntVar(&over, "over", over, "show functions with complexity > N only")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		funcDecl := n.(*ast.FuncDecl)

		d := parseDirective(funcDecl.Doc)
		if d.Ignore {
			return
		}

		fnName := funcName(funcDecl)

		fnComplexity := Complexity(funcDecl)

		if fnComplexity > over {
			pass.Reportf(funcDecl.Pos(), "cognitive complexity %d of func %s is high (> %d)", fnComplexity, fnName, over)
		}
	})

	return nil, nil
}
