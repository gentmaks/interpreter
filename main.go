package main

import (
	"fmt"

	"github.com/gentmaks/interpreter/lexer"
)

func main() {
	l := lexer.New("=+(){},;")
	tok := l.NextToken()
	for tok.Literal != "" {
		fmt.Println(tok)
		tok = l.NextToken()
	}
}
