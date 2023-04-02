package builder_pattern

type Book struct {
	id       int
	bookName string
	author   string
	content  string
	theme    string
	price    float64
}

func (b *Book) Builder(id int, name string) *BookBuilder {
	return NewBookBuilder(id, name)
}

func (b *Book) GetInfo() {

}
