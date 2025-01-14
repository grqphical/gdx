package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"gdx/lsp"
	"gdx/rpc"
)

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[gdx-server]", log.Ldate|log.Ltime|log.Lshortfile)
}

func handleMessage(method string, content []byte, logger *log.Logger, state *lsp.ServerState) error {
	if state.Shutdown {
		error := lsp.ResponseError{
			Code:    lsp.ErrCodeInvalidRequest,
			Message: "server is shutdown",
		}

		msg, _ := rpc.EncodeMessage(error)
		fmt.Printf("%s", msg)
		return nil
	} else {
		switch method {
		case "initialize":
			return lsp.HandleInitialize(content, logger)
		case "shutdown":
			lsp.HandleShutdown(state, logger)
			return nil
		case "exit":
			lsp.HandleExit(logger)
			return nil
		case "textDocument/didOpen":
			return lsp.HandleTextDocumentOpen(content, logger, state)
		case "textDocument/didChange":
			return lsp.HandleTextDocumentChange(content, logger, state)
		case "textDocument/completion":
			return lsp.HandleCompletion(content, logger)

		}

		return nil
	}
}

func main() {
	version := flag.Bool("version", false, "Prints the version")
	flag.Parse()

	if *version {
		fmt.Printf("GDX version: %s\n", lsp.ServerVersion)
		return
	}

	state := lsp.ServerState{
		Files: make(map[string]string),
	}
	logger := getLogger("/home/grqphical/development/go/gdx/log.txt")

	logger.Println("starting GDX")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg, logger)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}

		err = handleMessage(method, contents, logger, &state)
		if err != nil {
			logger.Printf("error while handling message: %s", err)
			continue
		}

	}
}
