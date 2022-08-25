package main

import "fmt"

func main() {

    a := make([]int, 5, 5)
    a = append(a, 1, 1, 1, 1, 1, 1)

    fmt.Println(cap(a))
}
