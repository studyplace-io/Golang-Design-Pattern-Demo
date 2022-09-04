package StrategyPatterns

// 一个接口，每个策略都实现这个接口
type StrategyInterface interface {
	method(c *Cache)
}
