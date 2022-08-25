package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    a := []int{1, 2, 3, 4, 5}
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < 100; i++ {
        fmt.Println(rand.Intn(len(a)))
    }

}
