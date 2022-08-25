package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func Test(a string) int {
	return len(a)
}

func main() {

	inspect(Test)
}

func inspect(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	switch t.Kind() {
	case reflect.Struct:
		fmt.Printf("number of field: %s\n", t.NumField())
		for i := 0; i < t.NumField(); i++ {
			fmt.Printf("field name: %s, type: %s, value: %s\n", t.Field(i).Name, t.Field(i).Type, v.Field(i))
		}
	case reflect.Func:
		fmt.Printf("in: %s, out: %s\n", t.NumIn(), t.NumOut())
	}
}

func monitor() {

}
