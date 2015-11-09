package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

type fileServer string

func (fs fileServer) Open(name string) (http.File, error) {
	return os.Open(path.Join(string(fs), name))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Please supply a directory to serve")
		os.Exit(1)
	}

	handler := http.FileServer(fileServer(os.Args[1]))
	http.ListenAndServe(":8080", handler)
}
