package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 范例一
	var m1 int = 2022
	valueOfm := reflect.ValueOf(m1)
	fmt.Println(valueOfm)

	getm1 := valueOfm.Interface().(int)
	getm2 := int(valueOfm.Int())
	fmt.Printf("getm1: %v, getm2: %v", getm1, getm2)
	fmt.Println("--------------------------------------------------------")

	// 范例二
	type example3 struct {
		a int
		b string
		float32
		bool
		next *example3
	}

	instance := example3{
		next: &example3{},
	}
	valueOfinstance := reflect.ValueOf(instance)
	fmt.Println("NumField: ", valueOfinstance.NumField())

	float32Field := valueOfinstance.Field(2)
	fmt.Printf("float32Field: %v, float32FieldType: %v \n", float32Field, float32Field.Type())
	boolField := valueOfinstance.Field(3)
	fmt.Printf("boolField: %v, boolFieldType: %v \n", boolField, boolField.Type())
	nextField := valueOfinstance.FieldByIndex([]int{4, 3}).Type()
	fmt.Println("FieldByIndex([]int{4, 3}).Type()", nextField)
	fmt.Println("--------------------------------------------------------")

	// 范例三
	var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())
	fmt.Println("--------------------------------------------------------")


	// 范例四

	var aa int = 2024
	//valueOfa := reflect.ValueOf(aa)	// 会panic
	valueOfa := reflect.ValueOf(&aa)
	valueOfa = valueOfa.Elem()	// 指针都需要加上Elem()方法。
	valueOfa.SetInt(200)
	fmt.Println(valueOfa.Int())
	fmt.Println("--------------------------------------------------------")


	// 范例五
	type example4 struct {
		LegCount int
	}
	// 如果是用指针，记得后面要加上Elem()方法，才能找到。
	valueOfexample4 := reflect.ValueOf(&example4{LegCount: 100}).Elem()
	vLegCount := valueOfexample4.FieldByName("LegCount")
	fmt.Println(vLegCount)
	vLegCount.SetInt(1000)
	fmt.Println(vLegCount)
	fmt.Println("--------------------------------------------------------")

	// 范例六
	type example5 struct {
		ID int
		Name string
	}
	instanceExample5 := example5{
		ID: 100,
		Name: "jiang",
	}
	valueOfexample5 := reflect.TypeOf(instanceExample5)
	instanceCopy := reflect.New(valueOfexample5).Elem()	// new出来会是没有初始化的struct，并且记得需要加上Elem()方法。
	for i := 0; i < instanceCopy.NumField(); i++ {
		fmt.Println(i)
		fmt.Println(instanceCopy.Field(i))
	}
	fmt.Println("--------------------------------------------------------")

	// 范例七
	funcValue := reflect.ValueOf(add)
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(30)}
	resList := funcValue.Call(paramList)
	fmt.Println(resList[0].Int())

	funcValue2 := reflect.ValueOf(reduce)
	paramList2 := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(30)}
	resList2 := funcValue2.Call(paramList2)
	fmt.Println(resList2[0].Int())



}

func add(a, b int) int {

	return a + b
}

func reduce(a, b int) int {
	return a - b
}