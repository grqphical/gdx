package main

import (
	"log"
	"os"
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
}
