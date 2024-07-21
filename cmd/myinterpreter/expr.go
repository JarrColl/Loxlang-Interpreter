package main

import "fmt"

func AstPrint(expr Expr) string {
	switch typedExpr := expr.(type) {
	case Literal:
		// if expr.(Literal).value == nil return "nil"
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








/* type AstPrinterVisitor struct {}

func (self *AstPrinterVisitor) Print(expr Expr) string {
	return expr.accept(self)
}

func (self *AstPrinterVisitor) VisitForBinary(expr *Binary) string {
	return self.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}

func (self *AstPrinterVisitor) VisitForGrouping(expr *Grouping) string {

}

func (self *AstPrinterVisitor) VisitForLiteral(expr *Literal) string {

}

func (self *AstPrinterVisitor) VisitForUnary(expr *Unary) string {

}

func (self *AstPrinterVisitor) parenthesize(name string, exprs ...Expr) string {

} */
