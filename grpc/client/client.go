package main

import (
    "context"
    "fmt"
    "github.com/kackerx/katools/proto"
    "google.golang.org/grpc"
    "log"
)

func main() {
    conn, err := grpc.Dial(":9999", grpc.WithInsecure())
    if err != nil {
        log.Fatalln(err)
    }
    
    client := proto.NewHelloClient(conn)
    resp, err := client.SayHello(context.Background(), &proto.Req{Msg: "kacker"})
    if err != nil {
        log.Fatalln(err)
    }
    
    fmt.Println(resp.Ans)

}
