package parser

import (
	"fmt"
	"os"

	"github.com/CedricFauth/math-lang-go/ast"
	"github.com/CedricFauth/math-lang-go/lexer"
)

type Parser struct {
	index  int
	tokens []*lexer.Token
}

func New(tokens []*lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) error(e string) {
	fmt.Printf("%v: '%v'\n", e, p.previous().Lexeme())
	os.Exit(1)
}

func (p *Parser) expect(tokenType lexer.TokenType, message string) {
	if !p.match(tokenType) {
		p.error(message)
	}
}

func (p *Parser) previous() *lexer.Token {
	if p.index-1 >= 0 {
		return p.tokens[p.index-1]
	}
	return p.tokens[0]
}

func (p *Parser) advance() {
	if p.index < len(p.tokens) {
		p.index++
	}
}

func (p *Parser) match(tokenType lexer.TokenType) bool {
	if p.index >= len(p.tokens) {
		return false
	}
	if tokenType == p.tokens[p.index].Type() {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) expression() ast.Expr {
	// find identifier
	expr := p.term()
	if p.match(lexer.EQUAL) {
		if variable, ok := expr.(*ast.Variable); ok {
			value := p.term()
			return &ast.Assignment{Name: variable.Name, Value: value}
		} else {
			p.error("Expected a variable")
		}
	}
	return expr
}

func (p *Parser) term() ast.Expr {
	expr := p.factor()
	for p.match(lexer.PLUS) || p.match(lexer.MINUS) {
		op := p.previous()
		right := p.factor()
		expr = &ast.Binary{expr, op, right}
	}
	return expr
}

func (p *Parser) factor() ast.Expr {
	expr := p.unary()
	for p.match(lexer.STAR) || p.match(lexer.SLASH) {
		op := p.previous()
		right := p.unary()
		expr = &ast.Binary{expr, op, right}
	}
	return expr
}

func (p *Parser) unary() ast.Expr {
	if p.match(lexer.MINUS) {
		op := p.previous()
		expr := p.unary()
		return &ast.Unary{op, expr}
	} else {
		return p.primary()
	}
}

func (p *Parser) primary() ast.Expr {
	if p.match(lexer.NUMBER) {
		return &ast.Literal{p.previous().Value()}
	}
	if p.match(lexer.PAREN_OPEN) {
		expr := p.expression() // TODO change to expression
		p.expect(lexer.PAREN_CLOSE, "Expect ')' after expression")
		return &ast.Grouping{expr}
	}
	if p.match(lexer.IDENTIFIER) {
		return &ast.Variable{p.previous().Lexeme()}
	}

	p.error("Missing or not an expression")
	return nil
}

func (p *Parser) Parse(tokens []*lexer.Token) ast.Expr {
	expr := p.expression() // TODO change to expression
	printer := NewASTPrinter()

	str, err := ast.Accept(expr, printer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	if p.index != len(p.tokens) {
		p.advance()
		p.error("Unallowed Token")
	}
	return expr
}
