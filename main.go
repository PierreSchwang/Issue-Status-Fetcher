package main

import (
	"io"
	"issue-status-fetcher/cmd"
	"log"
	"os"
)

func main() {
	// Add file based logging
	f, err := os.OpenFile("status.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Failed to close file stream: ", err)
		}
	}(f)
	writer := io.MultiWriter(os.Stdout, f)
	log.SetOutput(writer)
	cmd.Run()
}
