// Package token provides the definitions and utilities for lexical tokens
// used in the Monkey programming language. It defines token types, their
// representations, and helper functions for token classification.
//
// Constants:
// - ILLEGAL: Represents an unknown or invalid token.
// - EOF: Represents the end of file token.
// - IDENT: Represents an identifier (e.g., variable names like add, foobar, x, y).
// - INT: Represents an integer literal.
// - ASSIGN: Represents the assignment operator "=".
// - PLUS: Represents the addition operator "+".
// - COMMA: Represents the delimiter ",".
// - SEMICOLON: Represents the delimiter ";".
// - LPAREN: Represents the left parenthesis "(".
// - RPAREN: Represents the right parenthesis ")".
// - LBRACE: Represents the left brace "{".
// - RBRACE: Represents the right brace "}".
// - FUNCTION: Represents the "fn" keyword.
// - LET: Represents the "let" keyword.
//
// Variables:
// - keywords: A map of string keywords to their corresponding TokenType.
//
// Functions:
//   - LookupIdentifier: Checks if a given identifier is a keyword and returns
//     its corresponding TokenType. If the identifier is not a keyword, it
//     defaults to returning IDENT.
package token

// TokenType is a string alias used to represent the type of a token.
type TokenType string

// Token is a struct that encapsulates the type and literal value of a token.
type Token struct {
	Type    TokenType
	Literal string // storing this as a byte would be performant but debugging would be hard
}

const (
	// ILLEGAL represents an unknown or invalid token
	ILLEGAL = "ILLEGAL"
	// EOF represents the end of file token
	EOF = "EOF"

	// IDENT represents an identifier like add, foobar, x, y, etc.
	IDENT = "IDENT"
	// INT represents an integer literal.
	INT = "INT"

	// ASSIGN represents the assignment operator "=".
	ASSIGN = "="
	// PLUS represents the addition operator "+".
	PLUS = "+"
	// MINUS represents the subtraction operator "-".
	MINUS = "-"
	// BANG represents the logical NOT operator "!".
	BANG = "!"
	// ASTERISK represents the multiplication operator "*".
	ASTERISK = "*"
	// SLASH represents the division operator "/".
	SLASH = "/"

	// LT represents the less-than operator "<".
	LT = "<"
	// GT represents the greater-than operator ">".
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// COMMA represents the delimiter ",".
	COMMA = ","
	// SEMICOLON represents the delimiter ";".
	SEMICOLON = ";"
	// LPAREN represents the left parenthesis "(".
	LPAREN = "("
	// RPAREN represents the right parenthesis ")".
	RPAREN = ")"
	// LBRACE represents the left brace "{".
	LBRACE = "{"
	// RBRACE represents the right brace "}".
	RBRACE = "}"

	// FUNCTION represents the "fn" keyword.
	FUNCTION = "FUNCTION"
	// LET represents the "let" keyword.
	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdentifier checks if an identifier is a keyword and returns its TokenType.
// If the identifier is not a keyword, it returns IDENT.
func LookupIdentifier(identifier string) TokenType {
	tok, ok := keywords[identifier]

	if ok {
		return tok
	}

	return IDENT
}
