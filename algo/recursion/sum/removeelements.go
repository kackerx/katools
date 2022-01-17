package main

type Node struct {
    data int
    next *Node
}

func removeNode(head *Node, val int, depth int) *Node {
    
    if head == nil {
        return head
    }
    
    head.next = removeNode(head.next, val, depth)
    
    if head.data == val {
        return head.next
    } else {
        return head
    }
}

func recursionLinked(head *Node) *Node {
    if head == nil {
        return nil
    }
    
    return nil
}
