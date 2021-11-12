package main

import "fmt"

func main() {
    s := []int{0, 1, 2, 3}
    
    s = append(s[:2], s[3:]...)
    
    fmt.Println(s)
}
