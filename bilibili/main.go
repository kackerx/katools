package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Resp struct {
	Code int `json:"code"`
	Data struct {
		Dash struct {
			Video []struct {
				Id      int    `json:"id"`
				BaseUrl string `json:"baseUrl"`
			} `json:"video"`
			Audio []struct {
				Id      int    `json:"id"`
				BaseUrl string `json:"baseUrl"`
			} `json:"audio"`
		} `json:"dash"`
	} `json:"data"`
}

func GetHtml(url string) string {
	request, _ := http.NewRequest("GET", url, nil)
	cli := http.DefaultClient
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 Edg/89.0.774.57")
	request.Header.Add("referer", "https://api.bilibili.com/")
	resp, err := cli.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	resp_str, err := ioutil.ReadAll(resp.Body)
	return string(resp_str)
}

func GetJson(url string) (Resp, string, string) {
	var p string
	if strings.Contains(url, "p=") {
		p = strings.Split(url, "=")[1]
	} else {
		p = "0"
	}
	url_str := GetHtml(url)
	re := regexp.MustCompile(`"bvid":"(.*?)"`)
	bvid := re.FindStringSubmatch(url_str)[1]
	
	cre := regexp.MustCompile(`"cids":{"\d+":(.*?)}`)
	cid := cre.FindStringSubmatch(url_str)[1]
	
	api_url := fmt.Sprintf("https://api.bilibili.com/x/player/playurl?cid=%s&bvid=%s&qn=80&type=&otype=json&fourk=1&fnver=0&fnval=80&session=72dd92313cd27f9d461784d946fc28d0", cid, bvid)
	playUrl := GetHtml(api_url)
	
	var resp Resp
	err := json.Unmarshal([]byte(playUrl), &resp)
	if err != nil {
		panic(err)
	}
	return resp, bvid, p
}

func DownVideo(endCh chan interface{}, resp Resp, bvid string, p string) {
	fmt.Println("开始下载")
	defer wg.Done()
	go func() {
		for {
			select {
			case <-endCh:
				goto success
			default:
				for _, r := range "-\\|/" {
					fmt.Printf("\r%c", r)
					time.Sleep(50 * time.Millisecond)
				}
			}
			
		}
	success:
		fmt.Printf("\r%s\n", "下载完成")
		color.Cyan("❯❯❯ 输入bilbili播放页地址后回车, 输入0退出: ")
	}()
	
	videoUrl := resp.Data.Dash.Video[2].BaseUrl
	audioUrl := resp.Data.Dash.Audio[2].BaseUrl
	
	wg.Add(2)
	go DownToFile(videoUrl, bvid, p, "video")
	go DownToFile(audioUrl, bvid, p, "audio")
}

func DownToFile(url, bvid, p, tp string) {
	defer wg.Done()
	res := GetHtml(url)
	// Create the file
	//fmt.Println(cid)
	out, err := os.Create(fmt.Sprintf("%s_%s_%s.mp4", bvid, p, tp))
	if err != nil {
		panic(err)
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, strings.NewReader(res))
}

func main2() {
	
	endCh := make(chan interface{})
	
	color.Cyan("❯❯❯ 输入bilbili播放页地址后回车, 输入0退出: ")
	for {
		defer func() {
			pa := recover()
			fmt.Println(pa)
		}()
		var url string
		_, err := fmt.Scanf("%s", &url)
		if err != nil {
			panic(err)
		}
		
		if url == "0" {
			break
		}
		wg.Add(1)
		resp, bvid, p := GetJson(url)
		go DownVideo(endCh, resp, bvid, p)
		wg.Wait()
		endCh <- nil
	}
}



