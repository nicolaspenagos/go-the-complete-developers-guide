package main

import (
	"io"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, file)
}
