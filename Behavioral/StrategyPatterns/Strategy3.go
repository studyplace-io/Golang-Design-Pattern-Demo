package StrategyPatterns

import "fmt"

type Strategy3 struct {
}

func (s *Strategy3) method(c *Cache) {
	fmt.Println("using the Strategy3 method")
}
