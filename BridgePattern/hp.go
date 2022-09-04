package BridgePattern

import "fmt"

// 对象1
type Hp struct {
}

// 实现接口要实现的方法

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
