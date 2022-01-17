package main

import (
    "fmt"
    "sync"
)

var once sync.Once
var u *User

func main() {
    a := NewObj()
    b := NewObj()
    
    fmt.Printf("%p\n", a)
    fmt.Printf("%p\n", b)
}

type User struct {
    Name string
}

func NewObj() *User {
    once.Do(func() {
        u = &User{"kacker"}
    })
    return u
}
