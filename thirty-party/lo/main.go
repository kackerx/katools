package main

import (
    "fmt"
    "github.com/samber/lo"
    "math/rand"
    "time"
)

func main() {

    rand.Seed(time.Now().UnixNano())

    for i := 0; i < 10; i++ {
        res := lo.Sample([]int{1, 2, 3})
        fmt.Println(res)
    }

    lo.ForEach([]int{1, 2, 3, 4}, func(e int, i int) {
        fmt.Println(e, i)
        fmt.Println("----")
    })
}
