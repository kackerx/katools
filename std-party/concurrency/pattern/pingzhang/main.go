/*
   屏障模式:
   main阻塞所有子goroutine等待全部完成, 聚合结果返回
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Response struct {
	Resp string
	Err  error
}

var wg sync.WaitGroup

func main() {
	urls := []string{
		"http://httpbin.org/get?a=a",
		"http://httpbin.org/get?b=b",
	}

	in := make(chan Response, len(urls))

	for _, v := range urls {
		wg.Add(1)
		go doRequest(in, v)
	}

	go func() {
		wg.Wait()
		close(in)
	}()

	var resps []Response
	for resp := range in {
		resps = append(resps, resp)
	}

	fmt.Println(resps)
}

func doRequest(out chan<- Response, url string) {
	defer wg.Done()
	var resp Response
	res, err := http.Get(url)
	if err != nil {
		resp.Err = err
		out <- resp
		return
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		resp.Err = err
		out <- resp
		return
	}

	resp.Resp = string(resBytes)
	fmt.Println("send to out")
	out <- resp
}
