package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func GetHtml(uri string) (string, error) {
	var (
		resp *http.Response
		htmlByte []byte
		err error
	)
	if resp, err = http.Get(uri); err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	if htmlByte, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(htmlByte), nil
}

func main() {
	var (
		uri string
		err error
		html string
	)
	if uri, err = clipboard.ReadAll(); err != nil {
		fmt.Println(err)
		return
	}
	if html, err = GetHtml(uri); err != nil {
		return
	}

	reg := regexp.MustCompile(`.*?a:"(.*?)"`)
	ret := reg.FindStringSubmatch(html)
	m3u8 := ret[1]
	err = clipboard.WriteAll(strings.Replace(m3u8, "http:", "https:", -1))
	if err == nil {
		fmt.Println(m3u8)
	}
	
}
