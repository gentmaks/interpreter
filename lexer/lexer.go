// Package lexer has the logic of our lexer/tokenizer with the main NextToken() function
package lexer

import (
	"github.com/gentmaks/interpreter/token"
)

type Lexer struct{}

func (l *Lexer) NextToken() *token.Token {
	return &token.Token{Type: token.ASSIGN, Literal: "Hello"}
}

func New(s string) *Lexer {
	return &Lexer{}
}
