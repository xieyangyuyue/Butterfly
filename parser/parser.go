package parser

import (
	"Butterfly/ast"
	"Butterfly/lexer"
	"fmt"
	"strconv"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  lexer.Token
	peekToken lexer.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// 读取两个token，以填充 curToken 和 peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Expression = p.parseExpression()
	return program
}

// 解析表达式
func (p *Parser) parseExpression() ast.Expression {
	left := p.parseIntegerLiteral()

	for p.peekToken.Type == lexer.PLUS || p.peekToken.Type == lexer.MINUS {
		p.nextToken() // 移动到运算符
		left = p.parseInfixExpression(left)
	}
	return left
}

// 解析整数
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Value, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Value)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

// 解析中缀表达式
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Value,
		Left:     left,
	}

	p.nextToken() // 移动到右边的表达式

	expression.Right = p.parseIntegerLiteral()

	return expression
}
