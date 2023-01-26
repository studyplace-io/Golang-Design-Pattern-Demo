package builder_pattern

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T) {

	id := 100
	name := "test-builder-pattern"
	// 建造者模式，最后一定要调用Build()方法
	book := new(Book).Builder(id, name).SetPrice(10000.0).Build()

	fmt.Println(book)


}
