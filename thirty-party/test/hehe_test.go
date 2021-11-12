package test

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestDoc(t *testing.T) {
	resp, err := Get("https://www.hehedy.com/content/893.html")
	if err != nil {
		panic(err)
	}

	var (
		img      string
		title    string
		actor    string
		director string
		category string
		area     string
		year     string
		play string
	)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	img, _ = doc.Find(".stui-content__thumb>a").Attr("data-original")
	title = doc.Find(".stui-content__detail>h3").Text()

	doc.Find(".stui-content__detail>p:nth-of-type(2)").Each(func(i int, selection *goquery.Selection) {
		actor = strings.Replace(selection.Text(), "主演：", "", -1)
	})
	doc.Find(".stui-content__detail>p:nth-of-type(3)").Each(func(i int, selection *goquery.Selection) {
		director = strings.Replace(selection.Text(), "导演：", "", -1)
	})
	doc.Find(".stui-content__detail>p:nth-of-type(4)>a").Each(func(i int, selection *goquery.Selection) {
		category = strings.Replace(selection.Text(), "电影", "", -1)
	})
	doc.Find(".stui-content__detail>p:nth-of-type(4)").Each(func(i int, selection *goquery.Selection) {
		s := selection.Text()
		reg := regexp.MustCompile(`.*?地区：(.*?)年份：(.*?)$`)
		resp := reg.FindStringSubmatch(s)
		if len(resp) > 2 {
			area = resp[1]
			year = resp[2]
		} else {
			area = ""
		}
	})
	//play_address
	resp, err = Get("https://www.hehedy.com/play/11078-1-1.html")
	if err != nil {
		panic(err)
	}
	reg := regexp.MustCompile(`mac_url=unescape\('(.*?)'`)
	match := reg.FindStringSubmatch(string(resp))
	ret := match[1]
	ret = strings.ReplaceAll(ret, "%u", `\u`)
	ret, err = url2.PathUnescape(ret)
	if err != nil {
		panic(err)
	}

	retByte, err := UnescapeUnicode([]byte(ret))
	if err != nil {
		panic(err)
	}

	play = string(retByte)
	
	fmt.Println(play)
	fmt.Println(img)
	fmt.Println(title)
	fmt.Println(actor)
	fmt.Println(director)
	fmt.Println(category)
	fmt.Println(area)
	fmt.Println(year)
}

func Get(url string) ([]byte, error) {
	client := http.DefaultClient
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header["user-agent"] = []string{"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1 Edg/94.0.4606.71"}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}