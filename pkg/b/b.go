package main

import "fmt"

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := a[1:4]

	fmt.Println(s1)

	s2 := s1[2:4]

	fmt.Println(s2)

}
