package main

import (
	"fmt"
	"net/http"
)

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%v] %v\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
