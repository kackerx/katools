package main

import (
	"errors"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/atotto/clipboard"
	"github.com/kackerx/katools"
	"github.com/kackerx/katools/requests"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	err    error
	ret    string
	name   string
	rawUrl string
	db     *gorm.DB
)

type Target struct {
	Url  string
	Name string
	Img  string
}

func init() {
	//var video *Video
	db = initDb()
	rawUrl, _ = clipboard.ReadAll()
}

func initDb() *gorm.DB {
	var db *gorm.DB
	dsh := "kingvstr_dy:M3KHa7fjRkRjkkmJ@tcp(27.102.66.67:3306)/kingvstr_dy?charset=utf8"
	if db, err = gorm.Open(mysql.Open(dsh), &gorm.Config{}); err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func GetData(rawUrl string) (v *Video, p *Play, err error) {
	target, err := url.Parse(rawUrl)
	if ret, err = requests.Get(rawUrl); err != nil {
		fmt.Println(err)
		return
	}

	// mp4
	reg := regexp.MustCompile(`.*?"(http.*?mp4)"`)
	ret := reg.FindStringSubmatch(ret)
	if ret == nil {
		return nil, nil, errors.New("不是mp4")
	}
	videoUrl := strings.Replace(ret[1], "http://", "https://", -1)

	// name
	reg = regexp.MustCompile(`.*?"description" content="(.*?)"`)

	// imgUrl
	imgUrl := url.URL{
		Scheme:   "https",
		Host:     target.Host,
		RawQuery: "s=" + name,
	}
	var rets string
	if rets, err = requests.Get(imgUrl.String()); err != nil {
		fmt.Println(err)
		return
	}
	reg = regexp.MustCompile(`.*?=(https:.*?)&`)
	img := reg.FindStringSubmatch(rets)[1]
	mgImg, _ := katools.UploadImgTx(img)
	//var mgImgMap map[string]interface{}
	//err = json.Unmarshal([]byte(mgImg), &mgImgMap)
	//if err != nil {
	//	panic(err)
	//}

	//mgImg, _ = mgImgMap["data"].(map[string]string)["url"]

	return &Video{ // 34, 45, 46, 47, 59
			Tid:          46,
			Name:         name,
			Pic:          mgImg,
			Actor:        "41ts抢先电影网",
			VAddtime:     time.Now().Unix(),
			VDaytime:     time.Now().Unix(),
			VWeektime:    time.Now().Unix(),
			VMonthtime:   time.Now().Unix(),
			VPublishyear: 0,
			VNote:        "秒拍短视频",
			VDirector:    "41tss.com",
			Vip:          0,
		}, &Play{
			Tid:  46,
			Body: fmt.Sprintf("普通视频$$播放$%s$zhilian", videoUrl),
		}, nil
}

var wg sync.WaitGroup

func main() {

	//收集视频列表
	doc, err := htmlquery.LoadURL(rawUrl)
	if err != nil {
		panic(err)
	}

	nodes, err := htmlquery.QueryAll(doc, `//ul[@id="post_container"]/li`)
	if err != nil {
		panic(err)
	}

	for _, v := range nodes {
		href, err := htmlquery.Query(v, `//a`)
		if err != nil {
			panic(err)
		}

		var t Target
		t.Url = htmlquery.SelectAttr(href, "href")
		t.Name = htmlquery.SelectAttr(href, "title")
		img, err := htmlquery.Query(href, "//img")
		if err != nil {
			panic(err)
		}

		t.Img = htmlquery.SelectAttr(img, "src")
		wg.Add(1)
		go handler(t)
	}

	wg.Wait()
	os.Exit(0)

	//video, play, err := GetData(rawUrl)
	//if err != nil {
	//	panic(err)
	//}
	//
	//db.Create(video)
	//play.ID = video.ID
	//db.Create(play)
	//fmt.Println(video.Name, "成功入库!")
	////db.Order("v_addtime desc").First(&video)

}

func handler(t Target) {
	defer wg.Done()
	if ret, err = requests.Get(t.Url); err != nil {
		fmt.Println(err)
		return
	}

	// mp4
	reg := regexp.MustCompile(`.*?"(http.*?mp4)"`)
	ret := reg.FindStringSubmatch(ret)
	if ret == nil {
		return
	}

	videoUrl := strings.Replace(ret[1], "http://", "https://", -1)

	mgImg, _ := katools.UploadImgTx(t.Img)
	video := &Video{ // 34, 45, 46, 47, 59
		Tid:          59,
		Name:         t.Name,
		Pic:          mgImg,
		Actor:        "41ts抢先电影网",
		VAddtime:     time.Now().Unix(),
		VDaytime:     time.Now().Unix(),
		VWeektime:    time.Now().Unix(),
		VMonthtime:   time.Now().Unix(),
		VPublishyear: 0,
		VNote:        "秒拍短视频",
		VDirector:    "41tss.com",
		Vip:          0,
	}
	play := &Play{
		Tid:  59,
		Body: fmt.Sprintf("普通视频$$播放$%s$zhilian", videoUrl),
	}
	db.Create(video)
	play.ID = video.ID
	db.Create(play)
	fmt.Println(video.Name, "成功入库!")
}
