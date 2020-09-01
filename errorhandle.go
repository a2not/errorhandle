package errorhandle

import (
	"go/ast"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "errorhandle is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "errorhandle",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			obj, ok := pass.TypesInfo.Defs[n]
			if !ok || obj == nil {
				return
			}

			if obj.Name() != "_" {
				return
			}

			if !analysisutil.ImplementsError(obj.Type()) {
				return
			}
			pass.Reportf(n.Pos(), "receiving error with _")
		}
	})

	return nil, nil
}
