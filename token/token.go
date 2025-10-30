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
	ASSIGN		= "="
	PLUS		= "+"
	MINUS		= "-"
	BANG 		= "!"
	ASTERISK 	= "*"
	SLASH 		= "/"
	LT 			= "<"
	GT 			= ">"

	EQ 		= "=="
	NOT_EQ 	= "!="
 
	// Delimiters
	COMMA		= ","
	SEMICOLON	= ";"
	COLON		= ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keyword
	FUNCTION 	= "FUNCTION"
	STRING		= "STRING"
	LET 		= "LET"
	TRUE 		= "TRUE"
	FALSE 		= "FALSE"
	IF 			= "IF"
	ELSE 		= "ELSE"
	RETURN 		= "RETURN"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

// Checks the keywords table to see whether the given identifier is in fact a keyword.
// If it is, it returns the keyword’s TokenType constant. 
// If it isn’t, we just get back token.IDENT, which is the TokenType for all user-defined identifiers.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}