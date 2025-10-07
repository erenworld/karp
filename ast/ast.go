package ast

// Recursive descent parser. And in particular, it’s a “top down operator precedence” parser
// sometimes called “Pratt parser”, after its inventor Vaughan Pratt

import (
	"bytes"

	"karp/token"
) 

type Node interface {
	TokenLiteral() 	string
	String()		string
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

type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

type ReturnStatement struct {
	Token 		token.Token // return token
	ReturnValue	Expression
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

