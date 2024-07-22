package main

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) Parser {
	return Parser{tokens: tokens, current: 0}
}

func (self *Parser) Parse() Expr {
	expr, err := self.expression()
	if err != nil {
		return nil
	}

	return expr
}

func (self *Parser) advance() *Token {
	if !self.isAtEnd() {
		self.current++
	}
	return self.previous()
}

func (self *Parser) match(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		if self.check(tokenType) {
			self.advance()
			return true
		}
	}

	return false
}

func (self *Parser) check(tokenType TokenType) bool {
	if self.isAtEnd() {
		return false
	}

	return self.peekCurrent().tokenType == tokenType
}

func (self *Parser) isAtEnd() bool {
	return self.peekCurrent().tokenType == EOF
}

func (self *Parser) peekCurrent() *Token {
	return &self.tokens[self.current]
}

func (self *Parser) previous() *Token {
	return &self.tokens[self.current-1]
}

// Parsing rule functions below here

func (self *Parser) expression() (Expr, error) {
	return self.equality()
}

func (self *Parser) equality() (Expr, error) {
	expr, err := self.comparison()
	if err != nil {
		return nil, err
	}

	for self.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := self.previous()
		right, err := self.comparison()
		if err != nil {
			return nil, err
		}
		expr = Binary{left: expr, operator: operator, right: right}
	}
	return expr, nil
}

func (self *Parser) comparison() (Expr, error) {
	expr, err := self.term()
	if err != nil {
		return nil, err
	}

	for self.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := self.previous()
		right, err := self.term()
		if err != nil {
			return nil, err
		}
		expr = Binary{left: expr, operator: operator, right: right}
	}
	return expr, nil
}

func (self *Parser) term() (Expr, error) {
	expr, err := self.factor()
	if err != nil {
		return nil, err
	}

	for self.match(MINUS, PLUS) {
		operator := self.previous()
		right, err := self.factor()
		if err != nil {
			return nil, err
		}
		expr = Binary{left: expr, operator: operator, right: right}
	}
	return expr, nil
}

func (self *Parser) factor() (Expr, error) {
	expr, err := self.unary()
	if err != nil {
		return nil, err
	}

	for self.match(SLASH, STAR) {
		operator := self.previous()
		right, err := self.unary()
		if err != nil {
			return nil, err
		}
		expr = Binary{left: expr, operator: operator, right: right}
	}
	return expr, nil
}

func (self *Parser) unary() (Expr, error) {
	if self.match(BANG, MINUS) {
		operator := self.previous()
		right, err := self.unary()
		if err != nil {
			return nil, err
		}
		return Unary{operator: operator, right: right}, nil
	}

	expr, err := self.primary()
	if err != nil {
		return nil, err
	}

	return expr, nil
}

func (self *Parser) primary() (Expr, error) {
	if self.match(NUMBER, STRING) {
		return Literal{value: self.previous().literal}, nil
	}

	if self.match(FALSE) {
		return Literal{value: false}, nil
	}
	if self.match(TRUE) {
		return Literal{value: true}, nil
	}
	if self.match(NIL) {
		return Literal{value: nil}, nil
	}

	if self.match(LEFT_PAREN) {
		expr, err := self.expression()
		if err != nil {
			return nil, err
		}

		_, err = self.consume(RIGHT_PAREN, "Expect ') after expression.")
		if err != nil {
			return nil, err
		}
		return Grouping{expression: expr}, nil
	}

	return nil, self.parseError(self.peekCurrent(), "Expect expression.")
}

// Error Handling
type ParseError struct {
	line    int
	message string
}

func (err *ParseError) Error() string {
	return err.message
}

func (self *Parser) parseError(token *Token, message string) error {
	ReportTokenError(token, message)
	return &ParseError{
		line:    token.line,
		message: message,
	}
}

func (self *Parser) consume(tokenType TokenType, message string) (*Token, error) {
	if self.check(tokenType) {
		return self.advance(), nil
	}

	return nil, self.parseError(self.peekCurrent(), message)
}

func (self *Parser) synchronize() {
	self.advance()

	for !self.isAtEnd() {
		if self.previous().tokenType == SEMICOLON {
			return
		}

		switch self.peekCurrent().tokenType {
		case CLASS:
		case FUN:
		case VAR:
		case FOR:
		case IF:
		case WHILE:
		case PRINT:
		case RETURN:
			return
		}

		self.advance()
	}
}
