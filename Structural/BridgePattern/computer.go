package BridgePattern

// 抽象接口
type Computer interface {
	// 两个方法
	// 打印方法
	Print()
	// 设置具体实现的方法
	SetPrinter(Printer)
}