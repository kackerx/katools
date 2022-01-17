package main

import (
    "database/sql"
    "fmt"
    "github.com/pkg/errors"
)

type P struct {
    Name string `json:"name"`
    Age  int32  `json:"age"`
}

func foo() error {
    return errors.Wrap(sql.ErrNoRows, "foo error")
}

func bar() error {
    return errors.WithMessage(foo(), "bar error")
}

func main() {
    err := bar()
    if errors.Cause(err) == sql.ErrNoRows {
        fmt.Printf("%v\n", err)
        fmt.Println("-------------")
        fmt.Printf("%+v\n", err) // %+v是打印调用栈信息
    }
}

