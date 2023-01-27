package template_pattern

import (
	"fmt"
	"testing"
)

func TestTemplatePattern(t *testing.T) {

	cache := NewRedisCache()
	fmt.Println(cache.Output(123))

	cache2 := NewEtcdCache()
	fmt.Println(cache2.Output(123))

}
