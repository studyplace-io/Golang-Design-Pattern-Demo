package bridge_pattern

import "fmt"

// Epson 对象
type Epson struct {
}

func NewEpson() PrinterCreateFunc {
	return func() interface{} {
		return &Epson{}
	}
}

var _ Printer = &Epson{}

// 实现接口要实现的方法

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// Hp 对象
type Hp struct {
}

func NewHp() PrinterCreateFunc {
	return func() interface{} {
		return &Hp{}
	}
}

var _ Printer = &Hp{}

// 实现接口要实现的方法

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
