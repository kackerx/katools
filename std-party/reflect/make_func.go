package main

import (
	"fmt"
	"reflect"
	"time"
)

func MakeTimeFunc(f interface{}) interface{} {
	tf := reflect.TypeOf(f)
	vf := reflect.ValueOf(f)
	if tf.Kind() != reflect.Func {
		panic("f is not func")
	}

	retFunc := reflect.MakeFunc(tf, func(args []reflect.Value) (results []reflect.Value) {
		now := time.Now()
		results = vf.Call(args)
		fmt.Printf("func takes %v s\n", time.Now().Sub(now))
		return results
	})

	return retFunc.Interface()
}

func Takes() { time.Sleep(time.Second * 2) }

func test() string { return "kacker" }

func main() {
	t := reflect.ValueOf(test)
	if _, ok := t.Interface().(func() string); ok {
		fmt.Println(ok)
	} else {
		fmt.Println("no")
	}
}
