package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {

}

func Get() {
    request, err := http.NewRequest(http.MethodGet, "http://localhost:8888", nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    resp, err := http.DefaultClient.Do(request)
    if err != nil {
        fmt.Println(2)
        return
    }
    defer resp.Body.Close()

    respByte, err := ioutil.ReadAll(resp.Body)
    fmt.Println(3)
    fmt.Println(string(respByte), 1)
}
