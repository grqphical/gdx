package rpc_test

import (
	"bytes"
	"log"
	"testing"

	"gdx/rpc"
)

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 13\r\n\r\n{\"Foo\":\"bar\"}"

	actual, err := rpc.EncodeMessage(struct{ Foo string }{Foo: "bar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	message := []byte("Content-Length: 16\r\n\r\n{\"method\":\"foo\"}")
	logger := log.New(&bytes.Buffer{}, "", log.LstdFlags) // Use a buffer to avoid logging to stdout in tests

	method, _, err := rpc.DecodeMessage(message, logger)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if method != "foo" {
		t.Errorf("expected %q, got %q", "foo", method)
	}
}
