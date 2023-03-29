package wrapper_adapter_pattern

import "fmt"

// Computer 接口对象
type Computer interface {
	// DoSomething 一般的方法
	DoSomething()
	GetBrand() string

	InsertIntoLightningPort()
}

var _ Computer = &Mac{}

// Mac 对象
type Mac struct {
	brand string
}

func (m *Mac) DoSomething() {
	fmt.Println("do something with computer.")
}

func (m *Mac) GetBrand() string {
	return m.brand
}

// InsertIntoLightningPort 这个对象的方法 直接适配，不需要修改
func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

/*
 windows 对象并没有实现InsertIntoLightningPort()方法 (故意不实现的)
*/

//var _ Computer = &Windows{} 实现不了

// Windows 对象
type Windows struct {
	brand string
}

func (w *Windows) DoSomething() {
	fmt.Println("do something with computer.")
}

func (w *Windows) GetBrand() string {
	return w.brand
}

// insertIntoUSBPort
func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}