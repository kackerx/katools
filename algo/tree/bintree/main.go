package main

import "fmt"

type Node struct {
    data  int
    left  *Node
    right *Node
}

type BST struct {
    root *Node
    size int
}

func (t *BST) add(e int) {
    t.root = t.addNode(t.root, e)
}

// 返回插入新节点后二分搜索树的根
func (t *BST) addNode(node *Node, e int) *Node {
    if node == nil {
        t.size++
        return &Node{data: e}
    }
    
    if e < node.data {
        node.left = t.addNode(node.left, e)
    } else if e > node.data {
        node.right = t.addNode(node.right, e)
    }
    
    return node
}

func (t *BST) contains(node *Node, e int) bool {
    if node == nil {
        return false
    }
    
    if node.data == e {
        return true
    } else if node.data > e {
        return t.contains(node.left, e)
    } else {
        return t.contains(node.right, e)
    }
}

func (t *BST) mid(node *Node) {
    if node != nil {
        t.mid(node.left)
        fmt.Println(node.data)
        t.mid(node.right)
    }
}

func (t *BST) last(node *Node) {
    t.max(t.root)
    t.max(t.root)
    if node != nil {
        t.last(node.left)
        t.last(node.right)
        fmt.Println(node.data)
    }
}

func (t *BST) pre(node *Node) {
    if node != nil {
        fmt.Println(node.data)
        t.pre(node.left)
        t.pre(node.right)
    }
}

func (t *BST) max(node *Node) {
    if node.right != nil {
        t.max(node.right)
    } else {
        fmt.Println(node.data)
    }
}

func (t *BST) min(node *Node) {
    if node.left != nil {
        t.min(node.left)
    } else {
        fmt.Println(node.data)
    }
}

// 函数：删除以node为根节点的最小值, 返回删除后的该节点
func (t *BST) removeMin(node *Node) *Node {
    if node.left == nil {
        // 找到最小
        t.size--
        right := node.right
        node.right = nil
        return right
    }
    
    node.left = t.removeMin(node.left)
    return node
}

func main() {
    
    t := BST{}
    t.add(5)
    t.add(3)
    t.add(6)
    t.add(8)
    t.add(4)
    t.add(2)
    
    t.min(t.root)
    t.max(t.root)
    
    fmt.Println("-------")
    t.pre(t.root)
    
    fmt.Println("-------")
    t.removeMin(t.root)
    t.pre(t.root)
    
}
