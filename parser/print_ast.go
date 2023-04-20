package parser

import (
	"fmt"
	"strings"

	"github.com/CedricFauth/math-lang-go/ast"
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

func (p *ASTPrinter) VisitVariable(expr *ast.Variable) string {
	return fmt.Sprintf("%vVAR('%v')", p.pad(), expr.Name)
}

func (p *ASTPrinter) VisitAssignment(expr *ast.Assignment) string {
	p.indent()
	pad := p.pad()
	val, err := ast.Accept[string](expr.Value, p)
	if err != nil {
		return err.Error()
	}
	p.dedent()
	return fmt.Sprintf("%vASSIGN(\n%v'%v'\n%v\n%v)", p.pad(), pad, expr.Name, val, p.pad())
}

func (p *ASTPrinter) VisitLiteral(expr *ast.Literal) string {
	return fmt.Sprintf("%v%T(%v)", p.pad(), expr.Value, expr.Value)
}

func (p *ASTPrinter) VisitGrouping(expr *ast.Grouping) string {
	p.indent()
	val, err := ast.Accept[string](expr.Expression, p)
	if err != nil {
		return err.Error()
	}
	p.dedent()
	return fmt.Sprintf("%vGROUP(\n%v\n%v)", p.pad(), val, p.pad())
}

func (p *ASTPrinter) VisitUnary(expr *ast.Unary) string {
	p.indent()
	pad := p.pad()
	val, err := ast.Accept[string](expr.Expression, p)
	if err != nil {
		return err.Error()
	}
	p.dedent()
	return fmt.Sprintf("%vUNARY(\n%v%v\n%v\n%v)", p.pad(), pad, expr.Operator.Lexeme(), val, p.pad())
}

func (p *ASTPrinter) VisitBinary(expr *ast.Binary) string {
	p.indent()
	pad := p.pad()
	valLeft, err := ast.Accept[string](expr.Left, p)
	if err != nil {
		return err.Error()
	}
	valRight, err := ast.Accept[string](expr.Right, p)
	if err != nil {
		return err.Error()
	}
	p.dedent()

	return fmt.Sprintf("%vBINARY(\n%v%v\n%v\n%v\n%v)", p.pad(), pad, expr.Operator.Lexeme(), valLeft, valRight, p.pad())

}

func NewASTPrinter() ast.Visitor[string] {
	return &ASTPrinter{}
}
