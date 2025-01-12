package rpc_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"gdx/rpc"
)

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 13\r\n\r\n{\"Foo\":\"bar\"}"

	actual, err := rpc.EncodeMessage(struct{ Foo string }{Foo: "bar"})
	assert.NoError(t, err)
	assert.Equal(t, expected, actual, "data returned is different")
}

func TestDecodeMessage(t *testing.T) {
	message := []byte("Content-Length: 16\r\n\r\n{\"method\":\"foo\"}")

	method, _, err := rpc.DecodeMessage(message, log.Default())
	assert.NoError(t, err)
	assert.Equal(t, "foo", method, "decoded data is not equal")
}
