package main

func main() {

}

type User struct {
    Name string
}

func foo() (User, error) {
    return User{}, nil
}
