package main

import (
	"fmt"
	"reflect"
)

// 主要是
// 接口对象 := reflect.ValueOf(类型)

func main() {

	// 范例一：从reflect.ValueOf(m1)中获取值。
	var m1 int = 2022
	valueOfm := reflect.ValueOf(m1)
	fmt.Println(valueOfm)

	// 可以用空接口+类型断言拿到值。
	getm1 := valueOfm.Interface().(int)
	// 可以强制类型转换。
	getm2 := int(valueOfm.Int())
	fmt.Printf("getm1: %v, getm2: %v", getm1, getm2)
	fmt.Println("--------------------------------------------------------")

	// 范例二：当反射体是struct时，访问字段的方法。
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
	// 反射后取到valueOfinstance
	valueOfinstance := reflect.ValueOf(instance)
	fmt.Println("NumField: ", valueOfinstance.NumField())	// 数量

	float32Field := valueOfinstance.Field(2)	// 按照index取字段的value
	fmt.Printf("float32Field: %v, float32FieldType: %v \n", float32Field, float32Field.Type()) // 查看
	boolField := valueOfinstance.Field(3)	// 按照index取字段的value
	fmt.Printf("boolField: %v, boolFieldType: %v \n", boolField, boolField.Type())	// 查看
	nextField := valueOfinstance.FieldByIndex([]int{4, 3}).Type()	// 直接取嵌套类型的value
	fmt.Println("FieldByIndex([]int{4, 3}).Type()", nextField)
	fmt.Println("--------------------------------------------------------")

	// 范例三：判断有效性
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


	// 范例五：利用反射修改原值 需要有两个条件：1. 指针类型 2. 字段是大写
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

	// 范例六：反射创建新类型
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

	// 范例七：使用反射调用函数
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