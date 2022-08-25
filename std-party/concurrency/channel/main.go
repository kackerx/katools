package main

func main() {

    p := Person{Name: "li", Age: 27}

    p = p
}

func foo() *int {
    v := 0
    return &v
}

type Person struct {
    Name string
    Age  int
}

var p = Person{
    Name: "kacker",
    Age:  0,
}

type Pee struct {
    Name string
}
