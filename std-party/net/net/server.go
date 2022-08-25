package main

import (
    "fmt"
    "net"
)

func main() {

    listener, err := net.Listen("tcp", "127.0.0.1:9999")
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }

        go handle(conn)
    }
}

func handle(conn net.Conn) bool {
    defer conn.Close()
    var msg = make([]byte, 2048)
    n, err := conn.Read(msg)
    if err != nil {
        fmt.Println(err)
        return true
    }

    fmt.Printf("read len\n%d: %s\n", n, string(msg))

    fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", 11)
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprint(conn, "\r\n")
    fmt.Fprint(conn, "helloworld")

    // 消息体
    //body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title>Go example</title></head><body><strong>Hello World</strong></body></html>`
    // HTTP 协议及请求码
    //fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
    //// 内容长度
    //fmt.Fprintf(conn, "Content-Length: %d\r\n", 11)
    //// 内容类型
    //fmt.Fprint(conn, "Content-Type: text/html\r\n")
    //fmt.Fprint(conn, "\r\n")
    //fmt.Fprint(conn, "helloword")

    //res = "hello world"
    //n, err = conn.Write([]byte(res))
    if err != nil {
        fmt.Println(err)
        return true
    }

    //fmt.Printf("write len\n%d: %s\n", n, res)
    return false
}
