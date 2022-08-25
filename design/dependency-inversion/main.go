package main

import (
	"fmt"
	"net/http"
	"time"
)

type Person struct {
	Name string
	Age  int
	Err  string
}

func (p *Person) Error() string { return p.Err }

func (p *Person) GetName() string { return p.Name }

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Second * 5)
		fmt.Fprintf(writer, "hello %s", "kacker")
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}
