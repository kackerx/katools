package main

import (
    "fmt"
    "sync"
)

func fanIn(chans ...chan int) chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    for _, c := range chans {
        wg.Add(1)
        go func(in chan int) {
            defer wg.Done()
            for v := range in {
                out <- v
            }
        }(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}

func main() {
    data := []int{1, 2, 3}
    c1 := add(data)
    c2 := mul(data)
    out := fanIn(c1, c2)

    for v := range out {
        fmt.Println(v)
    }

}

func add(datas []int) chan int {
    out := make(chan int, len(datas))

    go func() {
        defer close(out)
        for _, v := range datas {
            out <- v + 1
        }
    }()

    return out
}

func mul(datas []int) chan int {
    out := make(chan int, len(datas))

    go func() {
        defer close(out)
        for _, v := range datas {
            out <- v * v
        }
    }()

    return out
}
