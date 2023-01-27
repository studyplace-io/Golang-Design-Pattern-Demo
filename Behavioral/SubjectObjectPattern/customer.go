package SubjectObjectPattern

import "fmt"

// 顾客对象
type Customer struct {
	// id
	Id string
}

// 更新
func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.Id, itemName)
}

// 拿到状态
func (c *Customer) getID() string {
	return c.Id
}
