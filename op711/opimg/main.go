package main

import (
	"fmt"
	"github.com/kackerx/ktools"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"sync"
)

type Video struct {
	//Vid  uint   `gorm:"primaryKey" gorm:"column:v_id"`
	ID   uint   `gorm:"column:v_id"`
	Name string `gorm:"column:v_name"`
	Pic  string `gorm:"column:v_pic"`
}

var (
	db *gorm.DB
	wg sync.WaitGroup
)

func (Video) TableName() string {
	return "sea_data"
}

func main() {
	var (
		upDatas map[uint]Video
		retMap  sync.Map
		imgUrl  string
		err     error
	)
	dsn := "mysqlConnectConfig"
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println(err)
		return
	}

	var videos []Video
	ret := db.Order("v_addtime desc").Limit(100).Find(&videos)
	if ret.Error != nil {
		fmt.Println(ret.Error)
		return
	}

	upDatas = make(map[uint]Video)
	for _, item := range videos {
		if !strings.Contains(item.Pic, "mgtv") && !strings.Contains(item.Pic, "alicdn") {
			fmt.Println("addr", item)
			upDatas[item.ID] = item
			//upDatas.Store(item.Vid, &item)
		}
	}
	fmt.Println(upDatas)

	wg.Add(len(upDatas))
	//upDatas.Range(func(key, value interface{}) bool {
	//	wg.Add(1)
	//	return true
	//})
	for _, v := range upDatas {
		go func(video Video) {
			defer wg.Done()
			if imgUrl, err = ktools.UploadImg(video.Pic); err != nil {
				fmt.Println(err)
				return
			}
			video.Pic = imgUrl
			retMap.Store(video.ID, video)
		}(v)
	}
	wg.Wait()

	retMap.Range(func(key, value interface{}) bool {
		video := value.(Video)
		db.Model(&video).Update("v_pic", video.Pic)
		return true
	})

}
