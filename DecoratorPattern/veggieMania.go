package DecoratorPattern


type VeggeMania struct {
	pizzaprice	int
}

func NewVeggeMania(price int) *VeggeMania {
	return &VeggeMania{
		pizzaprice: price,
	}
}

func (p *VeggeMania) GetPrice() int {
	return p.pizzaprice
}