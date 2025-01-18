package lexer_test

import (
	"reflect"
	"testing"

	"gdx/analysis/lexer"
)

func TestScanTokens(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []lexer.Token
		hasError bool
	}{
		{
			name:  "Single character tokens",
			input: "()[]",
			expected: []lexer.Token{
				{Type: lexer.TokenLParen, Value: "(", Line: 1},
				{Type: lexer.TokenRParen, Value: ")", Line: 1},
				{Type: lexer.TokenLBracket, Value: "[", Line: 1},
				{Type: lexer.TokenRBracket, Value: "]", Line: 1},
			},
			hasError: false,
		},
		{
			name:  "Two-character tokens",
			input: "== != >= <=",
			expected: []lexer.Token{
				{Type: lexer.TokenEqualsEquals, Value: "==", Line: 1},
				{Type: lexer.TokenNotEqual, Value: "!=", Line: 1},
				{Type: lexer.TokenGreaterOrEqual, Value: ">=", Line: 1},
				{Type: lexer.TokenLessOrEqual, Value: "<=", Line: 1},
			},
			hasError: false,
		},
		{
			name:  "Mixed single and multi-character tokens",
			input: "+= -= * /",
			expected: []lexer.Token{
				{Type: lexer.TokenPlusEqual, Value: "+=", Line: 1},
				{Type: lexer.TokenMinusEqual, Value: "-=", Line: 1},
				{Type: lexer.TokenStar, Value: "*", Line: 1},
				{Type: lexer.TokenSlash, Value: "/", Line: 1},
			},
			hasError: false,
		},
		{
			name:     "Unrecognized token",
			input:    "?",
			expected: nil,
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scanner := lexer.NewScanner(test.input)
			tokens, err := scanner.ScanTokens()

			if (err != nil) != test.hasError {
				t.Fatalf("expected error: %v, got: %v", test.hasError, err)
			}

			if !test.hasError && !reflect.DeepEqual(tokens, test.expected) {
				t.Errorf("expected tokens: %v, got: %v", test.expected, tokens)
			}
		})
	}
}
