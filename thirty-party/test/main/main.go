package main

import (
    "fmt"
    "runtime"
)

func main() {
    
    runtime.GOMAXPROCS(8)
    
    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Println(i)
        }(i)
    }
    
    
}
