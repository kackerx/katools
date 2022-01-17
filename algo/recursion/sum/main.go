package main

import "fmt"

func main() {
    head := &Node{
        data: 0,
        next: &Node{
            data: 1,
            next: &Node{
                data: 2,
                next: &Node{
                    data: 3,
                    next: nil,
                },
            },
        },
    }
    
    ret := removeNode(head, 1, 0)
    for ret != nil {
        fmt.Println(ret.data)
        ret = ret.next
    }
}

func sum(arr []int) int {
    if len(arr) == 1 {
        return arr[0]
    }
    
    return arr[0] + sum(arr[1:])
}

