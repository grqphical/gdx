package lexer

import (
	"fmt"
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

	// Single Character Tokens
	TokenLParen
	TokenRParen
	TokenLBracket
	TokenRBracket
	TokenPeriod
	TokenTilda
	TokenDash
	TokenPlus
	TokenEquals
	TokenBang
	TokenSlash
	TokenStar
	TokenPercent
	TokenAmpersand
	TokenOr
	TokenComma
	TokenGreater
	TokenLess
	TokenXOR

	// Two-Character Operators
	TokenEqualsEquals
	TokenNotEqual
	TokenGreaterOrEqual
	TokenLessOrEqual
	TokenShiftRight
	TokenShiftLeft
	TokenBooleanAnd
	TokenBooleanOr
	TokenPlusEqual
	TokenMinusEqual
	TokenTimesEqual
	TokenDivideEqual
	TokenPowerEqual
	TokenModEqual
	TokenAndEqual
	TokenOrEqual
	TokenXorEqual
	TokenRShiftEqual
	TokenLShiftEqual

	TokenEOF
)

type LexicalError struct {
	Line    int
	Column  int
	Message string
}

func NewLexicalError(line int, message string) *LexicalError {
	return &LexicalError{
		Line:    line,
		Message: message,
	}
}

func (s LexicalError) Error() string {
	return fmt.Sprintf("lexical error at line %d: %s", s.Line, s.Message)
}

type Token struct {
	Type  TokenType
	Value string
	Line  int
}

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	tokens  []Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
		tokens:  make([]Token, 0),
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() rune {
	char := rune(s.source[s.current])
	s.current += 1

	return char
}

func (s *Scanner) addToken(tokenType TokenType) {
	text := s.source[s.start : s.current+1]

	s.tokens = append(s.tokens, Token{
		Type:  tokenType,
		Value: text,
		Line:  s.line,
	})
}

func (s *Scanner) ScanTokens() ([]Token, *LexicalError) {
	c := s.advance()

	for !s.isAtEnd() {
		s.start = s.current

		switch c {
		case '(':
			s.addToken(TokenLParen)
		case ')':
			s.addToken(TokenRParen)
		case '[':
			s.addToken(TokenLBracket)
		case ']':
			s.addToken(TokenRBracket)
		case '.':
			s.addToken(TokenPeriod)
		case '~':
			s.addToken(TokenTilda)
		case '-':
			s.addToken(TokenDash)
		case '+':
			s.addToken(TokenPlus)
		case '=':
			s.addToken(TokenEquals)
		case '!':
			s.addToken(TokenBang)
		case '/':
			s.addToken(TokenSlash)
		case '*':
			s.addToken(TokenStar)
		case '%':
			s.addToken(TokenPercent)
		case '&':
			s.addToken(TokenAmpersand)
		case '|':
			s.addToken(TokenOr)
		case ',':
			s.addToken(TokenComma)
		case '>':
			s.addToken(TokenGreater)
		case '<':
			s.addToken(TokenLess)
		case '^':
			s.addToken(TokenXOR)
		default:
			// Handle unsupported characters if needed
		}

	}

	return s.tokens, nil
}
