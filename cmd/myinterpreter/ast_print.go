package main

import "fmt"

func AstPrint(expr Expr) string {
	switch typedExpr := expr.(type) {
	case Literal:
		if expr.(Literal).value == nil {return "nil"}
		return fmt.Sprintf("%v", typedExpr.value)
	case Unary:
		return parenthesize(typedExpr.operator.lexeme, typedExpr.right)
	case Binary:
		return parenthesize(typedExpr.operator.lexeme, typedExpr.left, typedExpr.right)
	case Grouping:
		return parenthesize("group", typedExpr.expression)
	default:
		// TODO: error make this an actual error
		return "error"
	}
}

func parenthesize(name string, exprs ...Expr) string {
	var returnStr string = "(" + name
	
	for _, expr := range exprs {
		returnStr += " "	
		returnStr += AstPrint(expr)
	}
	returnStr += ")"

	return returnStr
}

