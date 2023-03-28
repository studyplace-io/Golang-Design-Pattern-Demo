package decorator_pattern

import (
	"fmt"
	"testing"
)

func TestPizzaDecorator(t *testing.T) {
	pizza := NewPizza(20)

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		Pizza:              pizza,
		CheeseToppingPrice: 12,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		Pizza:              pizzaWithCheese,
		TomatoToppingPrice: 10,
	}

	fmt.Printf("Price of pizza with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}

func TestCarDecorator(t *testing.T) {

	// 普通的车
	car := NewCar()
	fmt.Println("car: ", car.Driving(), car.Parking())

	// 救护车
	ambulanceCar := NewAmbulance(car, 5)
	ambulanceCar.SavePeople()
	fmt.Println("ambulanceCar: ", ambulanceCar.Driving(), ambulanceCar.Parking())

	// 战车
	chariotCar := NewChariot(car, "missile", 5)
	chariotCar.AttackPeople()
	fmt.Println("chariotCar: ", chariotCar.Driving(), chariotCar.Parking())

	// 既有救人功能也有战斗功能的车
	chariotCar1 := NewChariot(ambulanceCar, "missile", 5)
	chariotCar1.AttackPeople()
	//chariotCar1.SavePeople()

}
