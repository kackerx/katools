package main

import (
    "context"
    "errors"
    "fmt"
    "github.com/go-redis/redis/v8"
    uuid "github.com/satori/go.uuid"
    "sync"
    "time"
)

var (
    c     = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: 0, DialTimeout: time.Second * 2, ReadTimeout: time.Millisecond * 3000, WriteTimeout: time.Second * 3, PoolSize: 50})
    count int
    lock  = Lock{
        rds:     c,
        Name:    "k",
        Expires: time.Second * 10,
        ctx:     context.Background(),
        Value:   uuid.NewV4().String(),
    }
)

type Lock struct {
    ctx     context.Context
    rds     *redis.Client
    Name    string
    Value   string
    Expires time.Duration
}

func (l *Lock) TryLock() error {
    isLocked, err := l.rds.SetNX(l.ctx, l.Name, l.Value, l.Expires).Result()
    if err != nil {
        return err
    }

    if !isLocked {
        return errors.New("lock is 占用")
    }

    return nil
}

func (l *Lock) UnLock() error {
    _, err := l.rds.Del(l.ctx, l.Name).Result()
    if err != nil {
        return err
    }

    return nil
}

func GetSource() {
    for i := 0; i < 10000; i++ {
        count++
    }
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            //GetSource()
            for {
                err := lock.TryLock()
                if err != nil {
                    fmt.Println(err)
                    time.Sleep(time.Millisecond * 10)
                    continue
                }

                GetSource()

                err = lock.UnLock()
                fmt.Println(err)
                break
            }
        }()
    }

    wg.Wait()
    fmt.Println(count)
}
