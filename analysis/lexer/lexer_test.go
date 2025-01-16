package lexer_test

import (
	"gdx/analysis/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanSource(t *testing.T) {
	// test comments
	testSource := "# this is a comment and should be ignored"
	tokens, err := lexer.ScanSource(testSource)
	assert.NoError(t, err)
	assert.Empty(t, tokens, "no tokens should be scanned")

	// test number literals
	testSource = "105\n0.15"
	expected := []lexer.Token{
		{
			Type:  lexer.TypeNumber,
			Value: "105",
		},
		{
			Type:  lexer.TypeNumber,
			Value: "0.15",
		},
	}
	tokens, err = lexer.ScanSource(testSource)
	assert.NoError(t, err)
	assert.ElementsMatch(t, expected, tokens, "tokens should match")

	// test error handling
	testSource = "INVALID"
	tokens, err = lexer.ScanSource(testSource)
	assert.Error(t, err)

	testSource = "0.a"
	tokens, err = lexer.ScanSource(testSource)
	assert.Error(t, err)
}
