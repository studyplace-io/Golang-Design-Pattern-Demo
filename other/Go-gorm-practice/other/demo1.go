package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Name string
	ID int
}


func main() {

	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	DBname := "mysql"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDb, err := db.DB()
	sqlDb.SetConnMaxLifetime(10)
	sqlDb.SetMaxOpenConns(100)

	fmt.Println("成功", db)

	user := User{
		ID: 2,
		Name: "jiang",

	}
	result := db.Create(&user)
	fmt.Println(result.RowsAffected, result.Error)

}

