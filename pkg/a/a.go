package aa

import "fmt"

func AFunc() {
	u := User{Name: "kacker", Age: 0}
	k := NewUser("k", 27)

	if u.Name == k.Name {
		fmt.Println("same name")
	}
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) Error() string {
	return u.Name
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}
