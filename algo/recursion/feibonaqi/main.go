package main

import "fmt"

// Foo 1, 1, 2, 3, 5, 8, 13
func Foo(n int) int {
    if n <= 1 {
        return 1
    }
    
    return Foo(n-1) + Foo(n-2)
}

func Bar(n int) {
    a, b := 1, 1
    for b < n {
        fmt.Println(a, b)
        a, b = b, a+b
    }
}

func dynamicFib(n int) int {
    memo := make([]int, n+1)
    memo[0] = 0
    memo[1] = 1
    for i := 2; i <= n; i++ {
        memo[i] = memo[i-1] + memo[i-2]
    }
    
    return memo[n]
}


func main() {
    fmt.Println(dynamicFib(8))
}

