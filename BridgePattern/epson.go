package BridgePattern

import "fmt"

// 对象2
type Epson struct {
}

// 实现接口要实现的方法

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
