// Package lexer provides lexical analysis for the Monkey programming language.
//
// It reads an input string character-by-character and produces a stream of tokens
// that represent the smallest meaningful units in the language syntax.
// These tokens are then used by the parser to construct an abstract syntax tree (AST).
//
// The lexer handles single-character tokens (like '+', '-', etc.), multi-character
// operators (like '==' and '!='), identifiers, numbers, and keywords.
package lexer

import "github.com/saikrishnamohan7/monkey/token"

// Lexer represents a lexical analyzer for the Monkey programming language.
// It takes an input string and produces a stream of tokens for parsing.
type Lexer struct {
	// input is the source code being tokenized.
	input string

	// position is the current position in input (points to current char).
	position int

	// readPosition is the next position to read (used for lookahead).
	readPosition int

	// charByte holds the current character under examination.
	charByte byte
}

// NewLexer Instantiates and returns a Lexer instance
func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
		charByte:     0,
	}
	lexer.readChar()

	return lexer
}

func (lexer *Lexer) readChar() {
	// Have reached the end of the input?
	if lexer.readPosition >= len(lexer.input) {
		lexer.charByte = 0
	} else {
		lexer.charByte = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition = lexer.readPosition + 1
}

// NextToken retrieves the next token from the input and advances the lexer.
func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.charByte {
	case '=':
		tok = newToken(token.ASSIGN, lexer.charByte)
	case '+':
		tok = newToken(token.PLUS, lexer.charByte)
	case '(':
		tok = newToken(token.LPAREN, lexer.charByte)
	case ')':
		tok = newToken(token.RPAREN, lexer.charByte)
	case ',':
		tok = newToken(token.COMMA, lexer.charByte)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.charByte)
	case '{':
		tok = newToken(token.LBRACE, lexer.charByte)
	case '}':
		tok = newToken(token.RBRACE, lexer.charByte)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.charByte) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(lexer.charByte) {
			tok.Literal = lexer.readNumber()
			tok.Type = token.INT
			return tok
		}
		tok = newToken(token.ILLEGAL, lexer.charByte)
	}

	lexer.readChar()

	return tok
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.charByte) {
		// changes lexer.position
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.charByte == ' ' || lexer.charByte == '\t' || lexer.charByte == '\n' || lexer.charByte == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.charByte) {
		lexer.readChar()
	}

	return lexer.input[position : lexer.position]
}

func newToken(tokenType token.TokenType, charByte byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(charByte),
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
