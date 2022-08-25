package main

import (
    "context"
    "fmt"
    "github.com/kackerx/katools/proto"
    "google.golang.org/grpc"
    "log"
    "net"
)

type Server struct {
    Name string
}

func (s *Server) SayHello(ctx context.Context, req *proto.Req) (*proto.Res, error) {
    return &proto.Res{Ans: "hellow" + req.Msg}, nil
}

func main() {
    server := grpc.NewServer()
    proto.RegisterHelloServer(server, &Server{})

    fmt.Println("监听: ")
    listener, err := net.Listen("tcp", ":9999")
    if err != nil {
        log.Fatalln(err)
    }

    server.Serve(listener)
}

func test(st string) {
}
