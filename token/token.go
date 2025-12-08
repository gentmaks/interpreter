// Package token defines the token struct
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// special token types
	ILLEGALL = "ILLEGAl"
	EOF      = "EOF"

	// identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"
	BOOL  = "BOOL"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
)
