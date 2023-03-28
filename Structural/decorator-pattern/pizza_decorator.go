package decorator_pattern

type CheeseTopping struct {
	Pizza              IPizza
	CheeseToppingPrice int
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice() + c.CheeseToppingPrice
	return pizzaPrice
}

type TomatoTopping struct {
	Pizza              IPizza
	TomatoToppingPrice int
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice() + c.TomatoToppingPrice
	return pizzaPrice
}
