package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

var TokenTypeStrings = [39]string{
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

var Identifiers = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

func TokenTypeToString(token_type TokenType) string {
	return TokenTypeStrings[token_type]
}

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   any
	line      int
}

func (self *Token) toString() string {
	var output string = TokenTypeToString(self.tokenType) + " " + self.lexeme

	switch self.literal.(type) {
	case float64:
		var decimal = fmt.Sprintf(" %.12f", self.literal)
		decimal = strings.TrimRight(decimal, "0")

		if decimal[len(decimal)-1] == '.' {
			decimal = decimal + string('0')
		}
		return output + decimal
	default:
		return output + fmt.Sprintf(" %v", self.literal)
	}
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
	self.addTokenWithLiteral(t, "null")
}

func (self *Scanner) addTokenWithLiteral(t TokenType, literal any) {
	var text string = self.source[self.start:self.current]
	self.tokens = append(self.tokens, Token{t, text, literal, self.line})
}

func (self *Scanner) matchCurrentChar(expected byte) bool {
	if self.isAtEnd() {
		return false
	}
	if self.source[self.current] != expected {
		return false
	}

	self.incrementCurrent()
	return true
}

func (self *Scanner) peekCurrent() byte {
	if self.isAtEnd() {
		return 0
	}
	return self.source[self.current]
}

func (self *Scanner) peekNext() byte {
	if (self.current + 1) >= len(self.source) {
		return 0
	}
	return self.source[self.current+1]
}

func (self *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (self *Scanner) isAlpha(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_'
}

func (self *Scanner) isAlphaNumeric(c byte) bool {
	return self.isAlpha(c) || self.isDigit(c)
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
	case '!':
		if self.matchCurrentChar('=') {
			self.addToken(BANG_EQUAL)
		} else {
			self.addToken(BANG)
		}
	case '=':
		if self.matchCurrentChar('=') {
			self.addToken(EQUAL_EQUAL)
		} else {
			self.addToken(EQUAL)
		}
	case '<':
		if self.matchCurrentChar('=') {
			self.addToken(LESS_EQUAL)
		} else {
			self.addToken(LESS)
		}
	case '>':
		if self.matchCurrentChar('=') {
			self.addToken(GREATER_EQUAL)
		} else {
			self.addToken(GREATER)
		}
	case '/':
		if self.matchCurrentChar('/') {
			for self.peekCurrent() != '\n' && !self.isAtEnd() {
				self.advance()
			}
		} else {
			self.addToken(SLASH)
		}
	case ' ':
	case '\t':
	case '\r':
	case '\n':
		self.line++
	case '"':
		self.stringFunc()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		self.numberFunc()
	default:
		if self.isAlpha(c) {
			self.identifierFunc()
		} else {
			ReportError(self.line, "Unexpected character: "+string(c))
		}
	}
}

func (self *Scanner) numberFunc() {
	self.start = self.current - 1

	for self.isDigit(self.peekCurrent()) {
		self.advance()
	}

	if self.peekCurrent() == '.' && self.isDigit(self.peekNext()) {
		self.advance() // consume the "."

		for self.isDigit(self.peekCurrent()) {
			self.advance()
		}
	}

	// Parse float from string
	if floatValue, err := strconv.ParseFloat(self.source[self.start:self.current], 64); err == nil {
		self.addTokenWithLiteral(NUMBER, floatValue)
	} else {
		ReportError(self.line, fmt.Sprintf("Error parsing float: %v", err))
	}
}

func (self *Scanner) stringFunc() {
	// Set start to the first "
	self.start = self.current - 1
	for self.peekCurrent() != '"' && !self.isAtEnd() {
		self.advance()
	}

	if self.isAtEnd() {
		ReportError(self.line, "Unterminated string.")
		return
	}

	self.advance()

	var str_value string = self.source[self.start+1 : self.current-1]
	self.addTokenWithLiteral(STRING, str_value)

}

func (self *Scanner) identifierFunc() {
	self.start = self.current - 1
	for self.isAlphaNumeric(self.peekCurrent()) {
		self.advance()
	}

	text := self.source[self.start:self.current]
	token_type, exists := Identifiers[text]

	if exists {
		self.addToken(token_type)
	} else {
		self.addToken(IDENTIFIER)
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
