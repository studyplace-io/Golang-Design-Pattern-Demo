package StrategyPatterns

import "fmt"

type Strategy2 struct {
}

func (s *Strategy2) method(c *Cache) {
	fmt.Println("using the Strategy2 method")
}
