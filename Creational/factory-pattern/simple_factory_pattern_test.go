package factory_pattern

import (
	"fmt"
	"testing"
)

func TestFactoryPattern(t *testing.T) {

	//admin := NewAdmin()
	//user := NewUser()

	user := CreateUser(NormalUser)(111, "test")
	fmt.Println(user)


}
