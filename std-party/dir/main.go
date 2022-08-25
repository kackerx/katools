package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func foo() error {
	return errors.New("kacker")
}
