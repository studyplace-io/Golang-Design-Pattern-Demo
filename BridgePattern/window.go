package BridgePattern

import "fmt"

// 对象2
type Windows struct {
	printer Printer
}

// 实现接口要实现的方法
// 具体

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
