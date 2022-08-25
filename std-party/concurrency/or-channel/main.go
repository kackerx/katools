package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    <-or(sig(time.Second), sig(time.Second*2), sig(time.Second*10))
    fmt.Println(time.Since(start))
}

// 利用递归等, 等待所有的chan, 任一个结束orDone结束
func or(channels ...chan interface{}) chan interface{} {
    switch len(channels) {
    case 0:
        return nil
    case 1:
        return channels[0]
    }

    orDone := make(chan interface{})
    go func() {
        defer close(orDone)
        switch len(channels) {
        case 2:
            select {
            case <-channels[0]:
            case <-channels[1]:
            }
        default:
            m := len(channels) / 2
            select {
            case <-or(channels[m:]...):
            case <-or(channels[:m]...):
            }
        }
    }()
    return orDone
}

func sig(after time.Duration) chan interface{} {
    out := make(chan interface{})
    go func() {
        defer close(out)
        time.Sleep(after)
    }()
    return out
}
