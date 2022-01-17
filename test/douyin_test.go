package test

import (
    "fmt"
    "testing"
    "time"
)

func TestTip(t *testing.T) {
    that := time.Date(2021, 11, 22, 0, 0, 0, 0, time.Local)
    res := time.Now().Sub(that).Hours()
    fmt.Println(res / 24)
}

func TestFoo(t *testing.T) {
    fmt.Println("kacker")
}



