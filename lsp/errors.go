package lsp

const (
	ErrCodeParseError     int = -32700
	ErrCodeInvalidRequest int = -32600
	ErrCodeMethodNotFound int = -32601
	ErrCodeInvalidParams  int = -32602
	ErrCodeInternalError  int = -32603
)

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
