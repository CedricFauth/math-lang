package parser

import (
	"fmt"
	"strings"
)

type ASTPrinter struct {
	padding int
}

func (p *ASTPrinter) indent() int {
	p.padding += 1
	return p.padding
}
func (p *ASTPrinter) dedent() int {
	if p.padding > 0 {
		p.padding -= 1
	}
	return p.padding
}

func (p *ASTPrinter) pad() string {
	return strings.Repeat("    ", p.padding)
}
func (p *ASTPrinter) visitLiteral(expr *Literal) string {
	return fmt.Sprintf("%v%T(%v)", p.pad(), expr.value, expr.value)
}

func (p *ASTPrinter) visitGrouping(expr *Grouping) string {
	p.indent()
	val, err := accept[string](expr.expression, p)
	if err != nil {
		return err.Error()
	}
	p.dedent()
	return fmt.Sprintf("%vGROUP(\n%v\n%v)", p.pad(), val, p.pad())
}

func (p *ASTPrinter) visitUnary(expr *Unary) string {
	p.indent()
	pad := p.pad()
	val, err := accept[string](expr.expression, p)
	if err != nil {
		return err.Error()
	}
	p.dedent()
	return fmt.Sprintf("%vUNARY(\n%v%v\n%v\n%v)", p.pad(), pad, expr.operator.Lexeme(), val, p.pad())
}

func (p *ASTPrinter) visitBinary(expr *Binary) string {
	p.indent()
	pad := p.pad()
	valLeft, err := accept[string](expr.left, p)
	if err != nil {
		return err.Error()
	}
	valRight, err := accept[string](expr.right, p)
	if err != nil {
		return err.Error()
	}
	p.dedent()

	return fmt.Sprintf("%vBINARY(\n%v%v\n%v\n%v\n%v)", p.pad(), pad, expr.operator.Lexeme(), valLeft, valRight, p.pad())

}

func NewASTPrinter() Visitor[string] {
	return &ASTPrinter{}
}
