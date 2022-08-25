package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rds = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
)

func main() {

	//rds.Watch(ctx, func(tx *redis.Tx) error {
	//    fmt.Println("k1")
	//    return nil
	//}, "k1", "k2")

}
