package subject_object_pattern

import "fmt"

type Subject interface {
	Notify()
	Add()
	Delete()
}


type Observer interface {
	Update(id int)
}

// 被观察者

type Order struct {
	Id int
	Observers []Observer // 观察者
}

// Notify 被观察者要通知观察者
func (o *Order) Notify() {
	for _, observer := range o.Observers {
		observer.Update(o.Id)
	}
}

// 加入观察者中
func (o *Order) Add(observer Observer) {
	o.Observers = append(o.Observers, observer)
}

func (o *Order) Cancel() {
	fmt.Printf("订单号%d被取消\n", o.Id)
	o.Notify()
}



// 两个观察者

type Stock struct {}

func (s *Stock) Update(id int) {
	fmt.Printf("订单号%d的订单库存被返还\n", id)

}

type Issue struct {}

func (s *Issue) Update(id int) {
	fmt.Printf("订单号%d的工单库存被关闭\n", id)

}