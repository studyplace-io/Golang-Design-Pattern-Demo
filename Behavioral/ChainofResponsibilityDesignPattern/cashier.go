package ChainofResponsibilityDesignPattern

import "fmt"

type Cashier struct {
	// 可以有自己业务逻辑的对象

	// 需要关连下一个部门
	next Department
}

// 执行
// 最后阶段。
func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

// 关连到下一个
func (c *Cashier) SetNext(next Department) {
	c.next = next
}