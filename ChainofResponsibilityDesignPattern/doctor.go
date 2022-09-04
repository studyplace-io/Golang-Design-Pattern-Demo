package ChainofResponsibilityDesignPattern

import "fmt"

type Doctor struct {
	// 可以有自己业务逻辑的对象

	// 需要关连下一个部门
	next Department
}


// 执行
func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.Execute(p)
}

// 关连到下一个
func (d *Doctor) SetNext(next Department) {
	d.next = next
}