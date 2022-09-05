package main

import (
	"fmt"
	"reflect"
)

// 主要是
// 接口对象 := reflect.TypeOf(类型)

func main() {

	// 范例一
	var m string
	typeOfm := reflect.TypeOf(m)
	fmt.Printf("name: %v, kind: %v \n", typeOfm.Name(), typeOfm.Kind()) // name: string, kind: string
	fmt.Println("--------------------------------------------------------")


	// 范例二
	type example1 struct {}

	// 这里可以比较一下，不用指针的结果，在Elem()会报错。 reflect.TypeOf(example{})
	typeOfexample := reflect.TypeOf(&example1{}) // 这里只会取指针类型，不会显示struct类型
	fmt.Printf("name: %v, kind: %v \n", typeOfexample.Name(), typeOfexample.Kind()) // name: , kind: ptr

	typeOfexampleElem := typeOfexample.Elem() // 调用这个才会显示struct类型
	fmt.Printf("element name: %v, element kind: %v \n", typeOfexampleElem.Name(), typeOfexampleElem.Kind()) // element name: example, element kind: struct
	fmt.Println("--------------------------------------------------------")

	// 范例三
	type example2 struct {
		Name string
		Type int `json:"type" id:"100"`
	}
	instance := example2{
		Name: "jiang",
		Type: 100,
	}
	typeOfinstance := reflect.TypeOf(instance)
	for i := 0; i < typeOfinstance.NumField(); i++ {
		fieldType := typeOfinstance.Field(i)
		fmt.Printf("name: %v, tag: '%v' \n", fieldType.Name, fieldType.Tag)
	}

	if typeOfinstance1, ok := typeOfinstance.FieldByName("Name"); ok {
		fmt.Println(typeOfinstance1.Name, typeOfinstance1.Type)
	}
	if typeOfinstance2, ok := typeOfinstance.FieldByName("Type"); ok {
		fmt.Println(typeOfinstance2.Tag.Get("json"), typeOfinstance2.Tag.Get("id"))
	}


}
