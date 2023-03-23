package builder_pattern

type BookBuilder struct {
	Id       int
	BookName string
	Price    float64
}

func NewBookBuilder(id int, name string) *BookBuilder {
	return &BookBuilder{
		Id:       id,
		BookName: name,
	}
}

func (b *BookBuilder) SetPrice(price float64) *BookBuilder {
	b.Price = price
	return b
}

// 构造函数最后一定要加上Build方法，才能返回struct实体
func (b *BookBuilder) Build() *Book {
	// 实例化
	book := &Book{
		id:       b.Id,
		bookName: b.BookName,
	}

	// 这里可以抽象出check()方法，做入参校验
	if b.Price > 0 {
		book.price = b.Price
	}

	return book
}