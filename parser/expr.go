package parser

import (
	"errors"
	"fmt"

	"github.com/CedricFauth/math-lang-go/lexer"
)

type Visitor[T any] interface {
	visitLiteral(*Literal) T
	visitGrouping(*Grouping) T
	visitUnary(*Unary) T
	visitBinary(*Binary) T
}

type Expr interface{}

type Literal struct {
	value any
}

type Grouping struct {
	expression Expr
}

type Unary struct {
	operator   *lexer.Token
	expression Expr
}

type Binary struct {
	left     Expr
	operator *lexer.Token
	right    Expr
}

func accept[T any](expr any, v Visitor[T]) (T, error) {
	switch e := expr.(type) {
	case *Literal:
		return v.visitLiteral(e), nil
	case *Grouping:
		return v.visitGrouping(e), nil
	case *Unary:
		return v.visitUnary(e), nil
	case *Binary:
		return v.visitBinary(e), nil

	default:
		var t T
		return t, errors.New(fmt.Sprintf("Error accept wrong type: %T", expr))
	}
}
