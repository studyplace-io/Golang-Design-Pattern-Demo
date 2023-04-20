package factory_pattern

/*
	工厂方法：可以有多个工厂
*/

type Book struct {
	Id       int
	BookName string
	Info     string
}

func (b *Book) GetInfo() string {
	return b.Info
}

func (b *Book) SetInfo() {
	b.Info = "这是一本书"
}

type Car struct {
	Id      int
	CarName string
	Info    string
}

func (c *Car) GetInfo() string {
	return c.Info
}

func (c *Car) SetInfo() {
	c.Info = "这是一辆车"
}
