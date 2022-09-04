package main

import "golanglearning/new_project/for-gong-zhong-hao/Design-Patterns-practice/WrapperAdapterPattern"

func main() {

	client := &WrapperAdapterPattern.Client{}
	mac := &WrapperAdapterPattern.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	// windowsMachine并没有实现 computer接口
	windowsMachine := &WrapperAdapterPattern.Windows{}
	// windowsMachineAdapter才实现 computer接口
	windowsMachineAdapter := &WrapperAdapterPattern.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
