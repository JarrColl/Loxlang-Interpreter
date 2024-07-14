package main

type TokenType = int8

const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

var x = [39]string{
	// Single-character tokens.
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"COMMA",
	"DOT",
	"MINUS",
	"PLUS",
	"SEMICOLON",
	"SLASH",
	"STAR",
	// One or two character tokens.
	"BANG",
	"BANG_EQUAL",
	"EQUAL",
	"EQUAL_EQUAL",
	"GREATER",
	"GREATER_EQUAL",
	"LESS",
	"LESS_EQUAL",
	// Literals.
	"IDENTIFIER",
	"STRING",
	"NUMBER",
	// Keywords.
	"AND",
	"CLASS",
	"ELSE",
	"FALSE",
	"FUN",
	"FOR",
	"IF",
	"NIL",
	"OR",
	"PRINT",
	"RETURN",
	"SUPER",
	"THIS",
	"TRUE",
	"VAR",
	"WHILE",
	"EOF",
}

func TokenTypeToString(token_type TokenType) string {
	return x[token_type]
}

type Token struct {
	token_type TokenType
	lexeme     string
	literal    string
	line       int
}

func (self *Token) toString() string {
	return TokenTypeToString(self.token_type) + " " + self.lexeme + " " + self.literal
}

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) Scanner {
	return Scanner{source: source, start: 0, current: 0, line: 1}
}

func (self *Scanner) isAtEnd() bool {
	return self.current >= len(self.source)
}

func (self *Scanner) incrementCurrent() {
	self.current++
}

func (self *Scanner) advance() byte {
	defer self.incrementCurrent()
	return self.source[self.current]
}

func (self *Scanner) addToken(t TokenType) {
	self.addTokenWithLexeme(t, "null")
}

func (self *Scanner) addTokenWithLexeme(t TokenType, lexeme string) {
	var text string = self.source[self.start:self.current]
	self.tokens = append(self.tokens, Token{t, text, lexeme, self.line})
}

func (self *Scanner) scanToken() {
	var c byte = self.advance()

	switch c {
	case '(':
		self.addToken(LEFT_PAREN)
	case ')':
		self.addToken(RIGHT_PAREN)
	case '{':
		self.addToken(LEFT_BRACE)
	case '}':
		self.addToken(RIGHT_BRACE)
	case ',':
		self.addToken(COMMA)
	case '.':
		self.addToken(DOT)
	case '-':
		self.addToken(MINUS)
	case '+':
		self.addToken(PLUS)
	case ';':
		self.addToken(SEMICOLON)
	case '*':
		self.addToken(STAR)
	default:
		// TODO: lex.error
	}
}

func (self *Scanner) ScanTokens() []Token {
	for {
		if self.isAtEnd() {
			break
		}

		self.start = self.current
		self.scanToken()
	}

	self.tokens = append(self.tokens, Token{EOF, "", "null", self.line})
	return self.tokens
}
