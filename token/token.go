package token

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifier + literal
	IDENT = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "ASSIGN"
	PLUS = "PLUS"

	// Delimiters
	COMMA = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "fn"
	LET = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdentType(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
