package wrapper_adapter_pattern

import (
	"fmt"
	"testing"
)

func TestWrapperAdapterPattern(t *testing.T) {
	client := NewClient()
	mac := &Mac{brand: "mac"}

	client.InsertLightningConnectorIntoComputer(mac)

	// windowsMachine并没有实现 computer接口
	windowsMachine := &Windows{brand: "windows"}
	// windowsMachineAdapter才实现 computer接口
	windowsMachineAdapter := &WindowsAdapter{
		*windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

	windowsMachineAdapter.DoSomething()
	fmt.Println(windowsMachineAdapter.GetBrand())
}