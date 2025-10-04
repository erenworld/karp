package token

type TokenType string

type Token struct {
	Type 		TokenType
	Literal 	string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF 	= "EOF"

	// Identifiers + literals
	IDENT 	= "IDENT" 
	INT		= "INT"

	// Operators
	ASSIGN	= "="
	PLUS	= "+"
	MUL		= "*"
	SUB		= "-"

	// Delimiters
	COMMA		= ","
	SEMICOLON	= ";"


	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keyword
	FUNCTION 	= "FUNCTION"
	LET 		= "LET"
)