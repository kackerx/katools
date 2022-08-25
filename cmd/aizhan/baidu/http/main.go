package main

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "net/http"
)

var (
    rds *redis.Client
    ctx = context.Background()
)

func init() {
    rds = redis.NewClient(&redis.Options{
        Addr: "127.0.0.1:6379",
        DB:   0,
    })
}

func main() {

    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    r.Handle(http.MethodPost, "/put", func(c *gin.Context) {
        ips := c.PostFormArray("ip")

        for i, v := range ips {
            _, err := rds.Set(ctx, fmt.Sprintf("proxy_%d", i), v, 0).Result()
            if err != nil {
                fmt.Println(err)
                c.JSON(http.StatusOK, gin.H{
                    "errmsg": err,
                })
                return
            }
        }

        c.JSON(http.StatusOK, gin.H{
            "success": "ok",
        })
    })

    r.GET("/get", func(c *gin.Context) {
        var proxys []string
        for i := 0; i < 50; i++ {
            proxy, err := rds.Get(ctx, fmt.Sprintf("proxy_%d", i)).Result()
            if err != nil {
                fmt.Println(err)
                c.JSON(http.StatusOK, gin.H{
                    "errmsg": err,
                })
                return
            }
            proxys = append(proxys, proxy)
        }

        c.JSON(http.StatusOK, gin.H{
            "success": "ok",
            "ips":     proxys,
        })
    })

    r.Run("0.0.0.0:9999")
}

func GetIp() {
    res, err := rds.Get(ctx, "test").Result()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(res)
}
