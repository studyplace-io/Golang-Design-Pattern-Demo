package subject_object_pattern

import "fmt"

// Item 商品对象
type Item struct {
	// 记录观察者的对象，观察者模式的重点
	observerList []Observer
	// 名称
	name string
	// 是否在卖
	inStock bool
}

var _ Subject = &Item{}

// NewItem 创建对象
func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

// UpdateAllObserver 商品上架的方法，通知所有观察者方法
func (i *Item) UpdateAllObserver() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	// 改变对象状态
	i.inStock = true
	// 通知观察者
	i.notifyAll()
}

// UpdateSingleObserver 商品上架的方法，通知特定观察者方法
func (i *Item) UpdateSingleObserver(id string) {
	fmt.Printf("Item %s is now in stock\n", i.name)
	// 改变对象状态
	i.inStock = true
	// 通知观察者
	i.notifySingle(id)
}

// Register 注册到被观察者的名单中
func (i *Item) Register(o Observer) {
	i.observerList = append(i.observerList, o)
}

// Deregister 注销方法
func (i *Item) Deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

// NotifyAll 通知观察者
func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

// NotifySingle 通知单一观察者
func (i *Item) notifySingle(id string) {
	for _, observer := range i.observerList {
		if observer.getID() == id {
			observer.update(id)
			break
		}
	}
}

// removeFromSlice 注销操作内的方法
func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		// match到的id相同。
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
