package main

import (
    "fmt"
    "sync"
)

var (
    err error
    wg  sync.WaitGroup
)

type Result struct {
    Data int
    err  error
}

func main() {
    
    reCh := make(chan Result)
    
    for _, v := range map[string]int{"k": 1, "v": 2, "a": 4} {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            reCh <- Result{Data: i, err: nil}
        }(v)
    }
    
    go func() {
        wg.Wait()
        close(reCh)
    }()
    
    for v := range reCh {
        if v.err != nil {
            fmt.Println(v.err)
            continue
        }
        fmt.Println(v.Data)
    }
    
}
