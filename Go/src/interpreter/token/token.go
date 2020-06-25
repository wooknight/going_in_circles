package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	//operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	DIVIDE = "/"

	//Delimiters
	SEMICOLON = ";"
	PERIOD    = "."
	COMMA     = ","

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "["
	RBRACE = "]"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
