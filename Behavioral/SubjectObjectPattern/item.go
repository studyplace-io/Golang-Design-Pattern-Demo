package SubjectObjectPattern

import "fmt"

// 商品对象
type Item struct {
	// 记录观察者的对象
	observerList []Observer
	// 名称
	name         string
	// 是否在卖
	inStock      bool
}

// 创建对象
func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

// 商品上架的方法
func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	// 改变对象状态
	i.inStock = true
	// 通知观察者
	i.NotifyAll()
}

// 注册到里面
func (i *Item) Register(o Observer) {
	i.observerList = append(i.observerList, o)
}

// 注销方法
func (i *Item) Deregister(o Observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

// 通知观察者。
func (i *Item) NotifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

// 注销操作内的方法
func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
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
