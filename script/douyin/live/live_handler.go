package main

import (
    "bufio"
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "github.com/kackerx/katools/middleware"
    "github.com/kackerx/katools/pkg"
    "github.com/tidwall/gjson"
    "golang.org/x/sync/errgroup"
    "log"
    "os"
    "strings"
)

var (
    res     []string
    eg      errgroup.Group
    err     error
    ret     bool
    f       *os.File
    rds     *redis.Client
    scanner *bufio.Scanner
    roomApi = "https://webcast-hl.amemv.com/webcast/room/reflow/info/?room_id=7091773472690539267&type_id=0&user_id=%s&live_id=1&app_id=1128"
)

func init() {
    rds = middleware.NewRecli()
    f, err = os.Open("script/douyin/live/test.txt")
    if err != nil {
        log.Fatalln(err)
    }
    scanner = bufio.NewScanner(f)
}

func IsMember() {
    for scanner.Scan() {
        uid := scanner.Text()
        uid = strings.TrimSpace(uid)

        if ret, err = rds.SIsMember(context.TODO(), "douyin:live:df:userid", uid).Result(); err != nil {
            fmt.Println(err)
        }

        if !ret {
            fmt.Println(uid)
        }
    }
}

func IsCorrect() {
    for scanner.Scan() {
        id := strings.TrimSpace(scanner.Text())
        url := fmt.Sprintf(roomApi, id)
        eg.Go(func() error {
            ret, err := pkg.Get(url)
            //fmt.Println(url)
            if err != nil {
                return err
            }

            retStr := string(ret)

            shortId := gjson.Get(retStr, "data.user.short_id").Int()
            //followerNum := gjson.Get(string(ret), "data.user.follow_info.follower_count").Int()
            douyinId := gjson.Get(retStr, "data.user.display_id").String()
            if shortId != 0 {
                //fmt.Println(id, "正确: ", shortId, ": ", douyinId)
                //if _, err := rds.SAdd(context.TODO(), "douyin:live:df:userid", id).Result(); err != nil {
                //    fmt.Println(err)
                //}

                //res = append(res, id)
            } else {
                fmt.Println(id, "错误: ", shortId, ": ", douyinId)
            }

            return nil
        })
    }

    if err = eg.Wait(); err != nil {
        log.Fatalln(err)
    }
}

func AddMember() {
    n, err := rds.SCard(context.TODO(), "douyin:live:df:userid").Result()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("插入前: ", n)

    for scanner.Scan() {
        id := strings.TrimSpace(scanner.Text())
        fmt.Println(id)
        if _, err := rds.SAdd(context.TODO(), "douyin:live:df:userid", id).Result(); err != nil {
            fmt.Println(err)
        }
    }

    n, err = rds.SCard(context.TODO(), "douyin:live:df:userid").Result()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("插入后: ", n)
}

func Delete() {

    res, err := rds.SCard(context.TODO(), "douyin:live:df:userid").Result()
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(res)

    a := []interface{}{3202239665809261, 102158250898, 1469422174213790, 31442173066, "1293923266998934023168", "61297086941m"}

    for _, v := range a {
        res, err = rds.SRem(context.TODO(), "douyin:live:df:userid", v).Result()
        if err != nil {
            log.Fatalln(err)
        }

        fmt.Println(res, "deleted")
    }

    res, err = rds.SCard(context.TODO(), "douyin:live:df:userid").Result()
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(res)

}

func GetFollower() {

    if rest, err := rds.HGetAll(context.TODO(), "douyin:live:df:follower_num").Result(); err != nil {
        fmt.Println(err)
    } else {
        for k, v := range rest {
            fmt.Println(k, "---", v)
        }
    }

}

func AllMembers() {
    r, err := rds.SMembers(context.TODO(), "douyin:live:df:userid").Result()
    if err != nil {
        log.Fatalln(err)
    }

    for _, v := range r {
        fmt.Println(v)
    }
}

func main() {
    //Delete()
    //IsMember()
    AllMembers()
    //AddMember()
    //IsCorrect()
    //AddMember()

}
