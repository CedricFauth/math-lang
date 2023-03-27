package parser

import (
	"fmt"

	"github.com/CedricFauth/math-lang-go/lexer"
)

func TestASTPrinter1() {
	expr := &Binary{
		&Unary{
			lexer.NewTestToken(lexer.MINUS),
			&Grouping{
				&Binary{
					&Literal{0.33},
					lexer.NewTestToken(lexer.PLUS),
					&Literal{0.67},
				},
			},
		},
		lexer.NewTestToken(lexer.STAR),
		&Literal{42},
	}
	printer := NewASTPrinter()
	str, err := accept(expr, printer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

func TestASTPrinter2() {
	expr := &Binary{
		&Unary{
			lexer.NewTestToken(lexer.MINUS),
			&Literal{9.99},
		},
		lexer.NewTestToken(lexer.STAR),
		&Literal{42},
	}
	printer := NewASTPrinter()
	str, err := accept(expr, printer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
