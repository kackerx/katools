package main

import (
    "context"
    "fmt"
    "log"
)

func main() {
    
    client := NewEs()
    
    data := map[string]interface{}{
        "keyword": "完美日记",
    }
    
    ret, err := client.Update().
        Index("yuntu_keyword-2021").
        Id("20211216-0f9a01f2608cc62aaff1026d64713b6e").
        Doc(data).
        DocAsUpsert(true).
        Do(context.Background())
    if err != nil {
        log.Fatalln(err)
    }
    
    fmt.Println(ret)
    
}
