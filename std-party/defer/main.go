package main

import (
    "fmt"
    "time"
)

func main() {
    Foo()
}

func Foo() {
    now := time.Now()
    //defer fmt.Println(time.Since(now))
    defer func() {
        fmt.Println(time.Since(now))
    }()

    time.Sleep(time.Second * 2)
    fmt.Println(now)

}

func bar() string {
    return ""
}

type T interface {
    Test() string
}

type myt struct{}

func (m *myt) Test() string {
    return ""
}

func NewMyt() T {
    return &myt{}
}
