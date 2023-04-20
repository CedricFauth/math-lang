package parser

import (
	"fmt"

	"github.com/CedricFauth/math-lang-go/ast"
	"github.com/CedricFauth/math-lang-go/lexer"
)

func TestASTPrinter1() {
	expr := &ast.Binary{
		&ast.Unary{
			lexer.NewTestToken(lexer.MINUS),
			&ast.Grouping{
				&ast.Binary{
					&ast.Literal{0.33},
					lexer.NewTestToken(lexer.PLUS),
					&ast.Literal{0.67},
				},
			},
		},
		lexer.NewTestToken(lexer.STAR),
		&ast.Literal{42},
	}
	printer := NewASTPrinter()
	str, err := ast.Accept(expr, printer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

func TestASTPrinter2() {
	expr := &ast.Binary{
		&ast.Unary{
			lexer.NewTestToken(lexer.MINUS),
			&ast.Literal{9.99},
		},
		lexer.NewTestToken(lexer.STAR),
		&ast.Literal{42},
	}
	printer := NewASTPrinter()
	str, err := ast.Accept(expr, printer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
