package main

import (
	"log"
	"os"
	"strings"
)

var expressions = [4][2]string{
	{"Binary", "left Expr, operator Token, right Expr"},
	{"Grouping", "expression Expr"},
	{"Literal", "value any"},
	{"Unary", "operator Token, right Expr"},
}

func main() {
	f, err := os.Create("../../expr_generated.go")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString("package main\n")
	f.WriteString(`
type Expr interface {
	Accept(Visitor)
}

`)

	for _, expression := range expressions {
		f.WriteString("type " + expression[0] + " struct {")
		var properties = strings.Split(expression[1], ", ")
		for _, str_prop := range properties {
			f.WriteString("\n\t" + str_prop)
		}
		f.WriteString("\n}\n\n")

		f.WriteString("func (self *" + expression[0] + ") Accept(visitor Visitor) {")
		f.WriteString("\n\tvisitor.VisitFor" + expression[0] + "(self)")
		f.WriteString("\n}\n\n")
	}

	f.WriteString("type Visitor interface {")
	for _, expression := range expressions {
		f.WriteString("\n\tVisitFor" + expression[0] + "(expr *" + expression[0] + ")")
	}
	f.WriteString("\n}\n\n")

}
