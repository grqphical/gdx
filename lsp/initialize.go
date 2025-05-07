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

type WorkspaceFolder struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
}

type InitializeRequestParams struct {
	ClientInfo struct {
		Name    string `json:"name"`
		Version string `json:"version,omitempty"`
	} `json:"clientInfo"`
	RootPath string `json:"rootPath"`
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
}

type ServerCapabilities struct {
	TextDocumentSync   int               `json:"textDocumentSync"`
	CompletionProvider CompletionOptions `json:"completionProvider"`
}

func HandleInitialize(content []byte, logger *log.Logger, state *ServerState) error {
	var request InitializeRequest
	if err := json.Unmarshal(content, &request); err != nil {
		return err
	}

	logger.Printf(
		"connected to client %s, version %s, workspace %s\n",
		request.Params.ClientInfo.Name,
		request.Params.ClientInfo.Version,
		request.Params.RootPath,
	)

	state.WorkspacePath = request.Params.RootPath

	var response InitializeResponse = InitializeResponse{
		Result: InitializeResult{
			ServerInfo: ServerInfo{
				Name:    ServerName,
				Version: version.Version,
			},
			Capabilities: ServerCapabilities{
				TextDocumentSync:   1,
				CompletionProvider: CompletionOptions{},
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
