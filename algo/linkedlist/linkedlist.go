package main

import (
    "fmt"
    "strconv"
)

type Node struct {
    data int
    next *Node
}

type Linkedlist struct {
    Size int
    Head *Node
}

func NewLinkedlist() *Linkedlist { return &Linkedlist{Head: &Node{}} }

func (l *Linkedlist) GetSize() int { return l.Size }

func (l *Linkedlist) IsEmpty() bool { return l.Size == 0 }

func (l *Linkedlist) AddFirst(e int) {
    l.Insert(e, 0)
    //l.Head.next = &Node{data: e, next: l.Head.next} // 新节点的next指向head的next
    //l.Size++
}

func (l *Linkedlist) AddLast(e int) {
    l.Insert(e, l.Size)
}

func (l *Linkedlist) Insert(e, index int) {
    pre := l.Head // 定位前一个节点, 插入前一个和后一个节点中间的位置
    for i := 0; i < index; i++ {
        pre = pre.next
    }
    
    pre.next = &Node{
        data: e,
        next: pre.next, // 新节点的next指向pre的next: 插入到中间
    }
    
    l.Size++
}

func (l *Linkedlist) Get(index int) int {
    cur := l.Head.next // 定位
    for i := 0; i < index; i++ {
        cur = cur.next
    }
    return cur.data
}

func (l *Linkedlist) GetFirst() int {
    return l.Get(0)
}

func (l *Linkedlist) GetLast() int {
    return l.Get(l.Size - 1)
}

func (l *Linkedlist) Show() {
    var res string
    cur := l.Head.next
    for cur != nil {
        res += strconv.Itoa(cur.data) + "->"
        cur = cur.next
    }
    fmt.Println(res)
}

func (l *Linkedlist) Contains(e int) bool {
    cur := l.Head.next
    for cur != nil {
        if cur.data == e {
            return true
        }
        cur = cur.next
    }
    return false
}

func (l *Linkedlist) Del(index int) int {
    pre := l.Head // 找到前一个节点, 连接要删除的节点的下一个节点
    for i := 0; i < index; i++ {
        pre = pre.next
    }
    
    retNode := pre.next
    pre.next = retNode.next
    l.Size--
    return retNode.data
}

func (l *Linkedlist) DelFirst() int {
    return l.Del(0)
}

func (l *Linkedlist) Rev() {
    cur := l.Head.next
    newHead := Node{}
    
    for cur != nil {
        next := cur.next
        cur.next = newHead.next
        newHead.next = cur
        cur = next
    }
    l.Head.next = newHead.next
}

func (l *Linkedlist) Rev2() {
    cur := l.Head.next
    var pre *Node
    for cur != nil {
        next := cur.next
        cur.next = pre
        pre = cur
        cur = next
    }
    l.Head.next = pre
}

func (l *Linkedlist) RecursiveRec() {
    l.Head.next = l.RecursiveNode(l.Head.next)
}

func (l *Linkedlist) RecursiveNode(head *Node) *Node {
    if head == nil || head.next == nil {
        return head
    }
    
    rev := l.RecursiveNode(head.next)
    head.next.next = head
    head.next = nil
    return rev
}
