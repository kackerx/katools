package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func main() {
    //f, err := os.Open("std-party/dir/test.txt")
    //if err != nil {
    //    log.Fatalln(err)
    //}
    
    bs, err := ioutil.ReadFile("std-party/dir/test.txt")
    if err != nil {
        log.Fatalln(err)
    }
    
    b := string(bs)
    
    bb := strings.Split(b, "\n")
    for _, b := range bb {
        fmt.Println(b)
    }
    
    
}
