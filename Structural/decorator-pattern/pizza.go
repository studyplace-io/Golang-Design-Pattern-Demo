package decorator_pattern


// IPizza 披萨对象接口
type IPizza interface {
	GetPrice() int
}

type Pizza struct {
	pizzaprice int
}

func NewPizza(price int) *Pizza {
	return &Pizza{
		pizzaprice: price,
	}
}

func (p *Pizza) GetPrice() int {
	return p.pizzaprice
}
