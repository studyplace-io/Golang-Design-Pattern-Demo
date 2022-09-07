package main

import (
	"golanglearning/new_project/for_gong_zhong_hao/Design-Patterns-practice/DecoratorPattern"
	"fmt"
)

func main() {
	pizza := DecoratorPattern.NewVeggeMania(20)

	//Add cheese topping
	pizzaWithCheese := &DecoratorPattern.CheeseTopping{
		Pizza: pizza,
		CheeseToppingPrice: 12,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &DecoratorPattern.TomatoTopping{
		Pizza: pizzaWithCheese,
		TomatoToppingPrice: 10,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}



