package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}
