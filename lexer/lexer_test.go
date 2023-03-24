package lexer

import (
	"fmt"
	"testing"
)

func TestScanEmpty(t *testing.T) {

	testCases := []string{
		"",
		" ",
		"\t",
	}

	for i, e := range testCases {
		l := Lexer{code: e}
		l.Scan()
		if l.Tokens() != nil {
			t.Fatal("Test failed:", i)
		}
	}

}

func TestScanOperators(t *testing.T) {

	testCases := map[string]([]*Token){
		"+": {&Token{PLUS, "+", nil}},
		"(+*/-)": {
			&Token{PAREN_OPEN, "(", nil},
			&Token{PLUS, "+", nil},
			&Token{STAR, "*", nil},
			&Token{SLASH, "/", nil},
			&Token{MINUS, "-", nil},
			&Token{PAREN_CLOSE, ")", nil},
		},
	}

	for k, v := range testCases {
		l := Lexer{code: k}
		l.Scan()
		result := fmt.Sprint(l.Tokens())
		expected := fmt.Sprint(v)
		if result != expected {
			t.Fatal("Test failed:", expected, result)
		}
	}
}

func TestScanNumbers(t *testing.T) {

	testCases := map[string]([]*Token){
		"0":    {&Token{NUMBER, "0", 0}},
		"0.0":  {&Token{NUMBER, "0.0", 0.0}},
		"123":  {&Token{NUMBER, "123", 123}},
		"38.":  {&Token{NUMBER, "38.", 38.0}},
		".13":  {&Token{NUMBER, ".13", 0.13}},
		"9.87": {&Token{NUMBER, "9.87", 9.87}},
	}

	for k, v := range testCases {
		l := Lexer{code: k}
		l.Scan()
		result := fmt.Sprint(l.Tokens())
		expected := fmt.Sprint(v)
		if result != expected {
			t.Fatal("Test failed:", expected, result)
		}
	}
}

func TestScanFull(t *testing.T) {

	testCases := map[string]([]*Token){
		"12+3": {
			&Token{NUMBER, "12", 12},
			&Token{PLUS, "+", nil},
			&Token{NUMBER, "3", 3},
		},
		"33.33*3+0.1": {
			&Token{NUMBER, "33.33", 33.33},
			&Token{STAR, "*", nil},
			&Token{NUMBER, "3", 3},
			&Token{PLUS, "+", nil},
			&Token{NUMBER, "0.1", 0.1},
		},
		"1+ ( 2 * 3.0 ) -(-2.5/5)": {
			&Token{NUMBER, "1", 1},
			&Token{PLUS, "+", nil},
			&Token{PAREN_OPEN, "(", nil},
			&Token{NUMBER, "2", 2},
			&Token{STAR, "*", nil},
			&Token{NUMBER, "3.0", 3.0},
			&Token{PAREN_CLOSE, ")", nil},
			&Token{MINUS, "-", nil},
			&Token{PAREN_OPEN, "(", nil},
			&Token{MINUS, "-", nil},
			&Token{NUMBER, "2.5", 2.5},
			&Token{SLASH, "/", nil},
			&Token{NUMBER, "5", 5},
			&Token{PAREN_CLOSE, ")", nil},
		},
	}

	for k, v := range testCases {
		l := Lexer{code: k}
		l.Scan()
		result := fmt.Sprint(l.Tokens())
		expected := fmt.Sprint(v)
		if result != expected {
			t.Fatal("Test failed:\n", expected, "\n", result)
		}
	}
}

func TestScanPanic(t *testing.T) {

	testCases := []string{
		".",
		"2+3_000",
		"10O0",
	}

	for i, e := range testCases {
		l := Lexer{code: e}
		err := l.Scan()
		if err == nil {
			t.Fatal("Test failed:", i)
		}
	}
}
