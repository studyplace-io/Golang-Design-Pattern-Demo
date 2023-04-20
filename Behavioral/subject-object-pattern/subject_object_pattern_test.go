package subject_object_pattern

import "testing"

/*
  观察者模式:提供了一种作用于任何实现了订阅者接口的对象的机制， 可对其事件进行订阅和取消订阅。

  定义一种订阅机制， 可在对象事件发生时通知多个 “观察” 该对象的其他对象。

  客户只订阅其感兴趣的特定商品， 商品可用时便会收到通知。 同时，多名客户也可订阅同一款产品。
  1. 会在有任何事发生时发布事件的主体。
  2. 订阅了主体事件并会在事件发生时收到通知的观察者。

*/

func TestSubjectObjectPattern(t *testing.T) {
	// 建立一个被观察者
	shirtItem := NewItem("Nike Shirt")

	// 建立两个观察者
	observerFirst := NewCustomer("abc@gmail.com", "jiang")
	observerSecond := NewCustomer("xyz@gmail.com", "jiangaaaaa")

	// 注册一下
	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	// 通知所有
	shirtItem.UpdateAllObserver()
	// 通知特定
	shirtItem.UpdateSingleObserver("abc@gmail.com")

}
