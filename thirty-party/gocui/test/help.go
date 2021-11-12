package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jroimartin/gocui"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	url2 "net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func getHehedy() error {
	resp, err := Get("https://www.hehedy.com/type/1-1.html")
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return err
	}

	var ids []string
	doc.Find(".stui-vodlist__box").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find("a").Attr("href")
		req := regexp.MustCompile(`.*?(\d+).*?`)
		id := req.FindStringSubmatch(href)[1]
		ids = append(ids, id)
	})

	for _, v := range ids {
		go getHehedyDetail(v)
	}

	return nil
}

func getMiaopai() {
	urls := []string{
		"https://www.mo190.com/wpfuli",
		"https://www.mo190.com/wpfuli",
		"https://www.mo190.com/wpguangchang",
		"https://www.mo190.com/xiee",
		"https://www.mo190.com/dapian",
		"https://www.mo190.com/gaoxiao",
		"https://www.mo190.com/niubi",
		//"https://www.mo190.com/wpguangchang/page/2",
		"https://www.mo190.com/qita",
	}
	for _, v := range urls {
		resp, err := Get(v)
		if err != nil {
			panic(err)
		}

		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
		if err != nil {
			panic(err)
		}

		var title string
		var img string
		doc.Find("#post_container>li").Each(func(i int, selection *goquery.Selection) {
			img, _ = selection.Find("img").Attr("src")
			title, _ = selection.Find("img").Attr("alt")
			year := selection.Find(".info>span:first-child").Text()

			y := strconv.Itoa(time.Now().Year())
			timeStr := y + "-" + year
			tt, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
			if err != nil {
				panic(err)
			}

			cha := time.Now().Sub(tt).Hours()
			if cha < 48 {
				href, _ := selection.Find(".thumbnail>a").Attr("href")
				resp, err = Get(href)
				if err != nil {
					panic(err)
				}

				reg := regexp.MustCompile(`.*?(http.*?mp4).*?`)
				play := reg.FindStringSubmatch(string(resp))

				if len(play) > 1 {
					a := []int{34, 45, 46, 47, 59}
					rand.Seed(time.Now().UnixNano())
					b := rand.Intn(4)
					playRet := strings.Replace(play[1], "http://", "https://", -1)

					img, _ = UploadImgTx(img)
					miaopaiVideos = append(miaopaiVideos, video{
						img:   img,
						title: title,
						play:  "普通视频$$播放$" + playRet + "$zhilian",
						year:  y,
						note:  "秒拍福利",
						tid:   a[b],
					})
				}
			}
		})
	}
}

func getFanqie() {
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
	//title := gjson.Get(jsonStr, "data.searchResults.#.title").Array()
	for _, v := range code {
		//fanqieVideos = append(fanqieVideos, video{id: v.String(), title: title[i].String()})
		go getFanqieDetail(v.String())
	}
}

func getFanqieDetail(id string) {
	ret, err := Get(fmt.Sprintf("https://moblie-api.fqdy.pro/cms/app/promotion/channel/get_movie_info?movieId=%s&channelType=MPC", id))
	if err != nil {
		panic(err)
	}

	var d FqData
	json.Unmarshal(ret, &d)
	data := d.Data
	Img, _ := UploadImgTx(data.CoverVerticalUrl)
	HImg, _ := UploadImgTx(data.CoverHorizontalUrl)

	fanqieVideos = append(fanqieVideos, video{
		id:       id,
		tid:      GetTid(data.TagNameList[0]),
		img:      Img,
		title:    data.Title,
		play:     fmt.Sprintf("番茄云播$$HD$%d_%s$fq", data.MovieEpisodeInfoList[0].MovieEpisodeId, id),
		actor:    "",
		director: "",
		category: "",
		area:     data.AreaName,
		year:     strconv.Itoa(data.Year),
		note:     "HD高清",
		content:  data.Desc,
		spic:     HImg,
	})

}

func UploadImgTx(rawImg string) (ret string, err error) {
	targetUrl := "https://om.qq.com/image/orginalupload"
	resp, err := http.Get(rawImg)
	if err != nil {
		return "http://www.hulupa.com/zuoz/img/load.gif", nil
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	fileWriter, err := w.CreateFormFile("Filedata", rawImg)
	if err != nil {
		panic(err)
	}

	io.Copy(fileWriter, bytes.NewReader(payload))
	contentType := w.FormDataContentType()
	w.Close()

	resp, err = http.Post(targetUrl, contentType, buf)
	if err != nil {
		panic(err)
	}

	retImg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var m Img
	err = json.Unmarshal(retImg, &m)
	if err != nil {
		return "http://www.hulupa.com/zuoz/img/load.gif", nil
	}

	return m.Data.Url, nil
}

func getHehedyDetail(id string) {
	resp, err := Get(fmt.Sprintf("https://www.hehedy.com/content/%s.html", id))
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
		play     string
		content  string
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
	content = doc.Find(".detail-sketch").Text()
	//play_address
	resp, err = Get(fmt.Sprintf("https://www.hehedy.com/play/%s-1-1.html", id))
	if err != nil {
		panic(err)
	}
	reg := regexp.MustCompile(`mac_url=unescape\('(.*?)'`)
	match := reg.FindStringSubmatch(string(resp))
	ret := match[1]
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

	play = string(retByte)
	play = strings.ReplaceAll(play, "#", "$$$")
	play = getPlay(play)

	v := video{
		id:       id,
		tid:      GetTid(category),
		img:      img,
		title:    title,
		play:     play,
		actor:    actor,
		director: director,
		category: category,
		area:     area,
		year:     year,
		content:  content,
	}
	hehedyVideos = append(hehedyVideos, v)
}

func GetTid(s string) int {
	switch s {
	case "剧情":
		return 9
	case "动作":
		return 5
	case "喜剧":
		return 6
	case "爱情":
		return 7
	case "科幻":
		return 8
	case "惊悚":
		return 11
	case "恐怖":
		return 11
	case "犯罪":
		return 12
	case "动画":
		return 48
	case "奇幻":
		return 49
	case "警匪":
		return 50
	case "悬疑":
		return 51
	case "冒险":
		return 52
	case "武侠":
		return 55
	case "国产剧":
		return 13
	case "大陆剧":
		return 13
	case "香港剧":
		return 14
	case "台湾剧":
		return 14
	case "港台剧":
		return 14
	case "日本剧":
		return 15
	case "韩国剧":
		return 15
	case "日韩剧":
		return 15
	case "美国剧":
		return 16
	case "欧美剧":
		return 16
	case "大陆综艺":
		return 29
	case "日韩综艺":
		return 30
	case "港台综艺":
		return 30
	case "欧美综艺":
		return 30
	case "台湾综艺":
		return 30
	case "国产动漫":
		return 60
	case "国漫":
		return 60
	case "日韩动漫":
		return 61
	case "日漫":
		return 61
	case "欧美动漫":
		return 62
	case "美漫":
		return 62
	default:
		return 9
	}
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
