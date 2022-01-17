package main

import "fmt"

type Emp struct {
    id   int
    name string
    next *Emp
}

type empLink struct {
    head *Emp
}

func (e *empLink) insert(emp *Emp) {
    cur := e.head
    if cur == nil {
        e.head = emp
        return
    }
    
    var pre *Emp
    
    
    for {
        
        if cur.id > emp.id {
            break
        } else {
            cur = cur.next
        }
        
    }
    for cur != nil && cur.id < emp.id {
        pre = cur
        cur = cur.next
    }
    
    if cur == nil {
        pre.next = emp
        return
    }
    
    emp.next = pre.next
    pre.next = emp
}

func (e *empLink) show(no int) {
    if e.head == nil {
        fmt.Printf("链表%d为空\n", no)
        return
    }
    
    cur := e.head
    for cur != nil {
        fmt.Printf("链表%d: %d, %s\n", no, cur.id, cur.name)
        cur = cur.next
    }
    

    
}
