package lexer

import (
	"errors"
	"fmt"
	"strconv"
)

type TokenType uint

const (
	PLUS TokenType = iota
	MINUS
	STAR
	SLASH
	PAREN_OPEN
	PAREN_CLOSE
	NUMBER
	EOL
)

type Token struct {
	tokenType TokenType
	lexeme    string
	value     any
}

func (t Token) String() string {
	switch t.value.(type) {
	case float64:
		return fmt.Sprintf("Token(%v, '%v', %.2f)", t.tokenType, t.lexeme, t.value)
	default:
		return fmt.Sprintf("Token(%v, '%v', %v)", t.tokenType, t.lexeme, t.value)
	}
}

type Lexer struct {
	tokens []*Token
	index  int
	code   string
}

func New(input string) *Lexer {
	return &Lexer{
		code: input,
	}
}

func (l *Lexer) Tokens() []*Token {
	return l.tokens
}

func (lexer *Lexer) Scan() error {

	for {
		var tokenErr error = nil
		char, err := lexer.peek()
		if err != nil {
			break
		}
		switch char {
		case ' ', '\t':
			lexer.advance()
		case '\n':
			tokenErr = lexer.consumeToken(EOL)
		case '+':
			tokenErr = lexer.consumeToken(PLUS)
		case '-':
			tokenErr = lexer.consumeToken(MINUS)
		case '*':
			tokenErr = lexer.consumeToken(STAR)
		case '/':
			tokenErr = lexer.consumeToken(SLASH)
		case '(':
			tokenErr = lexer.consumeToken(PAREN_OPEN)
		case ')':
			tokenErr = lexer.consumeToken(PAREN_CLOSE)
		default:
			if (char >= '0' && char <= '9') || char == '.' {
				tokenErr = lexer.consumeToken(NUMBER)
			} else {
				return errors.New(fmt.Sprintln("Unknown character: '"+string(char)+"' at ", lexer.index))
			}
		}
		if tokenErr != nil {
			return tokenErr
		}
	}
	return nil
}

func (lexer *Lexer) consumeToken(tokenType TokenType) error {
	lexeme := ""
	var value any = nil

	if tokenType == NUMBER {
		startIndex := lexer.index
		length := 0
		char, err := lexer.peek()
		isFloat := false
		for {
			if err != nil {
				break
			}
			if char >= '0' && char <= '9' {
				length++
			} else if char == '.' {
				length++
				isFloat = true
			} else {
				break
			}
			char, err = lexer.advance()
		}
		lexeme = lexer.code[startIndex : startIndex+length]
		if isFloat {
			value, err = strconv.ParseFloat(lexeme, 64)
		} else {
			value, err = strconv.ParseInt(lexeme, 10, 64)
		}
		if err != nil {
			return err
		}
	} else {
		char, _ := lexer.peek()
		lexeme = string(char)
		lexer.advance()
	}

	lexer.tokens = append(lexer.tokens, &Token{
		tokenType: tokenType,
		lexeme:    lexeme,
		value:     value,
	})
	return nil
}

func (lexer *Lexer) peek() (byte, error) {
	if lexer.index < len(lexer.code) {
		return lexer.code[lexer.index], nil
	} else {
		return 0, errors.New("end of input")
	}
}

func (lexer *Lexer) advance() (byte, error) {
	lexer.index++
	return lexer.peek()
}
