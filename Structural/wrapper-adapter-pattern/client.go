package wrapper_adapter_pattern

import "fmt"

// Client 对象
type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

// InsertLightningConnectorIntoComputer 调用 传入接口对象
func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
