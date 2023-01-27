package ChainofResponsibilityDesignPattern

import "fmt"

// 对象
type Reception struct {
	// 可以有自己业务逻辑的对象

	// 需要关连下一个部门
	next Department
}

// 执行
func (r *Reception) Execute(p *Patient) {
	// 需要判断一下是否执行过
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.Execute(p)
}

// 关连到下一个
func (r *Reception) SetNext(next Department) {
	r.next = next
}