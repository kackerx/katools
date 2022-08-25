package main

import (
    "fmt"
)

type Option func(*User)

type User struct {
    ID   uint `gorm:""`
    Name string
    Age  int
}

func (u *User) WithAge(age int) *User {
    u.Age = age
    return u
}

func NewUser(ID uint, name string, opts ...Option) *User {
    u := &User{ID: ID, Name: name}
    for _, v := range opts {
        v(u)
    }
    return u
}

func WithAge(age int) Option { return func(user *User) { user.Age = age } }

func main() {
    u := GetUser()

    fmt.Println(u)

}

func GetUser() *User {
    opts := []Option{
        WithAge(1),
        WithAge(1),
    }
    u := NewUser(1, "kacker", opts...)
    u.WithAge(2)
    return u
}
