package strategy_pattern

import "fmt"

type Travel struct {
	station        DestinationType
	travelExpenses int

	// 用来区分是否执行过此流程
	beijingDone    bool
	shanghaiUpDone bool
	guangzhouDone  bool
}

func NewTravel(station DestinationType, travelExpenses int) *Travel {
	return &Travel{station: station, travelExpenses: travelExpenses}
}

type DestinationType int

const (
	Beijing = iota
	Shanghai
	Guangzhou
)

// Process 流程接口，表示相关但负责不同功能的对象均实现相同方法
type Process interface {
	Execute(*Travel)                                          // 执行该流程的方法。
	ChooseTransportation(strategyInterface StrategyInterface) // 选择交通工具
	SetNext(Process)                                          // 交给下一个流程
}

var _ Process = &BeijingStation{}

// BeijingStation 对象
type BeijingStation struct {
	// 可以有自己业务逻辑的对象
	// 需要关连下一个部门
	strategyInterface StrategyInterface
	next              Process
}

func NewBeijingStation() *BeijingStation {
	return &BeijingStation{}
}

func (r *BeijingStation) ChooseTransportation(strategyInterface StrategyInterface) {
	r.strategyInterface = strategyInterface
}

// Execute 执行
func (r *BeijingStation) Execute(t *Travel) {
	// 需要判断一下是否执行过
	if t.beijingDone {
		fmt.Println("Beijing travel already done")
		r.next.Execute(t)
		return
	}
	r.strategyInterface.method(t, Beijing) // 出发去了～
	fmt.Println("go to Beijing!!")
	t.beijingDone = true
	fmt.Println("have fun in Beijing")
	// 为了判断是否为最后一个流程
	if r.next == nil {
		return
	}
	r.next.Execute(t) // 执行下个流程
}

// SetNext 关连到下一个
func (r *BeijingStation) SetNext(next Process) {
	r.next = next
}

var _ Process = &ShanghaiStation{}

// ShanghaiStation 对象
type ShanghaiStation struct {
	// 可以有自己业务逻辑的对象
	strategyInterface StrategyInterface
	next              Process
}

func NewShanghaiStation() *ShanghaiStation {
	return &ShanghaiStation{}
}

func (r *ShanghaiStation) ChooseTransportation(strategyInterface StrategyInterface) {
	r.strategyInterface = strategyInterface
}

// Execute 执行
func (r *ShanghaiStation) Execute(t *Travel) {
	// 需要判断一下是否执行过
	if t.shanghaiUpDone {
		fmt.Println("Shanghai travel already done")
		r.next.Execute(t)
		return
	}
	r.strategyInterface.method(t, Shanghai) // 出发去了～
	fmt.Println("go to Shanghai!!")

	t.shanghaiUpDone = true
	fmt.Println("have fun in Shanghai")
	if r.next == nil {
		return
	}
	r.next.Execute(t) // 执行下个流程
}

// SetNext 关连到下一个
func (r *ShanghaiStation) SetNext(next Process) {
	r.next = next
}

var _ Process = &GuangzhouStation{}

// GuangzhouStation 对象
type GuangzhouStation struct {
	// 可以有自己业务逻辑的对象
	// 需要关连下一个部门
	strategyInterface StrategyInterface
	next              Process
}

func NewGuangzhouStation() *GuangzhouStation {
	return &GuangzhouStation{}
}

func (r *GuangzhouStation) ChooseTransportation(strategyInterface StrategyInterface) {
	r.strategyInterface = strategyInterface
}

// Execute 执行
func (r *GuangzhouStation) Execute(t *Travel) {
	// 需要判断一下是否执行过
	if t.guangzhouDone {
		fmt.Println("Guangzhou travel already done")
		r.next.Execute(t)
		return
	}
	r.strategyInterface.method(t, Guangzhou) // 出发去了～
	fmt.Println("go to Guangzhou!!")

	t.guangzhouDone = true
	fmt.Println("have fun in Guangzhou")
	if r.next == nil {
		return
	}
	r.next.Execute(t) // 执行下个流程
}

// SetNext 关连到下一个
func (r *GuangzhouStation) SetNext(next Process) {
	r.next = next
}
