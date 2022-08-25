package main

import "fmt"

func main() {
	s := "我是china"
	for i, v := range []rune(s) {
		fmt.Println(i)
		fmt.Println(v)
	}
}
