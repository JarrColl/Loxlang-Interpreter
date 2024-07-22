package main

import (
	"fmt"
	"strings"
)

func AstPrint(expr Expr) string {
	switch typedExpr := expr.(type) {
	case Literal:
		if typedExpr.value == nil {
			return "nil"
		}

		switch typedExpr.value.(type) {
		case float64:
			var decimal = fmt.Sprintf("%.12f", typedExpr.value)
			decimal = strings.TrimRight(decimal, "0")

			if decimal[len(decimal)-1] == '.' {
				decimal = decimal + string('0')
			}
			return decimal
		}

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
