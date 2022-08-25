package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int = 1
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	fmt.Println(v.Interface().(int))

	field, _ := t.FieldByName("name")
	field.Tag.Get("name")
}
