package errorhandle

import (
	"errors"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "errorhandle is a static analysis tool which checks if received errors are properly handled."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "errorhandle",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var errType = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
		(*ast.ValueSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.AssignStmt:
			if flag, err := assignErrorToBlank(n.Lhs, n.Rhs, pass); err == nil && flag {
				pass.Reportf(n.Pos(), "receiving error with _")
			}
		case *ast.ValueSpec:
			if flag, err := assignErrorToBlank(n.Names, n.Values, pass); err == nil && flag {
				pass.Reportf(n.Pos(), "receiving error with _")
			}
		}
	})

	return nil, nil
}

func assignErrorToBlank(lhs interface{}, rhs []ast.Expr, pass *analysis.Pass) (bool, error) {
	rhsTypes := make([]types.Type, 0)
	for _, expr := range rhs {
		// ignore when assigning of _ = (error)(nil)
		callexpr, ok := expr.(*ast.CallExpr)
		if ok && len(callexpr.Args) == 1 {
			ident, ok := callexpr.Args[0].(*ast.Ident)
			if ok && ident.Name == "nil" {
				rhsTypes = append(rhsTypes, types.Typ[types.Int8])
				continue
			}
		}

		typ := pass.TypesInfo.TypeOf(expr)
		switch typ := typ.(type) {
		case *types.Tuple:
			for i := 0; i < typ.Len(); i++ {
				rhsTypes = append(rhsTypes, typ.At(i).Type())
			}
		default:
			rhsTypes = append(rhsTypes, typ)
		}
	}

	lhsNames := make([]string, 0)
	switch lhs := lhs.(type) {
	case []ast.Expr:
		for _, expr := range lhs {
			switch expr := expr.(type) {
			case *ast.Ident:
				lhsNames = append(lhsNames, expr.Name)
			}
		}
	case []*ast.Ident:
		for _, expr := range lhs {
			lhsNames = append(lhsNames, expr.Name)
		}
	default:
		return false, errors.New("Unexpected type of LHS")
	}

	if len(lhsNames) != len(rhsTypes) {
		return false, errors.New("(# of idents in LHS) != (# of return values in RHS)")
	}

	for i := 0; i < len(lhsNames); i++ {
		if lhsNames[i] == "_" && isError(rhsTypes[i]) {
			return true, nil
		}
	}

	return false, nil
}

func isError(typ types.Type) bool {
	return types.Implements(typ, errType) || types.Implements(types.NewPointer(typ), errType)
}
