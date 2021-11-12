package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-redis/redis/v8"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestPlay(t *testing.T) {
	play := "HD高清$test1m3u8$$$HD$test.mp4"
	playRet := make([]string, 1)
	if strings.Contains(play, "$$$") {
		playRet = strings.Split(play, "$$$")
	} else {
		playRet[0] = play
	}

	for i, v := range playRet {
		if strings.Contains(v, "m3u8") {
			playRet[i] = "38云播$$" + v + "$md"
		} else if strings.Contains(v, "mp4") {
			playRet[i] = "普通视频$$" + v + "$zhilian"
		} else if strings.Contains(v, "$D") {
			playRet[i] = "南瓜云播$$" + v + "$ng"
		} else {
			playRet[i] = "芒乐秒播$$" + v + "$ml"
		}
	}

	fmt.Println(playRet)
	plays := strings.Replace(strings.Trim(fmt.Sprint(playRet), "[]"), " ", "$$$", -1)
	fmt.Println(plays)
}

func TestFanqie(t *testing.T) {
	url := "https://moblie-api.fqdy.pro/cms/app/promotion/channel/home"

	body := map[string]interface{}{"order": "gmt_just_in", "area": "", "tag": "", "year": "", "summaryTagIds": "", "size": 30, "channelType": "MPC"}
	bodyData, _ := json.Marshal(body)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyData))
	if err != nil {
		panic(err)
	}

	cli := http.DefaultClient
	request.Header["content-type"] = []string{"application/json;charset=UTF-8"}

	resp, err := cli.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	jsonStr := string(ret)
	code := gjson.Get(jsonStr, "data.searchResults.#.movieId").Array()
	title := gjson.Get(jsonStr, "data.searchResults.#.title").Array()
	fmt.Println(jsonStr)
	for i, v := range code {
		fanqieVideos = append(fanqieVideos, video{id: v.String(), title: title[i].String()})
	}
	//for _, item := range fanqieVideos {
	//fmt.Printf("  %s: %s: %s\n", item.id, item.title)
	//}
	//fmt.Println(code)
	//
	//ret, err = Get(fmt.Sprintf("https://moblie-api.fqdy.pro/cms/app/promotion/channel/get_movie_info?movieId=%s&channelType=MPC", url[1:]))
	//
	//var d FqData
	//json.Unmarshal(ret, &d)
	//data := d.Data
	//Img, _ := UploadImgTx(data.CoverVerticalUrl)
	//HImg, _ := UploadImgTx(data.CoverHorizontalUrl)
}

func TestHehePlay(t *testing.T) {
	ret := "%u9ad8%u6e05%u4e2d%u5b57%24%09Dzbec4d07f502f6749485e7910b"

	ret = strings.ReplaceAll(ret, "%09", ``)
	ret = strings.ReplaceAll(ret, "%u", `\u`)
	ret, err = url2.PathUnescape(ret)
	if err != nil {
		panic(err)
	}

	retByte, err := UnescapeUnicode([]byte(ret))
	if err != nil {
		panic(err)
	}

	play := string(retByte)
	fmt.Println(play)
	play = strings.ReplaceAll(play, "#", "$$$")
	play = getPlay(play)
	fmt.Println(play)
}

func TestFanqiePlay(t *testing.T) {
	getFanqieDetail("13697")
}

func TestRedis(t *testing.T) {
	RdsCli := redis.NewClient(&redis.Options{
		Addr:         "10.0.3.100:6379",
		Password:     "EfcHGSzKqg6cfzWq",
		DB:           0,
		PoolSize:     70,
		MinIdleConns: 50,
	})
	f, err := os.Open("u.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	buf := make([]byte, 0)
	for err != io.EOF {
		buf, _, err = reader.ReadLine()
		id := string(buf)
		if err == io.EOF {
			break
		}
		r, _ := RdsCli.SIsMember(context.TODO(), "douyin:live:df:userid", id).Result()
		if !r {
			id = strings.Trim(id, " ")
			resp, _ := Get("https://webcast-hl.amemv.com/webcast/room/reflow/info/?room_id=7012166810174733548&type_id=0&user_id=" + id + "&live_id=1&app_id=1128")
			ret := gjson.Get(string(resp), "data.user.id").String()
			res, _ := strconv.Atoi(ret)
			if res != 0 {
				fmt.Println(res)
				RdsCli.SAdd(context.TODO(), "douyin:live:df:userid", res)
			}
		}
	}
}

func TestWujn(t *testing.T) {
	url := "http://wujinzy.com/"
	resp, err := Get(url)
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		panic(err)
	}

	doc.Find(".xing_vb4 a").Each(func(i int, selection *goquery.Selection) {
		content, _ := selection.Attr("href")
		go func(url string) {
			resp, _ := Get(url)
			doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
			if err != nil {
				panic(err)
			}

			category := doc.Find(".nvc dd").Text()
			var play string
			doc.Find(".vodplayinfo ul:nth-of-type(2)").Find("li").Each(func(i int, selection *goquery.Selection) {
				play += selection.Text() + "$wjm3u8#"
			})
			play = "无尽云播$$" + play
			play = play[0 : len(play)-1]

			doc.Find(".vod").Each(func(i int, selection *goquery.Selection) {
				img, _ := selection.Find("img").Attr("src")
				title := selection.Find(".vodh h2").Text()
				director := selection.Find(".vodinfobox ul li:nth-child(2) span").Text()
				actor := selection.Find(".vodinfobox ul li:nth-child(3) span").Text()
				area := selection.Find(".vodinfobox ul li:nth-child(5) span").Text()
				year := selection.Find(".vodinfobox ul li:nth-child(7) span").Text()
				note := selection.Find(".vodh span").Text()
				desc := selection.Find(".vodplayinfo").Text()
				category = strings.Split(category, "/")[0]
				tid := GetTid(strings.Replace(category, "片", "", 0))

				wujinVideos = append(wujinVideos, video{
					tid:      tid,
					img:      img,
					title:    title,
					play:     play,
					actor:    actor[0:20],
					director: director,
					area:     area,
					year:     year,

					note:    note,
					content: desc,
				})
			})
		}("http://wujinzy.com" + content)
	})
	time.Sleep(time.Second * 3)
	for _, v := range wujinVideos {
		fmt.Println(v.title, v.year)
	}
}

func TestFun(t *testing.T) {
	that, _ := time.Parse("2006-01-02","2019-06-12")
	fmt.Println(time.Now().Sub(that))
	
	//getWujin("http://wujinzy.com/")
	//play := "南瓜云播$$高清中字$110$ng$$$38云播$$高清中字$muu8$md"
	//raw := "麻花云播$$HD$https://cdn4.mh-qiyi.com/20210415/2403_a6e28120/index.m3u8$mahua$$$南瓜云播$$高清中字$Dzbe6d7793b8fe3ad9661d11059$ng"
	//ret := isPlayEqual(play, raw)
	//fmt.Println(ret)
}

func TestTcp(t *testing.T) {
	s := make([]int, 3)
	fmt.Println(cap(s))
}

func TestTip(t *testing.T) {
	api := "https://fapi.binance.com"
	//resp, err := Get(api + "/fapi/v1/premiumIndex?symbol=BTCUSDT")
	resp, err := Get(api + "/fapi/v1/allOrders?symbol=BTCUSDT")
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(resp))
}



