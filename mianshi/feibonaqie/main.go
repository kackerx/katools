package main

import "fmt"

func main() {
	ch := make(chan int)
	var res int

	go func() {
		fmt.Println("1")
		var r int
		for i := 1; i <= 10; i++ {
			r += i
		}
		ch <- r
	}()

	go func() {
		fmt.Println("2")
		var r int
		for i := 1; i <= 10; i++ {
			r *= i
		}
		ch <- r
	}()

	go func() {
		fmt.Println(3)
		ch <- 11
	}()

	for i := 0; i < 3; i++ {
		res += <-ch
	}

	close(ch)
	fmt.Println(res)
}
