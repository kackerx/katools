package main

import (
	"fmt"
	"strings"
)

func main() {
	a := strings.Builder{}
	a.WriteRune('哦')

	fmt.Println(a.String())
}
