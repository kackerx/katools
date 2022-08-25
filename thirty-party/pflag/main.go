package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

var ()

func main() {
	// P表示支持短选项, Var表示存到变量中而不是指针
	var flagVar = pflag.StringP("name", "n", "kackerr", "help usage for name")

	var age int
	pflag.IntVarP(&age, "age", "a", 27, "you age")

	pflag.Parse()

	fmt.Println(*flagVar)

	fmt.Println(age)
}
