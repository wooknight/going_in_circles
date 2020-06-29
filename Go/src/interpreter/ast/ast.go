package ast

import (
	"interpreter/token"
)

type Node interface {
	TokenLiteral() string //debugging
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement //program is a series (array) of statements
}

//TokenLiteral turns the array of statements into a Node!! to print
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}

}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) expressionNode()      {}
func (id *Identifier) TokenLiteral() string { return id.Token.Literal }
