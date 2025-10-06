package ast

// Recursive descent parser. And in particular, it’s a “top down operator precedence” parser
// sometimes called “Pratt parser”, after its inventor Vaughan Pratt
// let <identifier> = <expression>;

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}