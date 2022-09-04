package BridgePattern

import "fmt"

// 对象1
type Mac struct {
	printer Printer
}

// 实现接口要实现的方法
// 具体

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}
