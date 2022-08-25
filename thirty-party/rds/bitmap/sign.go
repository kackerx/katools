package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "github.com/kackerx/katools/thirty-party/rds"
    "time"
)

var (
    year, month, day = time.Now().Date()
    rdsKey           = fmt.Sprintf(`id:%d:%d:%d`, year, month)
)

func Sign() {
    _, err := rds.C.SetBit(context.Background(), rdsKey, int64(day-1), 1).Result()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func IsSign() (isSign int64) {
    isSign, err := rds.C.GetBit(context.Background(), rdsKey, int64(day-1)).Result()
    if err != nil {
        fmt.Println(err)
    }
    return
}

func SignCount() {
    result, err := rds.C.BitCount(context.Background(), rdsKey, &redis.BitCount{
        Start: 0,
        End:   100,
    }).Result()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
}

func main() {
    //Sign()

    fmt.Println(IsSign())
    SignCount()
}
