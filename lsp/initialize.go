package lsp

import (
	"encoding/json"
	"fmt"
	"gdx/rpc"
	"log"

	"gdx/version"
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

type CompletionOptions struct {
	ResolveProvider   bool     `json:"resolveProvider"`
	TriggerCharacters []string `json:"triggerCharacters"`
}

type ServerCapabilities struct {
	TextDocumentSync   int               `json:"textDocumentSync"`
	CompletionProvider CompletionOptions `json:"completionProvider"`
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
				Version: version.Version,
			},
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1,
				CompletionProvider: CompletionOptions{
					ResolveProvider: true,
					TriggerCharacters: []string{
						// Lowercase letters
						"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
						// Uppercase letters
						"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
						// Numbers
						"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
					},
				},
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

	fmt.Printf("%s", msg)

	return nil
}
