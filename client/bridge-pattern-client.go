package main

import (
	"fmt"
	"golanglearning/new_project/for-gong-zhong-hao/Design-Patterns-practice/Structural/BridgePattern"
)

func main() {

	// 打印机
	hpPrinter := &BridgePattern.Hp{}
	epsonPrinter := &BridgePattern.Epson{}

	// mac电脑
	macComputer := &BridgePattern.Mac{}

	// mac电脑可以随时替换 两种不同的打印机

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	// window电脑
	winComputer := &BridgePattern.Windows{}

	// window电脑也可以随时替换 不同的打印机
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
