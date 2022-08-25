package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	err error
	db  *gorm.DB
)

type TMG struct {
	ID   uint
	Name string
}

func main() {
	db, _ = NewDb()
	db.AutoMigrate(&TMG{})

	db.Transaction(func(tx *gorm.DB) error {
		t1 := TMG{Name: "1"}
		t2 := TMG{Name: "2"}
		if true {
			return errors.New("事务出错")
		}
		t3 := TMG{Name: "3"}
		tx.Create([]TMG{t1, t2, t3})
		return nil
	})

}

func NewDb() (db *gorm.DB, err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "t_",  // 配置表前缀
			SingularTable: false, // 配置表名单复数t_user, 否则t_users
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止自动外键约束
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)                  // 设置MySQL的最大空闲连接数（推荐100）
	sqlDB.SetMaxOpenConns(100)                 // 设置MySQL的最大连接数（推荐100）
	sqlDB.SetConnMaxLifetime(time.Second * 10) // 设置MySQL的空闲连接最大存活时间（推荐10s）

	return
}
