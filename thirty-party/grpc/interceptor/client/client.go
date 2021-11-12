package main

import (
    "context"
    "fmt"
    grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "learn_go/thirty_party/grpc/interceptor/proto"
    "time"
)

func main() {
    
    opts := []grpc.DialOption{}
    
    interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
        start := time.Now()
        err := invoker(ctx, method, req, reply, cc, opts...)
        fmt.Println("耗时: ", time.Since(start))
        return err
    }
    
    opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
    
    retryOpts := []grpc_retry.CallOption{
        grpc_retry.WithMax(3),
        grpc_retry.WithPerRetryTimeout(1 * time.Second),
        grpc_retry.WithCodes(codes.Unknown, codes.Unavailable, codes.DeadlineExceeded),
    }
    opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)))
    
    opts = append(opts, grpc.WithInsecure())
    
    conn, err := grpc.Dial("127.0.0.1:9999", opts...)
    if err != nil {
        panic(err)
    }
    
    client := proto.NewHelloClient(conn)
    
    resp, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "kacker"})
    if err != nil {
        panic(err)
    }
    
    fmt.Println(resp.Name)
}
