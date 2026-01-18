package token

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifier + literal
	IDENT = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "ASSIGN"
	EQ = "=="
	NOT_EQ = "!="
	PLUS = "PLUS"
	MINUS = "MINUS"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdentType(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
