package main

import (
	"fmt"
	"os"
)

func run() error {
	srv := newServer()
	srv.store = &dbStore{}
	err := srv.store.Open()
	if err != nil {
		return err
	}
	movies, err := srv.store.GetMovies()
	if err != nil {
		return err
	}
	fmt.Println(movies)
	defer srv.store.Close()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
