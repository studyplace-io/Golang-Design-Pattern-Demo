package bridge_pattern

// Computer 接口
type Computer interface {
	// 其他需要实现的方法
	DoSomething()

	// 与桥接模式相关的实现。
	Print()  				// 抽象方法，此方法要去调用 内部"具体实现"的方法(ex: PrintFile())
	SetPrinter(Printer)	    // 重点：需要一个修改或设置"具体实现"的方法，传入一个接口对象
}

const (
	Wins = iota
	Macs
)

type ComputeType int

type ComputerCreateFunc func() interface{}

func NewComputer(t ComputeType) ComputerCreateFunc {
	switch t {
	case Wins:
		return NewWindows()
	case Macs:
		return NewMac()
	}

	return nil
}
