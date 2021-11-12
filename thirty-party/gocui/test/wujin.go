package main

import (
    "bytes"
    "github.com/PuerkitoBio/goquery"
    "strconv"
    "strings"
    "time"
)

func getWujin(url string) {
    resp, err := Get(url)
    if err != nil {
        panic(err)
    }
    
    doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
    if err != nil {
        panic(err)
    }
    
    lastTime := doc.Find(".xing_vb7").Last().Text()
    lastT, _ := time.Parse("2006-01-02 03:04:05", lastTime)
    floor, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
    if lastT.After(floor) {
        go getWujin("http://wujinzy.com/index.php/index/index/page/2.html")
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
                actorRune := []rune(actor)
                if len(actorRune) > 20 {
                    actor = string(actorRune[0:15])
                }
                area := selection.Find(".vodinfobox ul li:nth-child(5) span").Text()
                year := selection.Find(".vodinfobox ul li:nth-child(7) span").Text()
                if year == "" {
                    year = strconv.Itoa(time.Now().Year())
                }
                note := selection.Find(".vodh span").Text()
                desc := selection.Find(".vodplayinfo").Text()
                category = strings.Split(category, "/")[0]
                tid := GetTid(strings.Replace(category, "片", "", 0))
                
                txImg, _ := UploadImgTx(img)
                wujinVideos = append(wujinVideos, video{
                    tid:      tid,
                    img:      txImg,
                    title:    title,
                    play:     play,
                    actor:    actor,
                    director: director,
                    area:     area,
                    year:     year,
                    
                    note:    note,
                    content: desc,
                })
            })
        }("http://wujinzy.com" + content)
    })
}