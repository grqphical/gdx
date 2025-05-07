package lsp

import "log"

type Registration struct {
	Id     string `json:"id"`
	Method string `json:"method"`
}

type RegistrationParams struct {
	Registrations []Registration `json:"registrations"`
}

func HandleInitialized(logger *log.Logger) {

}
