package middleware

import "github.com/go-redis/redis/v8"

func NewRecli() *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr:         "10.0.3.100:6379",
        Password:     "EfcHGSzKqg6cfzWq",
        DB:           0,
        PoolSize:     70,
        MinIdleConns: 50,
    })
}
