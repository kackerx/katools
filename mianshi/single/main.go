package main

import (
	"fmt"
	"sync"
)

type Student struct {
	Name string
	Age  int
}

var once sync.Once
var s Student

func New() *Student {
	once.Do(func() {
		s = Student{
			Name: "kacker",
			Age:  10,
		}
	})
	return &s
}

func main() {
	s := New()
	a := New()

	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", a)
}
