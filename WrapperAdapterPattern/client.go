package WrapperAdapterPattern


import "fmt"

// 对象
type Client struct {
}

// 调用 传入接口对象
func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
