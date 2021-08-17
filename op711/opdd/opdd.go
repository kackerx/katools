package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/kackerx/katools"
	"github.com/kackerx/katools/requests"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	err  error
	ret  string
	name string
)

func initDb() *gorm.DB {
	var db *gorm.DB
	dsh := "read:Wasd4044516520@tcp(101.33.117.86:3306)/kingvstr_dy?charset=utf8"
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
	videoUrl := strings.Replace(reg.FindStringSubmatch(ret)[1], "http://", "https://", -1)

	// name
	reg = regexp.MustCompile(`.*?"description" content="(.*?)"`)
	name = reg.FindStringSubmatch(ret)[1]

	// imgUrl
	imgUrl := url.URL{
		Scheme:   "https",
		Host:     target.Host,
		RawQuery: "s=" + name,
	}
	if ret, err = requests.Get(imgUrl.String()); err != nil {
		fmt.Println(err)
		return
	}
	reg = regexp.MustCompile(`.*?=(https:.*?)&`)
	img := reg.FindStringSubmatch(ret)[1]
	mgImg, _ := katools.UploadImg(img)

	return &Video{
			Tid:          47,
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
			Tid:  47,
			Body: fmt.Sprintf("普通视频$$播放$%s$zhilian", videoUrl),
		}, nil
}

func main() {
	//var video *Video
	db := initDb()
	rawUrl, _ := clipboard.ReadAll()
	video, play, err := GetData(rawUrl)
	if err != nil {
		panic(err)
	}
	
	db.Create(video)
	play.ID = video.ID
	db.Create(play)
	fmt.Println(video.Name, "成功入库!")
	//db.Order("v_addtime desc").First(&video)

}
