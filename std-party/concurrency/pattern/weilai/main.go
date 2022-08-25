/*
   未来模式:
   goroutine返回一个chan, main去最后等待chan的结果就好, 不用管每个goro何时结束
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	teaCh := Tee()
	coffCh := Coffee()

	fmt.Println(<-teaCh)
	fmt.Println(<-coffCh)

	fmt.Println(time.Since(now).Seconds())
}

func Tee() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		time.Sleep(time.Second * 3)
		out <- "tea is ready"
	}()
	return out
}

func Coffee() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		time.Sleep(time.Second * 3)
		out <- "coffee is ready"
	}()
	return out
}
