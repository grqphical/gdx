package main

import (
	"bufio"
	"log"
	"os"

	"gdx/rpc"
)

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[gdx-server]", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	logger := getLogger("log.txt")

	logger.Println("starting GDX")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		_, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}

		logger.Printf("%s\n", contents)
	}
}
