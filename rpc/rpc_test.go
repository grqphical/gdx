package rpc_test

import (
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
	expected := struct{ Foo string }{Foo: "bar"}
	message := []byte("Content-Length: 13\r\n\r\n{\"Foo\":\"bar\"}")

	var actual struct{ Foo string }
	assert.NoError(t, rpc.DecodeMessage(message, &actual))
	assert.Equal(t, expected, actual, "decoded data is not equal")
}
