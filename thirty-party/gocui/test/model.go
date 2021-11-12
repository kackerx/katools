package main

type video struct {
	id       string
	tid      int
	img      string
	title    string
	play     string
	actor    string
	director string
	category string
	area     string
	year     string
	note     string
	content  string
	spic     string
}

type Video struct {
	ID           int    `gorm:"column:v_id"`
	Tid          int    `gorm:"column:tid"`
	Name         string `gorm:"column:v_name"`
	Pic          string `gorm:"column:v_pic"`
	Actor        string `gorm:"column:v_actor"`
	VAddtime     int64
	VDaytime     int64
	VWeektime    int64
	VMonthtime   int64 `gorm:"column:v_monthtime"`
	VPublishyear string
	VNote        string
	VDirector    string
	Spic         string `gorm:"column:v_spic"`
	Vip          string `gorm:"column:v_vip"`
}

type Play struct {
	ID   int `gorm:"column:v_id"`
	Tid  int64
	Body string
}

type Content struct {
	ID   int `gorm:"column:v_id"`
	Tid  int64
	Body string
}

func (Content) TableName() string {
	return "sea_content"
}

func (Play) TableName() string {
	return "sea_playdata"
}

func (Video) TableName() string {
	return "sea_data"
}

type FqData struct {
	Data struct {
		AreaName             string   `json:"areaName"`
		Desc                 string   `json:"briefIntroduction"`
		Title                string   `json:"title"`
		Year                 int      `json:"year"`
		CoverHorizontalUrl   string   `json:"coverHorizontalUrl"`
		CoverVerticalUrl     string   `json:"coverVerticalUrl"`
		TagNameList          []string `json:"TagNameList"`
		MovieEpisodeInfoList []struct {
			MovieEpisodeId int `json:"movieEpisodeId"`
		} `json:"movieEpisodeInfoList"`
	} `json:"data"`
}

type Img struct {
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}
