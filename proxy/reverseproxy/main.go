package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

var (
    err  error
    addr = "127.0.0.1:2002"
)

func main() {
    rs1 := "http://127.0.0.1:2003/base"

    url1, err := url.Parse(rs1)
    if err != nil {
        log.Fatalln(err)
    }

    proxy := httputil.NewSingleHostReverseProxy(url1)

    log.Println("listening on", addr)

    log.Fatalln(http.ListenAndServe(addr, proxy))

}
