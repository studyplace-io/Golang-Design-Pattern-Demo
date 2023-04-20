package wrapper_adapter_pattern

import "fmt"

// WindowsAdapter 适配器对象
type WindowsAdapter struct {
	Windows // 使用嵌套struct，实现剩下不好实现的接口方法。
}

// WindowsAdapter 才能实现Computer接口
var _ Computer = &WindowsAdapter{}

// InsertIntoLightningPort 适配器模式的重点，另外用一个对象，采取"嵌套"的方式，并增加本来需要的方法。
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	// 这里可以进行对象适配。 不需要直接修改代码，因为有时候代码不是自己维护的
	// 或是用的是第三方包 不能修改。
	w.insertIntoUSBPort()
}
