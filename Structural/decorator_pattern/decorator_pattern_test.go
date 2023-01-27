package decorator_pattern

import (
	"fmt"
	"testing"
)

func TestDecoratorPattern(t *testing.T) {

	user := GetInfo(0)
	fmt.Println(user)

	user2 := GetInfoWithRole(GetInfo)
	fmt.Println(user2(0))

}
