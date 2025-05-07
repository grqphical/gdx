package lsp

import "gdx/analysis"

const ServerName string = "gdx"

type ServerState struct {
	Shutdown      bool
	WorkspacePath string
	Files         map[string]string
	ProjectConfig analysis.GodotProjectFile
}

type RequestMessage struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
}

type ResponseMessage struct {
	RPC   string `json:"jsonrpc"`
	ID    int    `json:"id"`
	Error any    `json:"error,omitempty"`
}

type Notification struct {
	Method string `json:"method"`
}
