package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://pkg.go.dev/net/http

type LogWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %w", err))
		os.Exit(1)
	}

	lw := LogWriter{}
	io.Copy(lw, resp.Body)
}

func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
