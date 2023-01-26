package main

import (
	"errors"
	"fmt"
	model "golanglearning/new_project/for-gong-zhong-hao/Design-Patterns-practice/Go-gorm-practice/gorm_demo"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

type User1 struct {
	Name string
	ID int
}

func main() {

	rand.Seed(time.Now().UnixNano())
    id := rand.Int()%10000
	user1 := User1{
		ID: id,
		Name: "jiang",

	}


	result := model.DB.Create(&user1)
	fmt.Println(result.RowsAffected, result.Error)
	var userfast User1
	model.DB.First(&userfast)
	fmt.Println(userfast.ID, userfast.Name)

	id = rand.Int()%10000
	var Userlist = []User1{
		{ID: id, Name: "aaaa"},
		{ID: id+1, Name: "bbb"},
		{ID: id+3, Name: "44444"},
	}
	model.DB.Create(&Userlist)

	var userlast User1
	model.DB.Last(&userlast)
	fmt.Println(userlast.ID, userlast.Name)

	var userfast1 User1
	model.DB.First(&userfast1, "5")
	fmt.Println(userfast1.ID, userfast1.Name)

	var userall []User1
	result1 := model.DB.Find(&userall)
	fmt.Println(result1.RowsAffected, result1.Error)

	var userlist []User1
	result = model.DB.Find(&userlist, []int{2,3,4})
	fmt.Println(result.RowsAffected, result.Error)
	checkErrRecordNotFound(result.Error)
	result = model.DB.Where("name = ?", "aaaa").Find(&userlist)
	fmt.Println(result.RowsAffected, result.Error)
	checkErrRecordNotFound(result.Error)

	result = model.DB.Where("name = ? AND id >= ?", "jiang", "1000").Find(&userlist)
	fmt.Println(result.RowsAffected, result.Error)
	checkErrRecordNotFound(result.Error)

	var user3 = User1{ID: 5}
	model.DB.Find(&user3)
	fmt.Println(user3.Name)

	if err := model.DB.Where("name = ?", "jiang").Find(&userlist).Error; err!= nil {
		log.Fatal(err)
	}

	if result := model.DB.Where("name = ?", "jiang").Find(&userlist); result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(result.RowsAffected, result.Error)


}

// 当 First、Last、Take 方法找不到记录时，GORM 会返回 ErrRecordNotFound 错误。
// 如果发生了多个错误，你可以通过 errors.Is 判断错误是否为 ErrRecordNotFound
func checkErrRecordNotFound(err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}
}