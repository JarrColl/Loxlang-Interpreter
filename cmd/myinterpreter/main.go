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
		cmd_tokenise()
	} else {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	if had_error {
		os.Exit(65)
	}
}

func report_error(line int, message string) {
	report(line, "", message);
}

func report(line int, where string, message string) {
	fmt.Println("[line " + strconv.Itoa(line) + "] Error" + where + ": " + message);
	had_error = true;
}

func cmd_tokenise() {
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		scanner := NewScanner(string(fileContents))

		scanner.ScanTokens()

		for _, token := range scanner.tokens {
			fmt.Println(token.toString())
		}
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
