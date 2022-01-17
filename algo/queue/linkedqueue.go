package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedQueue struct {
    head *Node
    tail *Node
    size int
}

func (l *LinkedQueue) getSize() int { return l.size }

func (l *LinkedQueue) isEmpty() bool { return l.size == 0 }

func (l *LinkedQueue) enqueue(e int) {
    if l.tail == nil { // 空队列
        l.tail = &Node{data: e}
        l.head = l.tail
        l.size++
        return
    }
    l.tail.next = &Node{data: e}
    l.tail = l.tail.next
    l.size++
}

func (l *LinkedQueue) dequeue() int {
    defer func() {
        l.head = l.head.next
        if l.head == nil { // 如果只有一个元素, 维护尾指针
            l.tail = nil
        }
        l.size--
    }()
    return l.head.data
}

func (l *LinkedQueue) getFront() int { return l.head.data }

func (l *LinkedQueue) show() {
    cur := l.head
    for cur != nil {
        fmt.Println(cur.data)
        cur = cur.next
    }
}
