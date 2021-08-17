package main

type Video struct {
	ID           int    `gorm:"column:v_id"`
	Tid          int    `gorm:"column:tid"`
	Name         string `gorm:"column:v_name"`
	Pic          string `gorm:"column:v_pic"`
	Actor        string `gorm:"column:v_actor"`
	Vip          int `gorm:"column:v_vip"`
	VAddtime     int64
	VDaytime     int64
	VWeektime    int64
	VMonthtime   int64 `gorm:"column:v_monthtime"`
	VPublishyear int64
	VNote        string
	VDirector    string
}

type Play struct {
	ID   int `gorm:"column:v_id"`
	Tid  int64
	Body string
}

func (Play) TableName() string {
	return "sea_playdata"
}

func (Video) TableName() string {
	return "sea_data"
}
