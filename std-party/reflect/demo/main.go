package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    u := User{
        Name: "kacker",
        Age:  27,
    }
    typ := reflect.TypeOf(u)
    val := reflect.ValueOf(u)

    fmt.Println(typ.Name(), typ.String(), typ.Kind(), typ.Field(0).Tag)

    fmt.Println(val.Field(0))

}
