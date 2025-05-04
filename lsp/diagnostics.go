package lsp

import (
	"errors"
	"fmt"
	"gdx/analysis/lexer"
	"gdx/rpc"
	"log"
)

type Serverity = int

const (
	SeverityError       Serverity = 1
	SeverityWarning     Serverity = 2
	SeverityInformation Serverity = 3
	SeverityHint        Serverity = 4
)

type Position struct {
	Line      uint `json:"line"`
	Character uint `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Diagnostic struct {
	Range     Range     `json:"range"`
	Serverity Serverity `json:"serverity"`
	Source    string    `json:"source"`
	Message   string    `json:"message"`
}

type PublishDiagnosticParams struct {
	URI         string       `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type PublishDiagnosticNotification struct {
	Notification
	Method string                  `json:"method"`
	Params PublishDiagnosticParams `json:"params"`
}

func RunDiagnostics(serverState *ServerState, logger *log.Logger, documentURI string) error {
	source, ok := serverState.Files[documentURI]
	if !ok {
		logger.Println("error: invalid file URI while running diagnostics")
		return errors.New("invalid file URI")
	}

	logger.Printf("running diagnostics on '%s'\n", documentURI)

	scanner := lexer.NewScanner(source)

	_, err := scanner.ScanTokens()

	var diagnostics []Diagnostic = make([]Diagnostic, 0)

	if err != nil {
		var lerr *lexer.LexicalError = err.(*lexer.LexicalError)

		diagnostics = []Diagnostic{
			{
				Range: Range{
					Start: Position{
						Line:      uint(lerr.Line),
						Character: uint(lerr.StartChar),
					},
					End: Position{
						Line:      uint(lerr.Line),
						Character: uint(lerr.EndChar),
					},
				},
				Serverity: SeverityError,
				Source:    "gdx",
				Message:   lerr.Message,
			},
		}
	}

	params := PublishDiagnosticParams{
		URI:         documentURI,
		Diagnostics: diagnostics,
	}

	payload := PublishDiagnosticNotification{
		Method: "textDocument/publishDiagnostics",
		Params: params,
	}

	encodedPayload, err := rpc.EncodeMessage(payload)
	if err != nil {
		return err
	}

	fmt.Print(encodedPayload)

	return nil

}
