package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	// Content-Length: <number>
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}

// encodes an object into an RPC message
func EncodeMessage(toBeEncoded any) (string, error) {
	data, err := json.Marshal(toBeEncoded)
	if err != nil {
		return "", fmt.Errorf("unable to encode data to JSON: %s", err)
	}

	contentLength := len(data)

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", contentLength, data), nil
}

// decodes an RPC message returning which method it's using, the contents of the message, and an error if there is one
func DecodeMessage(msg []byte, logger *log.Logger) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("unable to find seperator in message")
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil

}
