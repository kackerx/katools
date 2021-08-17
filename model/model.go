package model

type Video struct {
    Vid          int    `db:"v_id"`
    Tid          int    `db:"tid"`
    Name         string `db:"v_name"`
    Pic          string `db:"v_pic"`
    Actor        string `db:"v_actor"`
    Vaddtime     int64  `db:"v_addtime"`
    Vdaytime     int64  `db:"v_daytime"`
    Vweektime    int64  `db:"v_weektime"`
    Vmonthtime   int64  `db:"v_monthtime"`
    VPublishyear int64  `db:"v_publishyear"`
    VNote        string `db:"v_note"`
    VDirector    string `db:"v_director"`
}

type Play struct {
    ID   int `gorm:"column:v_id"`
    Tid  int64
    Body string
}
