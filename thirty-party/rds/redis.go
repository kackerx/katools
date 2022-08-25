package rds

import (
    "github.com/go-redis/redis/v8"
    "time"
)

var C = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: 0, DialTimeout: time.Second * 2, ReadTimeout: time.Millisecond * 3000, WriteTimeout: time.Second * 3, PoolSize: 50})
