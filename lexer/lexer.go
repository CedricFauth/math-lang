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
	IDENTIFIER
	EQUAL
	LESS
	GREATER
	BANG
	TRUE
	FALSE
	//EOL
	BANG_EQUAL
	EQUAL_EQUAL
	GREATER_EQUAL
	LESS_EQUAL
)

var keywords = map[string]TokenType{
	"TRUE":  TRUE,
	"FALSE": FALSE,
}

type Token struct {
	tokenType TokenType
	lexeme    string
	value     any
}

func NewTestToken(t TokenType) *Token {
	return &Token{t, t.String(), nil}
}

func (t *Token) String() string {
	switch t.value.(type) {
	case float64:
		return fmt.Sprintf("&Token(%v, '%v', %.2f)", t.tokenType, t.lexeme, t.value)
	default:
		return fmt.Sprintf("&Token(%v, '%v', %v)", t.tokenType, t.lexeme, t.value)
	}
}

func (t *Token) Lexeme() string {
	return t.lexeme
}

func (t *Token) Type() TokenType {
	return t.tokenType
}

func (t *Token) Value() any {
	return t.value
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
		char, err := lexer.peek(0)
		if err != nil {
			break
		}
		switch char {
		case ' ', '\t', '\n':
			lexer.advance()
		/*case '\n':
		tokenErr = lexer.consumeToken(EOL)*/
		case '=':
			c, _ := lexer.peek(1)
			if c == '=' {
				tokenErr = lexer.consumeToken(EQUAL_EQUAL)
			} else {
				tokenErr = lexer.consumeToken(EQUAL)
			}
		case '!':
			c, _ := lexer.peek(1)
			if c == '=' {
				tokenErr = lexer.consumeToken(BANG_EQUAL)
			} else {
				tokenErr = lexer.consumeToken(BANG)
			}
		case '<':
			c, _ := lexer.peek(1)
			if c == '=' {
				tokenErr = lexer.consumeToken(LESS_EQUAL)
			} else {
				tokenErr = lexer.consumeToken(LESS)
			}
		case '>':
			c, _ := lexer.peek(1)
			if c == '=' {
				tokenErr = lexer.consumeToken(GREATER_EQUAL)
			} else {
				tokenErr = lexer.consumeToken(GREATER)
			}
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
			} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
				tokenErr = lexer.consumeToken(IDENTIFIER)
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

	switch tokenType {
	case NUMBER:
		startIndex := lexer.index
		length := 0
		char, err := lexer.peek(0)
		//isFloat := false
		for {
			if err != nil {
				break
			}
			if char >= '0' && char <= '9' {
				length++
			} else if char == '.' {
				length++
				//isFloat = true
			} else {
				break
			}
			char, err = lexer.advance()
		}
		lexeme = lexer.code[startIndex : startIndex+length]
		//if isFloat {
		value, err = strconv.ParseFloat(lexeme, 64)
		/*} else {
			value, err = strconv.ParseInt(lexeme, 10, 64)
		}*/
		if err != nil {
			return err
		}
	case IDENTIFIER:
		start := lexer.index
		for {
			char, _ := lexer.advance()
			if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z' || char >= '0' && char <= '9')) {
				break
			}
		}
		end := lexer.index
		lexeme = lexer.code[start:end]
		if t, ok := keywords[lexeme]; ok {
			tokenType = t
		}
	case BANG_EQUAL, EQUAL_EQUAL, GREATER_EQUAL, LESS_EQUAL:
		lexeme = lexer.code[lexer.index : lexer.index+2]
		lexer.advance()
		lexer.advance()
	default:
		char, _ := lexer.peek(0)
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

func (lexer *Lexer) peek(offset int) (byte, error) {
	if lexer.index+offset < len(lexer.code) && lexer.index+offset >= 0 {
		return lexer.code[lexer.index+offset], nil
	} else {
		return 0, errors.New("end of input")
	}
}

func (lexer *Lexer) advance() (byte, error) {
	lexer.index++
	return lexer.peek(0)
}
