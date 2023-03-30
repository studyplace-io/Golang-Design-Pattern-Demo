package strategy_pattern

import "fmt"

// StrategyInterface 一个接口，每个策略都实现这个接口
// 去往目的地的"不同策略"方法
type StrategyInterface interface {
	method(c *Travel, place DestinationType)
}

var _ StrategyInterface = &Vehicle{}

type Vehicle struct {
}

func NewVehicle() TransportationCreateFunc {
	return func() StrategyInterface {
		return &Vehicle{}
	}
}

func (s *Vehicle) method(c *Travel, place DestinationType) {
	c.station = place
	c.travelExpenses += 100
	fmt.Println("using the Vehicle method")
}

var _ StrategyInterface = &HighSpeedRail{}

type HighSpeedRail struct {
}

func NewHighSpeedRail() TransportationCreateFunc {
	return func() StrategyInterface {
		return &HighSpeedRail{}
	}
}

func (s *HighSpeedRail) method(c *Travel, place DestinationType) {
	c.station = place
	c.travelExpenses += 500
	fmt.Println("using the HighSpeedRail method")
}

var _ StrategyInterface = &Airplane{}

type Airplane struct {
}

func NewAirplane() TransportationCreateFunc {
	return func() StrategyInterface {
		return &Airplane{}
	}
}

func (s *Airplane) method(c *Travel, place DestinationType) {
	c.station = place
	c.travelExpenses += 600
	fmt.Printf("using the Airplane method")
}

const (
	VehicleType = iota
	HighSpeedRailType
	AirplaneType
)

type Transportation int

// TransportationCreateFunc 工厂方法要返回的func
type TransportationCreateFunc func() StrategyInterface

// CreateTransportation 简易工厂初始化不同策略对象
func CreateTransportation(t Transportation) TransportationCreateFunc {
	switch t {
	case VehicleType:
		return NewVehicle()
	case HighSpeedRailType:
		return NewHighSpeedRail()
	case AirplaneType:
		return NewAirplane()
	}
	return nil
}
