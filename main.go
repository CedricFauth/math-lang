package main

import (
	"fmt"
	"os"

	"github.com/CedricFauth/math-lang-go/interpreter"
	"github.com/CedricFauth/math-lang-go/lexer"
	"github.com/CedricFauth/math-lang-go/parser"
)

func failed(what ...any) {
	fmt.Println(what...)
	os.Exit(1)
}

func main() {

	//inputs := []string{"abc = 1", "a=abc+1"}
	inputs := []string{"abc TRUE FALSE == >= <= != < > = !"}

	interpet := interpreter.New()

	for _, input := range inputs {
		lexer := lexer.New(input)

		err := lexer.Scan()
		if err != nil {
			failed(err)
		}
		fmt.Printf("%v\n", lexer.Tokens())

		parser := parser.New(lexer.Tokens())
		expr := parser.Parse(lexer.Tokens())

		res := interpreter.Evaluate(expr, interpet)

		fmt.Printf("%v (%T)\n", res, res)
	}
	//parser.TestASTPrinter1()
	//parser.TestASTPrinter2()
}
