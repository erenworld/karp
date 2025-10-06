package ast

// Recursive descent parser. And in particular, it’s a “top down operator precedence” parser
// sometimes called “Pratt parser”, after its inventor Vaughan Pratt

import "github.com/erenworld/karp/token"

type Node interface {
	TokenLiteral() string
}

type LetStatement struct {
	Token 	token.Token // the token.LET token
	Name 	*Identifier
	Value 	Expression
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}


type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

type Program struct {
	Statement []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

