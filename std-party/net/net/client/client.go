package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8989")
	if err != nil {
		log.Fatal(err)
	}

	n, err := conn.Write([]byte("hello world kacker"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("write %s byte\n", string(n))
}
