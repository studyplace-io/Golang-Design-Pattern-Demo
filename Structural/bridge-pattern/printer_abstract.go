package bridge_pattern

// Printer "具体实现"接口
// 需要实现PrintFile 方法
type Printer interface {
	PrintFile()
}

const (
	Hps = iota
	Epsons
)

type PrinterType int

type PrinterCreateFunc func() interface{}

func NewPrinter(t ComputeType) PrinterCreateFunc {
	switch t {
	case Hps:
		return NewHp()
	case Epsons:
		return NewEpson()
	}

	return nil
}
