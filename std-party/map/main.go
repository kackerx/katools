package main

import (
    "fmt"
    "unsafe"
)

type S1 struct {
    f1 int16
    f2 int32
}

type Person interface {
    GetName() string
}

func main() {
    fmt.Println(unsafe.Sizeof(S1{}))
    s := S1{
        f1: 0,
        f2: 0,
    }

    fmt.Println(s.f2)
}
