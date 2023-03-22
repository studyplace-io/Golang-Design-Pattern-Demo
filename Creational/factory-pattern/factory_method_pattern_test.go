package factory_pattern

import (
	"fmt"
	"testing"
)

func TestFactoryMethodPattern(t *testing.T) {

	var test IProductFactory
	test = &StudyFactory{}

	book := test.CreateProduct(ProductStudyBook)
	book.SetInfo()
	fmt.Println(book.GetInfo())

	test = &MachineryFactory{}
	car := test.CreateProduct(ProductMachineryCar)
	car.SetInfo()
	fmt.Println(car.GetInfo())

}
