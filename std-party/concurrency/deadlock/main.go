package main

import (
    "fmt"
    "time"
)

func main() {

    done := make(chan int)
    res := foo(done, nil)
    
    go func() {
        time.Sleep(time.Second * 2)
        close(done)
    }()

    <-res
    fmt.Println("main dnoe")
}

func foo(done chan int, strings chan string) chan interface{} {
    outCh := make(chan interface{})
    go func() {
        defer close(outCh)
        for {
            select {
            case <-done:
                fmt.Println("get done", <-done)
                return
            case s := <-strings:
                fmt.Println(s)
            }
        }
    }()
    return outCh
}
