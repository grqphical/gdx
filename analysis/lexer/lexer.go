package lexer

import (
	"fmt"
)

type TokenType = int

const (
	// Literals
	TokenIdentifier TokenType = iota
	TokenString
	TokenNumber

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
	TokenPower

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

// check if we are at the end of the source
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// advance the scanner
func (s *Scanner) advance() rune {
	char := rune(s.source[s.current])
	s.current += 1

	return char
}

// add a token to the list of scanned tokens
func (s *Scanner) addToken(tokenType TokenType) {
	text := s.source[s.start:s.current]

	s.tokens = append(s.tokens, Token{
		Type:  tokenType,
		Value: text,
		Line:  s.line,
	})
}

func (s *Scanner) addTokenWithValue(tokenType TokenType, value string) {
	s.tokens = append(s.tokens, Token{
		Type:  tokenType,
		Value: value,
		Line:  s.line,
	})

}

// check is the next character equals the given character
// and advances the scanner by one token if so
func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if rune(s.source[s.current]) != expected {
		return false
	}

	s.current += 1
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\x00'
	}
	return rune(s.source[s.current])
}

func (s *Scanner) ScanTokens() ([]Token, *LexicalError) {
	for !s.isAtEnd() {
		s.start = s.current
		c := s.advance()

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
			if s.match('=') {
				s.addToken(TokenMinusEqual)
			} else {
				s.addToken(TokenDash)
			}
		case '+':
			if s.match('=') {
				s.addToken(TokenPlusEqual)
			} else {
				s.addToken(TokenPlus)
			}
		case '=':
			if s.match('=') {
				s.addToken(TokenEqualsEquals)
			} else {
				s.addToken(TokenEquals)
			}
		case '!':
			if s.match('=') {
				s.addToken(TokenNotEqual)
			} else {
				s.addToken(TokenBang)
			}

		case '/':
			if s.match('=') {
				s.addToken(TokenDivideEqual)
			} else {
				s.addToken(TokenSlash)
			}
		case '*':
			if s.match('*') {
				if s.match('=') {
					s.addToken(TokenPowerEqual)
				} else {
					s.addToken(TokenPower)
				}
			} else if s.match('=') {
				s.addToken(TokenTimesEqual)
			} else {
				s.addToken(TokenStar)
			}
		case '%':
			if s.match('=') {
				s.addToken(TokenModEqual)
			} else {
				s.addToken(TokenPercent)
			}
		case '&':
			if s.match('&') {
				s.addToken(TokenBooleanAnd)
			} else {
				s.addToken(TokenAmpersand)
			}
		case '|':
			if s.match('|') {
				s.addToken(TokenBooleanOr)
			} else if s.match('=') {
				s.addToken(TokenOrEqual)
			} else {
				s.addToken(TokenOr)
			}
		case ',':
			s.addToken(TokenComma)
		case '>':
			if s.match('=') {
				s.addToken(TokenGreaterOrEqual)
			} else if s.match('>') {
				if s.match('=') {
					s.addToken(TokenRShiftEqual)
				} else {
					s.addToken(TokenShiftRight)
				}
			} else {
				s.addToken(TokenGreater)
			}
		case '<':
			if s.match('=') {
				s.addToken(TokenLessOrEqual)
			} else if s.match('<') {
				if s.match('=') {
					s.addToken(TokenLShiftEqual)
				} else {
					s.addToken(TokenShiftLeft)
				}
			} else {
				s.addToken(TokenLess)
			}
		case '^':
			if s.match('=') {
				s.addToken(TokenXorEqual)
			} else {
				s.addToken(TokenXOR)
			}
		case ' ', '\t', '\r':
			break
		case '\n':
			s.line += 1
			break

		// comments
		case '#':
			for s.peek() != '\n' {
				s.advance()
			}
			break
		// string literals
		case '"', '\'':
			if s.current+2 < len(s.source) && // Check if we have enough characters ahead
				rune(s.source[s.current]) == c &&
				rune(s.source[s.current+1]) == c { // Check for triple-quoted strings
				// Advance past the opening quotes
				s.advance()
				s.advance()

				for {
					if s.isAtEnd() {
						return nil, NewLexicalError(s.line, "unterminated triple-quoted string")
					}

					// Check for closing triple quotes
					if s.current+2 < len(s.source) &&
						rune(s.source[s.current]) == c &&
						rune(s.source[s.current+1]) == c &&
						rune(s.source[s.current+2]) == c {
						// Advance past the closing quotes
						s.advance()
						s.advance()
						s.advance()
						break
					}

					if s.peek() == '\n' {
						s.line++ // Track line numbers for multiline strings
					}
					s.advance()
				}

				// Extract the value inside the triple quotes
				value := s.source[s.start+3 : s.current-3]
				s.addTokenWithValue(TokenString, value)
			} else { // Handle single-quoted strings as usual
				for s.peek() != c && !s.isAtEnd() {
					if s.peek() == '\n' {
						return nil, NewLexicalError(s.line, "unterminated string due to newline")
					}
					s.advance()
				}

				if s.isAtEnd() {
					return nil, NewLexicalError(s.line, "unterminated string")
				}

				// Advance for the closing quote
				s.advance()

				value := s.source[s.start+1 : s.current-1]
				s.addTokenWithValue(TokenString, value)
			}
		default:
			return nil, NewLexicalError(s.line, fmt.Sprintf("unknown token '%s'", s.source[s.start:s.current]))
		}

	}

	return s.tokens, nil
}
