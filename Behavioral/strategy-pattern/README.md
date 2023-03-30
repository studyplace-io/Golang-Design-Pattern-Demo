### 策略模式介绍：
策略模式主要使用场景：当有需要执行同一动作，但使用不同方式的时候，可以使用。

- 一句话概括：将一组同性质的行为转换成对象， 并让它在对象内部能够相互替换。(有点类似接口的感觉)
- 实现方法：
    1. 不同的策略都实现同一接口，并且使它们可相互替换。
- 使用方法：
    1. 定义不同策略对象都实现此接口

### 示例：
```go
1. 策略模式+责任链模式+简易工厂模式实现
// StrategyInterface 一个接口，每个策略都实现这个接口
// 去往目的地的"不同策略"方法
type StrategyInterface interface {
    method(c *Travel, place DestinationType)
}
// 不同策略对象都实现此接口
var _ StrategyInterface = &Vehicle{}

var _ StrategyInterface = &HighSpeedRail{}

var _ StrategyInterface = &Airplane{}
```
