package main

import "golanglearning/new_project/for-gong-zhong-hao/Design-Patterns-practice/SubjectObjectPattern"

/*
  观察者模式:提供了一种作用于任何实现了订阅者接口的对象的机制， 可对其事件进行订阅和取消订阅。

  定义一种订阅机制， 可在对象事件发生时通知多个 “观察” 该对象的其他对象。


  客户只订阅其感兴趣的特定商品， 商品可用时便会收到通知。 同时，多名客户也可订阅同一款产品。
  1. 会在有任何事发生时发布事件的主体。
  2. 订阅了主体事件并会在事件发生时收到通知的观察者。
 */


func main() {

	// 建立一个被观察者
	shirtItem := SubjectObjectPattern.NewItem("Nike Shirt")

	// 建立两个观察者
	observerFirst := &SubjectObjectPattern.Customer{Id: "abc@gmail.com"}
	observerSecond := &SubjectObjectPattern.Customer{Id: "xyz@gmail.com"}

	// 注册一下
	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()
}
