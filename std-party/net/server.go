package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "hello world")
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	server.ListenAndServe()

}
