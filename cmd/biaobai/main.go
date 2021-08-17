package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

func main() {
    r := gin.Default()
    r.Static("/js", "template/js")
    r.Static("/css", "template/css")
    r.LoadHTMLGlob("template/index.html")
    r.Handle(http.MethodGet, "/", func(ctx *gin.Context) {
        ctx.HTML(http.StatusOK, "index.html", gin.H{
            "time": time.Now(),
        })
    })
    r.Run()
}
