package factory_pattern

/*
	工厂方法：可以有多个工厂
 */

type Book struct {
	Id int
	BookName string
}

func (b *Book) GetInfo() string {
	return "book"
}

type Car struct {
	Id int
	CarName string
}

func (c *Car) GetInfo() string {
	return "car"
}
