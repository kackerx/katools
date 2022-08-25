package main

import (
    "fmt"
    "sync"
)

type N struct {
    heh string
}

func main() {
    var wg sync.WaitGroup

    ch := make(chan int, 10) // 信号量
    s := make([]int, 100)

    for i := 0; i < 100; i++ {
        ch <- i
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            s = append(s, i)
            <-ch
        }(i)
    }

    wg.Wait()
    fmt.Println(len(s))
}
