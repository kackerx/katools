package main

import (
    "errors"
    "fmt"
)


type CycleQueue struct {
    data        [MAX_SIXE]int
    front, rear int
}

func (q *CycleQueue) isFull() bool {
    return (q.rear + 1) % MAX_SIXE == q.front
}

func (q *CycleQueue) isEmpty() bool {
    return q.front == q.rear
}

func (q *CycleQueue) put(e int) error {
    if q.isFull() {
        return errors.New("满")
    }
    q.data[q.rear] = e
    q.rear = (q.rear + 1) % MAX_SIXE
    return nil
}

func (q *CycleQueue) get() (int, error) {
    if q.isEmpty() {
        return 0, errors.New("空")
    }
    
    defer func() {
        q.front = (q.front + 1) % MAX_SIXE
    }()
    
    return q.data[q.front], nil
}

func (q *CycleQueue) show() {
    if q.isEmpty() {
        fmt.Println("空")
        return
    }
    
    for i := q.front; i < q.front + q.size(); i++ {
        fmt.Println(i)
    }
}

func (q *CycleQueue) size() int {
    return (q.rear + MAX_SIXE - q.front) % MAX_SIXE
}

func (q *CycleQueue) head() (int, error) {
    if q.isEmpty() {
        return 0, errors.New("空")
    }
    return q.data[q.front], nil
}

