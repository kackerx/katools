package main

import (
    "github.com/olivere/elastic/v7"
    "log"
    "os"
    "time"
)

func NewEs() *elastic.Client {
    options := []elastic.ClientOptionFunc{
        elastic.SetURL("http://@10.0.4.101:4000"),
        elastic.SetSniff(true),
        elastic.SetHealthcheckInterval(time.Second * 10),
        elastic.SetGzip(false),
        elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC: ", log.LstdFlags)),
        elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
    }
    
    options = append(options, elastic.SetBasicAuth("eswriter", "ZDM4MTUwZDNi"))
    client, err := elastic.NewClient(options...)
    if err != nil {
        log.Fatalln(err)
    }
    
    return client
}
