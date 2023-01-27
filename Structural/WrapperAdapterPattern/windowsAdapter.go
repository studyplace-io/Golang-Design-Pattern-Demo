package WrapperAdapterPattern

import "fmt"

// Windows适配器对象
type WindowsAdapter struct {
	WindowMachine *Windows
}

//
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	// 这里可以进行对象适配。 不需要直接修改代码，因为有时候代码不是自己维护的
	// 或是用的是第三方包 不能修改。
	w.WindowMachine.insertIntoUSBPort()
}
