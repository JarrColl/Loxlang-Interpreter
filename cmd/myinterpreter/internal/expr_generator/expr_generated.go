package main


type Expr interface {
	accept(Visitor)
}

type Binary struct {
	left Expr
	operator Token
	right Expr
}

func (self *Binary) accept(visitor Visitor) {
	visitor.VisitForBinary()
}

type Grouping struct {
	expression Expr
}

func (self *Grouping) accept(visitor Visitor) {
	visitor.VisitForGrouping()
}

type Literal struct {
	value any
}

func (self *Literal) accept(visitor Visitor) {
	visitor.VisitForLiteral()
}

type Unary struct {
	operator Token
	right Expr
}

func (self *Unary) accept(visitor Visitor) {
	visitor.VisitForUnary()
}

type Visitor interface {
	VisitForBinary()
	VisitForGrouping()
	VisitForLiteral()
	VisitForUnary()
}

