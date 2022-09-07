package main

import (
	"fmt"
	model "golanglearning/new_project/for-gong-zhong-hao/Go-gorm-practice/gorm_demo"
	"math/rand"
)

type User1 struct {
	Name string
	ID int
}

func main() {



	user := User1{
		ID: int(rand.Int31()%10000),
		Name: "jiang",

	}


	result := model.DB.Create(&user)
	fmt.Println(result.RowsAffected, result.Error)
	var usernil User1
	model.DB.First(&usernil)
	fmt.Println(usernil.ID, usernil.Name)

}
