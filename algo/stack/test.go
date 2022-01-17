package main

import "fmt"

func main() {
    s := ArrayStack{
        array: [10]int{},
        top:   -1,
    }
    s.push(0)
    s.push(1)
    s.push(2)
    s.push(3)
    
    fmt.Println(s.peek())
    fmt.Println(s.pop())
    fmt.Println(s.peek())
}
