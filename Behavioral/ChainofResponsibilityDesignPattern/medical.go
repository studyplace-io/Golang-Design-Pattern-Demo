package ChainofResponsibilityDesignPattern

import "fmt"

type Medical struct {
	// 可以有自己业务逻辑的对象

	// 需要关连下一个部门
	next Department
}


// 执行
func (m *Medical) Execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.Execute(p)
}

// 关连到下一个
func (m *Medical) SetNext(next Department) {
	m.next = next
}
