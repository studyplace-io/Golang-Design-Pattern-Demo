package subject_object_pattern

import "fmt"

// Customer 顾客对象
type Customer struct {
	// id
	id   string
	name string
}

var _ Observer = &Customer{}

func NewCustomer(id string, name string) *Customer {
	return &Customer{id: id, name: name}
}

// update 更新
func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}

func (c *Customer) getName() string {
	return c.name
}