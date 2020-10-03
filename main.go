package main

import (
	"fmt"
	"net/http"
	"os"
)

func run() error {
	srv := newServer()
	srv.store = &dbStore{}
	err := srv.store.Open()
	if err != nil {
		return err
	}
	defer srv.store.Close()

	http.HandleFunc("/", srv.serveHTTP)
	fmt.Println("Server listening on port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
