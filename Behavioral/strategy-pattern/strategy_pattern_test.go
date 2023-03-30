package strategy_pattern

import (
	"fmt"
	"testing"
)

func TestStrategyPattern(t *testing.T) {
	// 依序 Travel(对象) 经过：  Beijing -> Shanghai -> Guangzhou

	// 不同策略的交通工具
	v := CreateTransportation(VehicleType)()
	r := CreateTransportation(HighSpeedRailType)()
	a := CreateTransportation(AirplaneType)()

	// 一个行程流程。
	travel := NewTravel(Beijing, 100)

	// 选择好目的地后，可以使用责任链模式自由选择流程，注意先后顺序。

	bs := NewBeijingStation()
	ss := NewShanghaiStation()
	gs := NewGuangzhouStation()

	bs.ChooseTransportation(v)
	bs.SetNext(ss)

	ss.ChooseTransportation(a)
	ss.SetNext(gs)

	gs.ChooseTransportation(r)
	gs.SetNext(nil) // 最后一站，后续没有行程

	bs.Execute(travel)
	fmt.Println(travel.travelExpenses)
}
