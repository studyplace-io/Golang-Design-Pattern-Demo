package chain_of_responsibility_pattern

import "fmt"

// Process 流程接口，表示相关但负责不同功能的对象均实现相同方法
type Process interface {
	Execute(*Patient) // 执行该流程的方法。
	SetNext(Process)  // 交给下一个流程
}

// Reception 对象
type Reception struct {
	// 可以有自己业务逻辑的对象
	// 需要关连下一个部门
	next Process
}

// Execute 执行
func (r *Reception) Execute(p *Patient) {
	// 需要判断一下是否执行过
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.SetCost(p.GetCost() + 50) // 挂号费
	p.registrationDone = true
	r.next.Execute(p) // 执行下个流程
}

// SetNext 关连到下一个
func (r *Reception) SetNext(next Process) {
	r.next = next
}

type Doctor struct {
	// 可以有自己业务逻辑的对象

	// 需要关连下一个部门
	next Process
}

// Execute 执行
func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone || p.illness != "" {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.SetIllness("test-得病了")
	p.SetCost(p.GetCost() + 200) // 要加治病钱
	p.doctorCheckUpDone = true
	d.next.Execute(p)
}

// SetNext 关连到下一个
func (d *Doctor) SetNext(next Process) {
	d.next = next
}

type Medical struct {
	// 可以有自己业务逻辑的对象

	// 需要关连下一个部门
	next Process
}

// Execute 执行
func (m *Medical) Execute(p *Patient) {
	if p.medicineDone || p.prescriptions != "" {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.SetPrescriptions("test-配到药了")
	p.SetCost(p.GetCost() + 150) // 要加治病钱
	p.medicineDone = true
	m.next.Execute(p)
}

// SetNext 关连到下一个
func (m *Medical) SetNext(next Process) {
	m.next = next
}

type Cashier struct {
	// 可以有自己业务逻辑的对象
	// 需要关连下一个部门
	next Process
}

// Execute 执行
// 最后阶段
func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
	fmt.Println("总共需要多少钱: ", p.GetCost())
	fmt.Println("病症: ", p.GetIllness())
	fmt.Println("处方: ", p.GetPrescriptions())
}

// SetNext 关连到下一个，最后一个流程虽然有，但没有调用
func (c *Cashier) SetNext(next Process) {
	c.next = next
}
