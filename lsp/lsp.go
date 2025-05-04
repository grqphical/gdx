package lsp

const ServerName string = "gdx"

type FileState struct {
	Contents []string
}

type ServerState struct {
	Shutdown bool
	Files    map[string]FileState
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
