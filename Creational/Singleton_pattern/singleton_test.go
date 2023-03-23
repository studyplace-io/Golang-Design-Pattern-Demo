package Singleton_pattern

import (
	"fmt"
	"testing"
)

func TestSingletonPattern(t *testing.T) {

	// 使用GetInstance()拿到的都是唯一初始化的对象
	in := GetInstance()
	in.SetInstance("这是一个单例模式")
	fmt.Println(in.GetInstanceInfo())
	in.SetInstance("这是一个单例模式1111")
	in1 := GetInstance()
	fmt.Println(in1.GetInstanceInfo())
}