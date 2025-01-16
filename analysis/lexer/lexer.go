package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType = int

const (
	// Literals
	TypeIdentifier TokenType = iota
	TypeString
	TypeNumber

	// Keywords
	TokenIf
	TokenElif
	TokenElse
	TokenFor
	TokenWhile
	TokenMatch
	TokenWhen
	TokenBreak
	TokenContinue
	TokenPass
	TokenReturn
	TokenClass
	TokenClassName
	TokenExtends
	TokenIs
	TokenIn
	TokenAs
	TokenSelf
	TokenSuper
	TokenSignal
	TokenFunc
	TokenStatic
	TokenConst
	TokenEnum
	TokenVar
	TokenBreakpoint
	TokenPreload
	TokenAwait
	TokenYield
	TokenAssert
	TokenVoid
	TokenPI
	TokenTAU
	TokenINF
	TokenNAN
)

type SyntaxError struct {
	Line    int
	Column  int
	Message string
}

func NewSyntaxError(line int, col int, message string) SyntaxError {
	return SyntaxError{
		Line:    line,
		Column:  col,
		Message: message,
	}
}

func (s SyntaxError) Error() string {
	return fmt.Sprintf("syntax error on line %d, column %d: %s", s.Line+1, s.Column+1, s.Message)
}

type Token struct {
	Type  TokenType
	Value string
}

func makeNumber(line string, colNumber *int, lineNumber int) (Token, error) {
	dotCount := 0
	numberString := ""

	for i := 0; *colNumber+i < len(line); i++ {
		char := rune(line[*colNumber+i])

		if char == '.' {
			// Can't have more than one dot in a float literal
			if dotCount == 1 {
				return Token{}, NewSyntaxError(lineNumber, *colNumber+i, "invalid float literal")
			}

			dotCount++
			numberString += string(char)
		} else if unicode.IsDigit(char) {
			numberString += string(char)
		} else {
			break
		}
	}

	*colNumber += len(numberString)

	return Token{
		Type:  TypeNumber,
		Value: numberString,
	}, nil
}

func ScanSource(source string) ([]Token, error) {
	var tokens []Token = make([]Token, 0)

	lines := strings.Split(source, "\n")

	for lineNumber, line := range lines {
		colNumber := 0
		var char rune

	charLoop:
		for colNumber < len(line) {
			char = rune(line[colNumber])

			switch char {
			// line/rest of line is a comment so ignore the rest of it
			case '#':
				break charLoop
			case ' ', '\t', '\x00':
				break
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				fmt.Printf("character: %c\n", char)
				token, err := makeNumber(line, &colNumber, lineNumber)
				if err != nil {
					return nil, err
				}

				tokens = append(tokens, token)

				break
			default:
				return nil, NewSyntaxError(lineNumber, colNumber, fmt.Sprintf("unknown token %q", char))
			}
			colNumber += 1
		}
	}

	return tokens, nil
}
