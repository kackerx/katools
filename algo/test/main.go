package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	var err error

	if false {
		err = errors.New("true")
		goto Finish
	} else {
		err = errors.New("false")
		goto Finish
	}

	for {
		select {
		case <-time.NewTicker(time.Second).C:
			fmt.Println("kacker")

		}
	}

Finish:
	fmt.Println(err)
}
