package main

import (
    "fmt"
    "github.com/jroimartin/gocui"
    "strings"
    "time"
)

type dataPlay struct {
    VId   int    `json:"v_id"`
    VName string `json:"v_name"`
    Body  string `json:"body"`
    VPic  string `json:"v_pic"`
}

func updatePlay(g *gocui.Gui) {
    n, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
    ret := db.Table("sea_data").Select("sea_data.v_id, sea_data.v_name, sea_data.v_pic, sea_playdata.body").
        Joins("left join sea_playdata on sea_data.v_id = sea_playdata.v_id where sea_playdata.body like ? and sea_data.v_addtime > ?", "%龙腾%", n.Unix()).
        Limit(200).Find(&ltVideos)
    
    //ret := db.Where("body like ?", "%龙腾%").Limit(66).Find(&ltVideos)
    var count int
    if ret.RowsAffected > 0 {
        for _, v := range ltVideos {
            s := v.Body
            s = strings.Replace(s, "###", "", -1)
            var flag bool
            
            if strings.HasSuffix(s, ".m3u8") || strings.HasSuffix(s, ".m3u8#") {
                s = strings.Replace(s, ".m3u8", ".m3u8$ltm3u8", -1)
                flag = true
                count++
            } else if strings.Contains(s, ".m3u8$$$") || strings.Contains(s, ".m3u8#$$$") {
                count++
                flag = true
                s = strings.Replace(s, ".m3u8$$$", ".m3u8$ltm3u8$$$", -1)
                s = strings.Replace(s, ".m3u8#", ".m3u8$ltm3u8#", -1)
            }
            if flag {
                db.Model(&Play{}).
                    Where("v_id = ?", v.VId).
                    Update("Body", s)
    
                if !strings.Contains(v.VPic, "gtimg") && !strings.Contains(v.VPic, "mgtv") {
                    img, _ := UploadImgTx(v.VPic)
                    db.Model(&Video{}).
                        Where("v_id = ?", v.VId).
                        Update("v_pic", img)
                }
            }
        }
        if count > 0 {
            updateMsg(g, fmt.Sprintf("%d: 更新完毕", count))
        }
    }
}

func updateOne(v video) error {
    defer wg.Done()
    newVideo := Video{
        Name:         v.title,
        Pic:          v.img,
        Actor:        v.actor,
        VAddtime:     time.Now().Unix(),
        VDaytime:     time.Now().Unix(),
        VWeektime:    time.Now().Unix(),
        VMonthtime:   time.Now().Unix(),
        VPublishyear: v.year,
        VNote:        v.note,
        VDirector:    v.director,
        Spic:         v.spic,
        Tid:          v.tid,
    }
    
    var video Video
    var playData Play
    var contentData Content
    result := db.Where("v_name = ?", v.title).First(&video)
    if result.RowsAffected == 1 {
        //已存在先查看播放地址是否存在不同
        db.Where("v_id = ?", video.ID).First(&playData)
        playBody := isPlayEqual(v.play, playData.Body)
        
        db.Model(&video).Updates(newVideo)
        
        playData.ID = video.ID
        db.Model(&playData).Updates(Play{
            ID:   video.ID,
            Body: playBody,
        })
        
        contentData.ID = video.ID
        db.Model(&contentData).Updates(Play{
            ID:   video.ID,
            Body: v.content,
        })
        
    } else {
        newVideo.Tid = v.tid
        result := db.Create(&newVideo)
        if result.Error != nil {
            fmt.Println(result.Error)
        }
        
        db.Create(&Play{
            ID:   newVideo.ID,
            Tid:  int64(v.tid),
            Body: v.play,
        })
        db.Create(&Content{
            ID:   newVideo.ID,
            Tid:  int64(v.tid),
            Body: v.content,
        })
    }
    return nil
}

func isPlayEqual(play, rawPlay string) string {
    playRet := make([]string, 1)
    if strings.Contains(rawPlay, "$$$") {
        playRet = strings.Split(rawPlay, "$$$")
    } else {
        playRet[0] = rawPlay
    }
    
    var flag bool
    for i, v := range playRet {
        playType := strings.Split(v, "$$")[0]
        if strings.Contains(play, playType) {
            flag = true
            playRet[i] = play
            //fmt.Println("替换")
        }
    }
    if !flag {
        //fmt.Println("新增")
        playRet = append(playRet, play)
    }
    
    return strings.Replace(strings.Trim(fmt.Sprint(playRet), "[]"), " ", "$$$", -1)
    
}

func getPlay(play string) string {
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
    return strings.Replace(strings.Trim(fmt.Sprint(playRet), "[]"), " ", "$$$", -1)
}
