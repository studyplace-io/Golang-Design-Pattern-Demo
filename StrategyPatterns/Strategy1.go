package StrategyPatterns

import "fmt"

type Strategy1 struct {
}

func (s *Strategy1) method(c *Cache) {
	fmt.Println("using the Strategy1 method")
}
