package main

import (
    "errors"
    "fmt"
)

const MAX_SIXE = 3

type ArrayQueue struct {
    data        [MAX_SIXE]int `json:"data,omitempty"`
    front, rear int
}

func (q *ArrayQueue) isFull() bool {
    return q.rear == MAX_SIXE
}

func (q *ArrayQueue) isEmpty() bool {
    return q.front == q.rear
}

func (q *ArrayQueue) put(e int) error {
    if q.isFull() {
        return errors.New("满")
    }
    q.data[q.rear] = e
    q.rear++
    return nil
}

func (q *ArrayQueue) get() (int, error) {
    if q.isEmpty() {
        return 0, errors.New("空")
    }
    
    defer func() {
        q.front++
    }()
    
    return q.data[q.front], nil
}

func (q *ArrayQueue) show() {
    if q.isEmpty() {
        fmt.Println("空")
        return
    }
    
    for k, v := range q.data {
        fmt.Printf("array:%d=%d\n", k, v)
    }
}

func (q *ArrayQueue) head() (int, error) {
    if q.isFull() {
        return 0, errors.New("空")
    }
    return q.data[q.front], nil
}

