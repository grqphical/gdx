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

func HandleTextDocumentOpen(contents []byte, logger *log.Logger) error {
	var msg DidOpenTextDocumentNotification

	if err := json.Unmarshal(contents, &msg); err != nil {
		return err
	}

	logger.Printf("recieved textDocument/open. URI: %s\nContents: %s", msg.Params.TextDocument.URI, msg.Params.TextDocument.Text)

	return nil
}
