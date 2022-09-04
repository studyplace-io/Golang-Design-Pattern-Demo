package WrapperAdapterPattern

import "fmt"

// 对象二
type Windows struct{}

// 这个
func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
