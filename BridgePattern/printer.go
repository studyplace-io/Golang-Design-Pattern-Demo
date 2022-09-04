package BridgePattern

// 抽象接口
type Printer interface {
	// 方法
	// 因为Printer是被替换的对象，所以不用实现set方法。
	PrintFile()
}
