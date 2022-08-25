package main

import (
	"fmt"
	"reflect"
)

type book struct {
	name string
	age  int
}

type handler func(int, int) int

func meta(obj interface{}) {
	t := reflect.TypeOf(obj) // int, string, main.user, func(int, int) int, main.handler
	n := t.Name()            // int, string, user, "", handler
	k := t.Kind()            // int, string, struct, func, func
	v := reflect.ValueOf(obj)

	fmt.Printf("type: %s name: %s kind: %s value: %v\n", t, n, k, v)
}

//func main() {
//	var intVar int = 10
//	var strVar string = "hello"
//	var structVar user = user{
//		name: "kacker",
//		age:  27,
//	}
//	var sum = func(a, b int) int {
//		return a + b
//	}
//	var sub handler = func(a, b int) int {
//		return a - b
//	}
//
//	meta(intVar)
//	meta(strVar)
//	meta(structVar)
//	meta(sum)
//	meta(sub)
//}
