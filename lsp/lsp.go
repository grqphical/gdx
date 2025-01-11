package lsp

type RequestMessage struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
	Params any    `json:"params,omitempty"`
}

type ResponseMessage struct {
	ID     int `json:"id"`
	Result any `json:"result,omitempty"`
	Error  any `json:"error,omitempty"`
}
