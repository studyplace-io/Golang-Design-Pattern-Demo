package WrapperAdapterPattern

import "fmt"

// 对象一
type Mac struct {
}

// 这个对象的方法 直接适配，不需要修改
func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}