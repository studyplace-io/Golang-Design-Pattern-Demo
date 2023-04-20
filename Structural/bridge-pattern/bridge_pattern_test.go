package bridge_pattern

import (
	"fmt"
	"testing"
)

func TestBridgePattern(t *testing.T) {

	// 创建打印机
	tt := NewPrinter(Hps)()
	hpPrinter := tt.(*Hp)

	ee := NewPrinter(Epsons)()
	epsonPrinter := ee.(*Epson)

	// mac电脑
	a := NewComputer(Macs)()
	macComputer := a.(*Mac)
	// mac电脑可以随时替换 两种不同的打印机

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	macComputer.DoSomething()
	fmt.Println()

	// window电脑
	b := NewComputer(Wins)()
	winComputer := b.(*Windows)

	// window电脑也可以随时替换 不同的打印机
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	winComputer.DoSomething()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
