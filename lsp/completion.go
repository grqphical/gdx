package lsp

import (
	"encoding/json"
	"fmt"
	"gdx/rpc"
	"log"
)

type CompletionItemKind int

const (
	Text CompletionItemKind = iota + 1
	Method
	Function
	Constructor
	Field
	Variable
	Class
	Interface
	Module
	Property
	Unit
	Value
	Enum
	Keyword
	Snippet
	Color
	File
	Reference
	Folder
	EnumMember
	Constant
	Struct
	Event
	Operator
	TypeParameter
)

var keywords []string = []string{
	"if",
	"elif",
	"else",
	"for",
	"while",
	"match",
	"when",
	"break",
}

type CompletionRequest struct {
	RequestMessage
	Params CompletionRequestParams `json:"params"`
}

type CompletionRequestParams struct {
	Context CompletionContext `json:"context"`
}

type CompletionContext struct {
	TriggerKind      int    `json:"triggerKind"`
	TriggerCharacter string `json:"TriggerCharacter"`
}

type CompletionResponse struct {
	ResponseMessage
	Result []CompletionItem `json:"result"`
}

type CompletionItem struct {
	Label      string             `json:"label"`
	Kind       CompletionItemKind `json:"kind"`
	Detail     string             `json:"detail"`
	InsertText string             `json:"insertText"`
}

func generateCompletionItems(keywords []string) []CompletionItem {
	result := make([]CompletionItem, len(keywords))

	for _, keyword := range keywords {
		result = append(result, CompletionItem{
			Label:      keyword,
			Kind:       Keyword,
			Detail:     "a language keyword",
			InsertText: keyword,
		})
	}

	return result
}

func HandleCompletion(content []byte, logger *log.Logger) error {
	var request CompletionRequest
	if err := json.Unmarshal(content, &request); err != nil {
		return err
	}

	logger.Printf("recieved completion with trigger %d on '%s'\n", request.Params.Context.TriggerKind, request.Params.Context.TriggerCharacter)

	response := CompletionResponse{
		ResponseMessage: ResponseMessage{
			ID:  request.ID,
			RPC: "2.0",
		},
		Result: generateCompletionItems(keywords),
	}

	encodedResponse, err := rpc.EncodeMessage(response)
	if err != nil {
		return err
	}

	fmt.Print(encodedResponse)

	return nil
}
