package builder_pattern

type Book struct {
	Id 			int
	BookName 	string
	Price 		float64
}

func (b *Book) Builder(id int, name string) *BookBuilder {
	return NewBookBuilder(id, name)
}

func (b *Book) GetInfo() {

}
