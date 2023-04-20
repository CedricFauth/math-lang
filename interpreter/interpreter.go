package interpreter

import (
	"fmt"
	"os"

	"github.com/CedricFauth/math-lang-go/ast"
	"github.com/CedricFauth/math-lang-go/lexer"
)

type VarEnv map[string]any

func (env VarEnv) String() string {
	out := "VarEnv:\n"
	for k, v := range env {
		out += fmt.Sprintf("  %v: %v\n", k, v)
	}
	return out
}

type Interpreter struct {
	Env VarEnv
}

func (i *Interpreter) VisitVariable(expr *ast.Variable) any {
	val, present := i.Env[expr.Name]
	if !present {
		failed(fmt.Errorf("variable '%v' does not exist", expr.Name))
	}
	return val
}

func (i *Interpreter) VisitAssignment(expr *ast.Assignment) any {
	value := Evaluate(expr.Value, i)
	i.Env[expr.Name] = value
	fmt.Print(i.Env)
	return value
}

func (i *Interpreter) VisitLiteral(expr *ast.Literal) any {
	return expr.Value
}

func (i *Interpreter) VisitGrouping(expr *ast.Grouping) any {
	return Evaluate(expr.Expression, i)
}

func (i *Interpreter) VisitUnary(expr *ast.Unary) any {
	opType := expr.Operator.Type()
	if opType == lexer.MINUS {
		obj := Evaluate(expr.Expression, i)
		if v, ok := obj.(int64); ok {
			return -v
		} else if v, ok := obj.(float64); ok {
			return -v
		}
	}
	return nil
}

func (i *Interpreter) VisitBinary(expr *ast.Binary) any {
	left := Evaluate(expr.Left, i).(float64)
	right := Evaluate(expr.Right, i).(float64)

	opType := expr.Operator.Type()

	switch opType {
	case lexer.PLUS:
		return left + right
	case lexer.MINUS:
		return left - right
	case lexer.STAR:
		return left * right
	case lexer.SLASH:
		return left / right
	}
	return nil

}

func New() ast.Visitor[any] {
	i := &Interpreter{}
	i.Env = make(map[string]any)
	return i
}

func failed(e error) {
	fmt.Printf("%v\n", e)
	os.Exit(1)
}

func Evaluate(expr ast.Expr, i ast.Visitor[any]) any {
	obj, err := ast.Accept(expr, i)
	if err != nil {
		failed(err)
	}
	return obj
}
