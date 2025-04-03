package token

type TokenType string

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
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
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
