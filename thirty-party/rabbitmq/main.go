package main

import "fmt"

func main() {
    var m Myint
    m = "kacker"
    fmt.Println(len(m))
}

func publist() {}

type Myint string


