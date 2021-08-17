package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "github.com/kackerx/katools/model"
)

var (
	err error
	Db  *sqlx.DB
)

func initDb() (err error) {
	dsn := "read:Wasd4044516520@tcp(101.33.117.86:3306)/kingvstr_dy?charset=utf8"
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}

	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(10)
	return
}

func Testmain() {
    if err = initDb(); err != nil {
        panic(err)
    }
    
    var vs []model.Video
    sqlStr := `select v_id, v_name from sea_data limit 0, 5`
    err := Db.Select(&vs, sqlStr)
    if err != nil {
        panic(err)
    }
    
    fmt.Println(vs)
}

func main() {
    Testmain()
}
