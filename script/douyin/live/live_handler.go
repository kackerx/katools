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
	roomApi = "https://webcast-hl.amemv.com/webcast/room/reflow/info/?room_id=7031051525690415910&type_id=0&user_id=%s&live_id=1&app_id=1128"
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

			shortId := gjson.Get(string(ret), "data.user.short_id").Int()
			//followerNum := gjson.Get(string(ret), "data.user.follow_info.follower_count").Int()
			douyinId := gjson.Get(string(ret), "data.user.display_id").String()
			if shortId != 0 {
				fmt.Println(id, " ", douyinId)
				//if _, err := rds.SAdd(context.TODO(), "douyin:live:df:userid", id).Result(); err != nil {
				//    fmt.Println(err)
				//}

				//res = append(res, id)
			}

			return nil
		})
	}

	if err = eg.Wait(); err != nil {
		log.Fatalln(err)
	}
}

func AddMember() {
	for scanner.Scan() {
		id := strings.TrimSpace(scanner.Text())
		fmt.Println(id)
		if _, err := rds.SAdd(context.TODO(), "douyin:live:df:userid", id).Result(); err != nil {
			fmt.Println(err)
		}
	}
}

func Delete(data []int) {

	res, err := rds.SCard(context.TODO(), "douyin:live:df:userid").Result()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)

	a := []int{
		228320054875630,
		109464144885,
		73504363641,
		87988439293,
		1046361476965005,
		109622799295,
		96184438358,
		110583221811,
		96636483240,
		104953177238,
		739298334227771,
	}

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
	r, err := rds.HGetAll(context.TODO(), "douyin:live:df:follower_num").
		Result()
	if err != nil {
		log.Fatalln(err)
	}

	for k, v := range r {
		fmt.Println(k, " ", v)
	}
}

func test() {

}

func main() {
	//IsMember()
	//AllMembers()
	AddMember()
	//IsCorrect()
	//AddMember()

}
