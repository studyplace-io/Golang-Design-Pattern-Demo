package bridge_pattern

import "fmt"

// Windows 对象
type Windows struct {
	printer Printer
}

func NewWindows() ComputerCreateFunc {
	return func() interface{} {
		return &Windows{}
	}
}

var _ Computer = &Windows{}

// 实现接口要实现的方法
// 具体

func (w *Windows) DoSomething() {
	fmt.Println("Do something with windows")
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// Mac 对象1
type Mac struct {
	printer Printer
}

func NewMac() ComputerCreateFunc {
	return func() interface{} {
		return &Mac{}
	}
}

var _ Computer = &Mac{}

// 实现接口要实现的方法
// 具体

func (m *Mac) DoSomething() {
	fmt.Println("Do something with mac")
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}
