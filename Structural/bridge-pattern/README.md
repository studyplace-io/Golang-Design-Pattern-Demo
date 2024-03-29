### 桥接模式介绍：
桥接模式主要使用场景：可以把"抽象"与"具体实现"解耦，使得二者可以独立扩展。

- 一句话概括：通常在项目初期时，可以考虑解耦"抽象"与"实现"，方便后续扩展。
- 实现方法：
    1. 定义两种接口，一个"抽象"一个"具体实现"。在"抽象"接口中，设置有**变换具体实现**的方法。
- 使用方法：
    1. 使用工厂方法初始化对象。
    2. 在"抽象"中可以设置不同的"具体实现"。

![](https://github.com/StudyPlace-io/Golang-Design-Pattern-Demo/blob/main/image/bridge_pattern.jpg?raw=true)
### 示例：
```go
1. 桥接模式+简易工厂方法使用
// Computer 接口
type Computer interface {
    // 其他需要实现的方法
    DoSomething()
    
  // 与桥接模式相关的实现。
    Print()  			// 抽象方法，此方法要去调用 内部"具体实现"的方法(ex: PrintFile())
    SetPrinter(Printer)	// 重点：需要一个修改或设置"具体实现"的方法，传入一个接口对象
}


// Printer "具体实现"接口
// 需要实现PrintFile 方法
type Printer interface {
    PrintFile()
}
```
