package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		select {
		case <-time.After(time.Second * 4):
			fmt.Println("task end")
			fmt.Fprintf(writer, "hello world")
		case <-request.Context().Done():
			fmt.Println("任务被中断")
		}
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println(err)
	}
}
