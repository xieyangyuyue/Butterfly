package ast

import "Butterfly/lexer"

// 所有节点类型都必须实现 Node 接口
type Node interface {
	TokenLiteral() string
}

// 所有表达式节点都必须实现 Expression 接口
type Expression interface {
	Node
	expressionNode()
}

// Program 是每个 AST 的根节点
type Program struct {
	Expression Expression
}

func (p *Program) TokenLiteral() string { return "Program" }

// 整数リテラル
type IntegerLiteral struct {
	Token lexer.Token // a lexer.TOKEN_INT
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Value }

// 中缀表达式，如 a + b 或 a - b
type InfixExpression struct {
	Token    lexer.Token // 运算符Token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Value }
