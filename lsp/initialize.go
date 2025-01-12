package lsp

import (
	"encoding/json"
	"fmt"
	"gdx/rpc"
	"log"
)

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

type InitializeResponse struct {
	ResponseMessage
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	ServerInfo   ServerInfo         `json:"serverInfo"`
	Capabilities ServerCapabilities `json:"capabilities"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ServerCapabilities struct {
	TextDocumentSync int `json:"textDocumentSync"`
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

	var response InitializeResponse = InitializeResponse{
		Result: InitializeResult{
			ServerInfo: ServerInfo{
				Name:    ServerName,
				Version: ServerVersion,
			},
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1,
			},
		},
		ResponseMessage: ResponseMessage{
			RPC: "2.0",
			ID:  request.ID,
		},
	}
	msg, err := rpc.EncodeMessage(response)
	if err != nil {
		return err
	}

	logger.Println(msg)

	fmt.Printf("%s", msg)

	return nil
}
