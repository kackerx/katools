package main

import (
    "fmt"
    "github.com/spf13/pflag"
)

var (
    flagVar = pflag.Int("name", 123, "help usage for name")
)

func main() {
    pflag.Parse()
    
    fmt.Println(pflag.NArg())
    fmt.Println(pflag.Args())
    fmt.Println(pflag.Arg(0))
    

}

