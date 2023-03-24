package main

import (
	"fmt"
	"os"

	"github.com/CedricFauth/math-lang-go/lexer"
)

func failed(what ...any) {
	fmt.Println(what...)
	os.Exit(1)
}

func main() {
	input := ".12+-"
	lexer := lexer.New(input)

	err := lexer.Scan()
	if err != nil {
		failed(err)
	}
	fmt.Printf("%v\n", lexer.Tokens())
}
