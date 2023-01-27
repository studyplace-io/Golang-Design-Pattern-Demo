package DecoratorPattern

type TomatoTopping struct {
	Pizza              IPizza
	TomatoToppingPrice int
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice() + c.TomatoToppingPrice
	return pizzaPrice
}
