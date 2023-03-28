package decorator_pattern


import "fmt"

// 应该把功能抽象承接口

// ICar 车子对象接口
type ICar interface {
	Parking() string
	Driving() string
}

var _ ICar = &Car{}

type Car struct {
}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) Parking() string {
	return fmt.Sprintf("Parking!!")
}

func (c *Car) Driving() string {
	return fmt.Sprintf("Driving!!")
}