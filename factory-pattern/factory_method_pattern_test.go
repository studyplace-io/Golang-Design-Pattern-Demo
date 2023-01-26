package factory_pattern

import (
	"fmt"
	"testing"
)

func TestFactoryMethodPattern(t *testing.T) {


	fmt.Println(new(TechFactory).CreateProduct(ProductTechBook).GetInfo())

	fmt.Println(new(DailyFactory).CreateProduct(ProductDailyCar).GetInfo())

}
