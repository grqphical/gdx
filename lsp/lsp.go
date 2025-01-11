package lsp

import (
	"encoding/json"
	"log"
)

type RequestMessage struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
}

type ResponseMessage struct {
	ID     int `json:"id"`
	Result any `json:"result,omitempty"`
	Error  any `json:"error,omitempty"`
}

type InitializeRequest struct {
	RequestMessage
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo struct {
		Name    string `json:"name"`
		Version string `json:"version,omitempty"`
	} `json:"clientInfo"`
}

func HandleInitialize(content []byte, logger *log.Logger) error {
	var request InitializeRequest
	if err := json.Unmarshal(content, &request); err != nil {
		return err
	}

	logger.Printf(
		"connected to client %s, version %s\n",
		request.Params.ClientInfo.Name,
		request.Params.ClientInfo.Version,
	)

	return nil
}
