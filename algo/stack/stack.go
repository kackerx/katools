package main

type Stack interface {
    push(e int)
    pop() int
    peek() int
    getSize() int
    isEmpty() bool
}

const MAX_SIXE = 10

type ArrayStack struct {
    array [MAX_SIXE]int
    top   int
}

func (s *ArrayStack) push(e int) {
    if s.top == MAX_SIXE-1 {
        return
    }
    s.top++
    s.array[s.top] = e
}

func (s *ArrayStack) pop() int {
    defer func() {
        s.top--
    }()
    return s.array[s.top]
}

func (s *ArrayStack) peek() int {
    return s.array[s.top]
}

func (s *ArrayStack) getSize() int {
    panic("implement me")
}

func (s *ArrayStack) isEmpty() bool {
    panic("implement me")
}



//type LinkedStack struct {
//    *linkedlist.Linkedlist
//}
//
//func NewLinkedStack() *LinkedStack {
//    return &LinkedStack{Linkedlist: linkedlist.NewLinkedlist()}
//}
//
//func (l *LinkedStack) push(e int) {
//    l.AddFirst(e)
//}
//
//func (l *LinkedStack) pop() int {
//    return l.DelFirst()
//}
//
//func (l *LinkedStack) peek() int {
//    return l.GetFirst()
//}
//
//func (l *LinkedStack) getSize() int {
//    return l.Size
//}
//
//func (l *LinkedStack) isEmpty() bool {
//    return l.IsEmpty()
//}
//
//func (l *LinkedStack) show() {
//    l.Show()
//}
