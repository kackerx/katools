package randombalance

import (
    "fmt"
    "testing"
    "time"
)

func TestRandom(t *testing.T) {

    ch := make(chan struct{})
    go func() {
        time.Sleep(time.Second * 3)
        ch <- struct{}{}
    }()
    fmt.Println("----")
    for {
        select {
        case <-ch:
            fmt.Println("get ch break")
            break
        default:
            fmt.Println("de")
            time.Sleep(time.Second)
        }
    }
}
