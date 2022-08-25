package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Lock struct {
	Name     string
	IsLocked bool
}

var (
	err error
	ctx = context.TODO()
	rds = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
)

func main() {
	redis.NewScript(`
		local name = argv
	`)

}

func Do() {

	uuid := uuid.NewV4()
	isLocked, err := rds.SetNX(ctx, "lock", uuid, time.Duration(10000)).Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	if isLocked {

	}
}
