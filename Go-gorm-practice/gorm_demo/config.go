package gorm_demo

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
)

// 预先声明
var (
	Db 		 	string
	DbHost   	string
	DbPort   	string
	DbUser   	string
	DbPassWord 	string
	DbName   	string
)

func init() {
	workdir, _ := os.Getwd()
	wrpath := "/Go-gorm-practice/gorm_demo/"
	var str = []string{workdir, wrpath, "config.ini"}
	path := strings.Join(str, "")

	file, err := ini.Load(path)
	checkErr(err)
	LoadMysqlData(file)
	dbConfig := []string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ")/", DbName, "?charset=utf8&parseTime=true"}
	pathDB := strings.Join(dbConfig, "")
	DataBaseInit(pathDB)



}

// 取配置文件
func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()

}

// 错误处理
func checkErr(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

