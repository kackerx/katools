package main

import (
	"fmt"
)

type user struct {
    Name string `json:"name"`
}

func (u user) GetName(a string) string {
    fmt.Println(a)
    return u.Name
}

//func main() {
//	b := user{Name: "kacker"}
//	t := reflect.TypeOf(b)
//	v := reflect.ValueOf(b)
//	field, _ := t.FieldByName("Name")
//	fmt.Println(field.Tag.Get("json")) // 获取tag
//
//	method := v.MethodByName("GetName") // 调用方法
//	vSlice := method.Call([]reflect.Value{reflect.ValueOf("hehe")})
//	fmt.Println(vSlice[0].Interface().(string))
//}
