package main

import (
    "fmt"
    "unsafe"
)

type User struct {
    i int
    j int
}

func main() {
    var i int = 5
    var a = [100]int{}
    var s = []int{}
    var u User

    fmt.Println(unsafe.Sizeof(i)) // 8byte
    fmt.Println(unsafe.Sizeof(a)) // 800bytet
    fmt.Println(unsafe.Sizeof(s)) // 24byte: 切片描述符长度
    fmt.Println(unsafe.Sizeof(u)) // 16byte: 字段之和

    fmt.Println(unsafe.Alignof(i)) // 8byte
    fmt.Println(unsafe.Alignof(a)) // 8bytet
    fmt.Println(unsafe.Alignof(s)) // 8byte
    fmt.Println(unsafe.Alignof(u)) // 8byte

    Foo(&User{i: 0, j: 1})

    var p = &i
    fmt.Println(unsafe.Sizeof(p))
    var r rune = 's'
    fmt.Println(unsafe.Sizeof(r))
}

func Foo(u *User) {
    ptr := unsafe.Pointer(u)
    i := (*int)(ptr)                                                 // 第一字段地址
    j := (*int)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(u.j))) // 第二字段地址
    *j += 1                                                          // 修改第二字段值

    fmt.Println(i)
}
