package main

import (
    "log"
    "net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    h(writer, request)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        writer.Write([]byte("Hello"))
    })
    log.Fatalln(http.ListenAndServe(":9999", mux))
}

func Hello(writer http.ResponseWriter, request *http.Request) {
    writer.Write([]byte("Hello"))
}
