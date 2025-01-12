package lsp

import (
	"log"
	"os"
)

func HandleShutdown(state *ServerState, logger *log.Logger) {
	logger.Println("shutting down GDX")
	state.Shutdown = true
}

func HandleExit(logger *log.Logger) {
	logger.Println("exiting server")
	os.Exit(0)
}
