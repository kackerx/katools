package main

import (
    "fmt"
    "github.com/fatih/color"
    "github.com/jroimartin/gocui"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "regexp"
    "sync"
)

var (
    active        int
    nextIndex     int
    viewArr       = []string{"hehedy", "fanqie", "wujin", "miaopai", "lt"}
    hehedyVideos  []video
    fanqieVideos  []video
    ltVideos      = make([]dataPlay, 100)
    miaopaiVideos []video
    wujinVideos   []video
    err           error
    db            *gorm.DB
    wg            sync.WaitGroup
)

func initDb() error {
    fmt.Println("初始化数据库连接中...")
    dsh := "kingvstr_dy:M3KHa7fjRkRjkkmJ@tcp(27.102.66.66:3306)/kingvstr_dy?charset=utf8"
    if db, err = gorm.Open(mysql.Open(dsh), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    }); err != nil {
        return err
    }
    return nil
}

func init() {
   go getHehedy()
   go getFanqie()
   go getMiaopai()
   go getWujin("http://wujinzy.com/")
   if err := initDb(); err != nil {
      fmt.Println(err)
      initDb()
   }

}

func main() {
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        panic(err)
    }
    defer g.Close()
    
    g.Highlight = true
    g.SelFgColor = gocui.ColorGreen
    g.Mouse = true
    //g.Cursor = true
    
    g.SetManagerFunc(layout)
    
    go updatePlay(g)
    
    if err := g.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, updateVideo); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
        log.Panicln(err)
    }
    
    if err := g.SetKeybinding("msg", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
        log.Panicln(err)
    }
    
    if err := g.SetKeybinding("fanqie", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("wujin", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("lt", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("miaopai", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("hehedy", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
        log.Panicln(err)
    }
    
    if err := g.SetKeybinding("hehedy", gocui.KeyCtrlSpace, gocui.ModNone, update); err != nil {
        log.Panicln(err)
    }
    
    if err := g.SetKeybinding("lt", 'c', gocui.ModNone, updateAll); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("hehedy", 'c', gocui.ModNone, updateAll); err != nil {
        log.Panicln(err)
    }
    
    if err := g.SetKeybinding("", 's', gocui.ModNone, insert); err != nil {
        log.Panicln(err)
    }
    
    if err := g.SetKeybinding("miaopai", 'c', gocui.ModNone, updateAll); err != nil {
        log.Panicln(err)
    }
    
    if err := g.SetKeybinding("fanqie", 'c', gocui.ModNone, updateAll); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("wujin", 'c', gocui.ModNone, updateAll); err != nil {
        log.Panicln(err)
    }
    
    if err := g.SetKeybinding("fanqie", gocui.KeyCtrlSpace, gocui.ModNone, update); err != nil {
        log.Panicln(err)
    }
    
    if err := g.SetKeybinding("msg", gocui.KeyCtrlSpace, gocui.ModNone, delMsg); err != nil {
        log.Panicln("bye libao!")
    }
    
    if err := g.MainLoop(); err != nil {
        fmt.Println(color.CyanString("更新完毕! 小李宝"))
    }
}

func insert(g *gocui.Gui, v *gocui.View) error {
    if _, err := g.SetCurrentView("msg"); err != nil {
        return nil
    }
    return nil
}

func updateAll(g *gocui.Gui, v *gocui.View) error {
    switch v.Name() {
    case "hehedy":
        for _, vv := range hehedyVideos {
            wg.Add(1)
            go updateOne(vv)
        }
    case "fanqie":
        for _, vv := range fanqieVideos {
            wg.Add(1)
            go updateOne(vv)
        }
    case "wujin":
        for _, vv := range wujinVideos {
            wg.Add(1)
            go updateOne(vv)
        }
    case "miaopai":
        for _, vv := range miaopaiVideos {
            wg.Add(1)
            go updateOne(vv)
        }
    }
    wg.Wait()
    
    if err := g.DeleteView("updateMsg"); err != nil {
        //fmt.Println(err)
    }
    msg := color.CyanString(" 全部更新完毕")
    updateMsg(g, msg)
    return nil
}

func delMsg(g *gocui.Gui, v *gocui.View) error {
    if err := g.DeleteView("msg"); err != nil {
        panic(err)
    }
    if err := g.DeleteView("updateMsg"); err != nil {
    
    }
    if _, err := g.SetCurrentView("hehedy"); err != nil {
        panic(err)
    }
    return nil
}

func updateMsg(g *gocui.Gui, msg string) error {
    maxX, maxY := g.Size()
    if mv, err := g.SetView("updateMsg", maxX-24, maxY-24, maxX-2, maxY-20); err != nil {
        //fmt.Println(err)
        if err != gocui.ErrUnknownView {
            fmt.Println(err)
            return err
        }
        mv.Wrap = true
        mv.Title = "result"
        title := color.RedString("    操作结果提示")
        fmt.Fprint(mv, title+"\n\n")
        fmt.Fprint(mv, msg)
    }
    
    if _, err := g.SetCurrentView("msg"); err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}

func updateVideo(g *gocui.Gui, v *gocui.View) error {
    _, cy := v.Cursor()
    var l string
    var err error
    if l, err = v.Line(cy); err != nil {
        l = ""
        return err
    }
    var title string
    var flag bool
    
    if err := g.DeleteView("updateMsg"); err != nil {
        //fmt.Println(err)
    }
    
    reg := regexp.MustCompile(`.*?(a?\d+)$`)
    match := reg.FindStringSubmatch(l)
    if len(match) > 1 {
        l = match[1]
    } else {
        v.Clear()
        return nil
    }
    
    //if strings.Contains(l, "a") {
    for _, vv := range fanqieVideos {
        if vv.id == l {
            wg.Add(1)
            go updateOne(vv)
            title = vv.title
            flag = true
        }
    }
    //} else {
    for _, vv := range hehedyVideos {
        if vv.id == l {
            wg.Add(1)
            go updateOne(vv)
            title = vv.title
            flag = true
        }
        //}
    }
    
    if !flag {
        msg := color.CyanString("  该id不存在, 错误!")
        updateMsg(g, msg)
        v.Clear()
        return nil
    }
    wg.Wait()
    v.Clear()
    msg := color.CyanString("  %s:更新完毕!", title)
    updateMsg(g, msg)
    return nil
}

func update(g *gocui.Gui, v *gocui.View) error {
    //_, cy := v.Cursor()
    //title := v.Name()
    
    maxX, maxY := g.Size()
    if cv, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        cv.Editable = true
        cv.Wrap = true
        //fmt.Fprint(cv, title)
        if _, err := g.SetCurrentView("msg"); err != nil {
            return err
        }
    }
    return nil
}

func layout(g *gocui.Gui) error {
    maxX, maxY := g.Size()
    //first
    if _, err := g.SetView("sideUp", 0, 0, 23, maxY/2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
    }
    if v, err := g.SetView("hehedy", 3, 3, 15, 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprint(v, "hehedy")
        g.SetCurrentView("hehedy")
    }
    if v, err := g.SetView("fanqie", 3, 6, 15, 8); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprint(v, "fanqie")
    }
    if v, err := g.SetView("wujin", 3, 9, 15, 11); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprint(v, "wujin")
    }
    if v, err := g.SetView("miaopai", 3, 12, 15, 14); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprint(v, "miaopai")
    }
    if v, err := g.SetView("lt", 3, 15, 15, 17); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprint(v, "lt")
    }
    //second
    if _, err := g.SetView("sideDown", 0, maxY/2, 23, maxY); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
    }
    
    if v, err := g.SetView("out", 24, 0, maxX-1, maxY-1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprint(v, "          欢迎来到, 葫芦葩更新系统v1.0")
    }
    
    if cv, err := g.SetView("msg", maxX/2-30, maxY-5, maxX/2+30, maxY-3); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        cv.Editable = true
        cv.Wrap = true
        //fmt.Fprint(cv, title)
        //if _, err := g.SetCurrentView("msg"); err != nil {
        //	return err
        //}
    }
    
    return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
    nextIndex++
    if nextIndex == len(viewArr) {
        nextIndex = 0
    }
    //nextIndex = active + 1
    //if active == 0 {
    //	nextIndex = 1
    //}
    
    name := viewArr[nextIndex]
    out, err := g.View("out")
    if err != nil {
        panic(err)
    }
    
    out.Clear()
    switch name {
    case "hehedy":
        if hehedyVideos == nil {
            fmt.Fprint(out, "视频加载中...")
        } else {
            for _, item := range hehedyVideos {
                fmt.Fprintf(out, "  %s: %s: %s\n", item.id, item.title, item.play)
            }
        }
    case "fanqie":
        if fanqieVideos == nil {
            fmt.Fprintf(out, "视频加载中...")
        } else {
            for _, item := range fanqieVideos {
                fmt.Fprintf(out, "  %s: %s\n", item.id, item.title)
            }
        }
    case "wujin":
        if wujinVideos == nil {
            fmt.Fprintf(out, "视频加载中...")
        } else {
            for _, item := range wujinVideos {
                fmt.Fprintf(out, "  %s: %s\n", item.title, item.note)
            }
        }
    case "miaopai":
        if miaopaiVideos == nil {
            fmt.Fprintf(out, "视频加载中...")
        } else {
            for _, item := range miaopaiVideos {
                fmt.Fprintf(out, "  %s: %s\n", item.title, item.year)
            }
        }
    case "lt":
        if ltVideos == nil {
            fmt.Fprintf(out, "视频加载中...")
        } else {
            for _, item := range ltVideos {
                fmt.Fprintf(out, "  %s\n", item.VName)
            }
        }
    }
    
    if _, err := setCurrentViewOnTop(g, name); err != nil {
        return err
    }
    
    active = nextIndex
    return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}
