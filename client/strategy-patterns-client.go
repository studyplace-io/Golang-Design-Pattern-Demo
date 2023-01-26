package main

import "golanglearning/new_project/for-gong-zhong-hao/Design-Patterns-practice/StrategyPatterns"

func main() {
	Strategy1 := &StrategyPatterns.Strategy1{}
	cache := StrategyPatterns.InitCache(Strategy1)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	Strategy2 := &StrategyPatterns.Strategy2{}
	cache.SetStrategyInterface(Strategy2)

	cache.Add("d", "4")

	Strategy3 := &StrategyPatterns.Strategy3{}
	cache.SetStrategyInterface(Strategy3)

	cache.Add("e", "5")

}
