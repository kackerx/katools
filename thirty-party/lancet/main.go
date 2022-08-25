package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
    "github.com/gin-gonic/gin"
    "net/http"
)

type Duck interface {
    wang() string
}

type Person struct {
    name string `json:"name" gorm:"not null"`
    Age  int    `json:"age"`
}

func (p *Person) wang() string {
    return p.name
}

func (p *Person) GetName(i int) string {
    return p.name
}

func main() {

    s := "hello"
    rs := strutil.ReverseStr(s)
    fmt.Println(rs)

    r := gin.Default()
    r.Handle(http.MethodGet, "/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "test": "success",
        })
    })
    r.Run(":9999")

    var p = &Person{name: "kacker", Age: 0}

    p.GetName(1)
}
