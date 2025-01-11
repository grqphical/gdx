package main

import (
	"bufio"
	"errors"
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

func handleMessage(method string, content []byte, logger *log.Logger) error {
	switch method {
	case "initialize":
		return lsp.HandleInitialize(content, logger)
	default:
		return errors.New("invalid method")
	}
}

func main() {
	logger := getLogger("log.txt")

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

		err = handleMessage(method, contents, logger)
		if err != nil {
			logger.Printf("error while handling message: %s", err)
			continue
		}

	}
}
