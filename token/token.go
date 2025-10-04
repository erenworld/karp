package token

type TokenType string

type Token struct {
	Type 		TokenType
	Litteral 	string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF 	= "EOF"

	// Identifiers + litterals
	IDENT 	= "IDENT" 
	INT		= "INT"

	// Operators
	ASSIGN	= "="
	PLUS	= "+"

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