package main

import (
    "errors"
    "fmt"
)

type queue struct {
    array             [5]int
    front, tail, size int
}

func (q *queue) enqueue(e int) error {
    if (q.tail + 1) % 5 == q.front {
        return errors.New("q is full")
    }
    
    q.array[q.tail] = e
    q.tail = (q.tail + 1) % 5
    q.size++
    return nil
}

func (q *queue) cap() int {
    return (q.tail + 5 - q.front) % 5
}

func (q *queue) show() {
    tempHead := q.front
    for i := 0; i < q.cap(); i++ {
        fmt.Println(q.array[tempHead])
        tempHead = (tempHead + 1) % 5
    }
}

func main() {
    q := queue{
        front: 0,
        tail:  0,
        size:  0,
    }
    
    q.enqueue(1)
    q.enqueue(2)
    q.enqueue(3)
    q.enqueue(4)
    q.enqueue(5)
    
    q.show()
}

