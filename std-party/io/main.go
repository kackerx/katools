package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    var uid interface{} = "kacker"
    u, ok := uid.(string)
    if ok {
        fmt.Println(ok)
    }
    fmt.Println(u)
    //bufWrite()
}

type User struct {
    Name string
}

type T interface {
}

func bufWrite() {
    f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln(err)
    }

    defer func() {
        f.Sync()
        f.Close()
    }()

    data := []byte("kacker")

    bf := bufio.NewWriterSize(f, 32)

    n, err := bf.Write(data)
    if err != nil {
        log.Fatalln(err)
    }

    bf.Flush()
    fmt.Println("write: ", n)

}

func read() {
    f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
    if err != nil {
        log.Fatalln(err)
    }

    defer f.Close()

    data := make([]byte, 6)

    n, err := f.Read(data)
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println("read: ", n)

    fmt.Println(string(data))
}

func write() {
    f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln(err)
    }

    defer func() {
        f.Sync()
        f.Close()
    }()

    data := []byte("kacker")

    n, err := f.Write(data)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("write: ", n)
}
