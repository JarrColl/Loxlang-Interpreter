package main

import (
	"fmt"
	"os"
	"strconv"
)

var had_error = false

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command == "tokenize" {
		for _, token := range cmdTokenise() {
			fmt.Println(token.toString())
		}
	} else if command == "parse" {
		cmdParse()
	} else if command == "print" {
		var expression Expr = Binary{
			left: Unary{
				operator: &Token{tokenType: MINUS, lexeme: "-", literal: "null", line: 1},
				right:    Literal{123},
			},
			operator: &Token{tokenType: STAR, lexeme: "*", literal: "null", line: 1},
			right:    Grouping{expression: Literal{value: 45.67}},
		}

		fmt.Println(AstPrint(expression))
	} else {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	if had_error {
		os.Exit(65)
	}

}

func ReportError(line int, message string) {
	report(line, "", message)
}

func ReportTokenError(token *Token, message string) {
    if token.tokenType == EOF {
      report(token.line, " at end", message);
    } else {
      report(token.line, " at '" + token.lexeme + "'", message);
    }
}

func report(line int, where string, message string) {
	fmt.Fprintln(os.Stderr, "[line "+strconv.Itoa(line)+"] Error"+where+": "+message)
	had_error = true
}

func cmdTokenise() []Token {
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		scanner := NewScanner(string(fileContents))

		scanner.ScanTokens()

		return scanner.tokens
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}

	return nil //TODO: return error here
}

func cmdParse() {
	parser := NewParser(cmdTokenise())
	expr := parser.Parse()

	fmt.Println(AstPrint(expr))
}
