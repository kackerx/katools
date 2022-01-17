package main

import "fmt"

func main() {
    q := LinkedQueue{}
    q.enqueue(0)
    q.enqueue(1)
    q.enqueue(2)
    q.enqueue(3)
    fmt.Println(q.getSize())
    q.show()
}
