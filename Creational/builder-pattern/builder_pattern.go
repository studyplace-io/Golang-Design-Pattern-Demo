package builder_pattern

// BookBuilder 创建者
type BookBuilder struct {
	Id       int
	BookName string
	Author   string
	Content  string
	Theme    string
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

func (b *BookBuilder) SetAuthor(author string) *BookBuilder {
	b.Author = author
	return b
}

func (b *BookBuilder) SetContent(content string) *BookBuilder {
	b.Content = content
	return b
}

func (b *BookBuilder) SetTheme(theme string) *BookBuilder {
	b.Theme = theme
	return b
}

// 构造函数最后一定要加上Build方法，才能返回struct实体
func (b *BookBuilder) Build() *Book {
	// 实例化
	book := &Book{
		id:       b.Id,
		bookName: b.BookName,
		author: b.Author,
		content: b.Content,
		theme: b.Theme,
	}

	// 这里可以抽象出check()方法，做入参校验
	if b.Price > 0 {
		book.price = b.Price
	}

	return book
}

// DefaultCreate 如果建造的对象太复杂，可让用户只选择必选项的方法
func (b *BookBuilder) DefaultCreate() *Book {
	return b.SetAuthor("default").SetTheme("default").SetContent("default").Build()
}