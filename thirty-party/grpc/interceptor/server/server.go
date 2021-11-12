package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "learn_go/thirty_party/grpc/interceptor/proto"
    "net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (reply *proto.HelloReply, err error) {
    return &proto.HelloReply{Name: "hello" + request.Name}, nil
}

func main() {
    // 拦截器
    interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
        fmt.Println("收到一个请求")
        return handler(ctx, req)
    }
    opt := grpc.UnaryInterceptor(interceptor)
    
    server := grpc.NewServer(opt)
    proto.RegisterHelloServer(server, &Server{}) // 注册对象
    
    listener, err := net.Listen("tcp", "127.0.0.1:9999")
    if err != nil {
        panic(err)
    }
    
    err = server.Serve(listener)
    if err != nil {
        panic(err)
    }
    
}
