package main

import (
    "fmt"
)

type Person interface {
    Eat()
}

type Man struct {
    Name string
}

// 结构体实现方法, 或默认生成一个结构体指针实现的方法
func (m Man) Eat() {
    fmt.Println("man eat food")
}

func main() {

    var m any = &Man{Name: "kacker"}

    mm := m.(*Man)

    fmt.Println(mm.Name)
}
