package lsp

import (
	"encoding/json"
	"log"
)

type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageId string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

type DidOpenTextDocumentNotification struct {
	Notification
	Params DidOpenTextDocumentNotificationParams `json:"params"`
}

type DidOpenTextDocumentNotificationParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentNotificationParams `json:"params"`
}

type DidCloseTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type DidCloseTextDocumentNotification struct {
	Notification
	Params DidCloseTextDocumentParams `json:"params"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type DidChangeTextDocumentNotificationParams struct {
	TextDocument   TextDocumentIdentifier           `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

func HandleTextDocumentOpen(contents []byte, logger *log.Logger, state *ServerState) error {
	var msg DidOpenTextDocumentNotification

	if err := json.Unmarshal(contents, &msg); err != nil {
		return err
	}

	logger.Printf("recieved textDocument/open for %s\n", msg.Params.TextDocument.URI)

	state.Files[msg.Params.TextDocument.URI] = msg.Params.TextDocument.Text

	err := RunDiagnostics(state, logger, msg.Params.TextDocument.URI)
	if err != nil {
		return err
	}

	return nil
}

func HandleTextDocumentClose(contents []byte, logger *log.Logger, state *ServerState) error {
	var msg DidCloseTextDocumentNotification

	if err := json.Unmarshal(contents, &msg); err != nil {
		return err
	}

	logger.Printf("document %s closed", msg.Params.TextDocument.URI)

	delete(state.Files, msg.Params.TextDocument.URI)

	return nil
}

func HandleTextDocumentChange(contents []byte, logger *log.Logger, state *ServerState) error {
	var msg DidChangeTextDocumentNotification

	if err := json.Unmarshal(contents, &msg); err != nil {
		return err
	}

	logger.Printf("document %s changed", msg.Params.TextDocument.URI)

	state.Files[msg.Params.TextDocument.URI] = msg.Params.ContentChanges[0].Text

	err := RunDiagnostics(state, logger, msg.Params.TextDocument.URI)
	if err != nil {
		return err
	}

	return nil
}
