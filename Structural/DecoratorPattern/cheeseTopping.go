package DecoratorPattern


type CheeseTopping struct {
	Pizza              IPizza
	CheeseToppingPrice int
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice() + c.CheeseToppingPrice
	return pizzaPrice
}