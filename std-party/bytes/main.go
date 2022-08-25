package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	buf.WriteString("hello")

	fmt.Println(buf.Len(), buf.Cap()) // 5, 64

	var p = make([]byte, 3)
	n, _ := buf.Read(p)
	fmt.Println(n, p, buf.Len()) // 3 [*, *, *] 2, Len()是未被读取的字节长度

	for k, v := range map[string]int{"k": 1} {
		fmt.Println(k, v)
	}
}
