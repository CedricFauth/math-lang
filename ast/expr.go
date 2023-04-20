package ast

import (
	"fmt"

	"github.com/CedricFauth/math-lang-go/lexer"
)

type Visitor[T any] interface {
	VisitVariable(*Variable) T
	VisitAssignment(*Assignment) T
	VisitLiteral(*Literal) T
	VisitGrouping(*Grouping) T
	VisitUnary(*Unary) T
	VisitBinary(*Binary) T
}

type Expr interface{}

type Variable struct {
	Name string
}

type Assignment struct {
	Name  string
	Value Expr
}

type Literal struct {
	Value any
}

type Grouping struct {
	Expression Expr
}

type Unary struct {
	Operator   *lexer.Token
	Expression Expr
}

type Binary struct {
	Left     Expr
	Operator *lexer.Token
	Right    Expr
}

func Accept[T any](expr any, v Visitor[T]) (T, error) {
	switch e := expr.(type) {
	//case *Variable:
	//	return v.VisitVariable(e), nil
	case *Variable:
		return v.VisitVariable(e), nil
	case *Assignment:
		return v.VisitAssignment(e), nil
	case *Literal:
		return v.VisitLiteral(e), nil
	case *Grouping:
		return v.VisitGrouping(e), nil
	case *Unary:
		return v.VisitUnary(e), nil
	case *Binary:
		return v.VisitBinary(e), nil

	default:
		var t T
		return t, fmt.Errorf("error accept wrong type: %T", expr)
	}
}
