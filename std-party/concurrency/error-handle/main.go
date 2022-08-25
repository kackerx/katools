package main

import (
    "fmt"
    "net/http"
)

func main() {

    done := make(chan struct{})
    urls := []string{"https://www.baidu.com", "http://badhost"}
    for v := range GetResp(done, urls...) {
        fmt.Println(v.StatusCode)
    }
}

func GetResp(done chan struct{}, urls ...string) chan *http.Response {
    out := make(chan *http.Response)
    go func() {
        defer close(out)
        for _, url := range urls {
            res, err := http.Get(url)
            if err != nil {
                fmt.Println(err) // 没有办法传递回来error, 使用额外的结构体去封装, 把错误和并发分开
                continue
            }

            select {
            case <-done:
                return
            case out <- res:
            }
        }
    }()
    return out
}
