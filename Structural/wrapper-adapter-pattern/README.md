### 适配器模式介绍：
适配器模式主要使用场景：当旧有对象需要实现新的接口时使用。(解决兼容问题)

- 一句话概括：对象变动太复杂，需要急救火的时候使用。
- 实现方法：
    1. 适配器模式将一个类的接口，转换成被调用方期望的另一个接口。
- 使用方法：
    1. 创建一个xxxAdapter对象，并使用嵌套struct。
    2. 实现新接口中需要的方法。

### 示例：
```go
/*
 windows 对象并没有实现InsertIntoLightningPort()方法 (故意不实现的)
*/

//var _ Computer = &Windows{} 实现不了

// Windows 对象
type Windows struct {
    brand string
}

func (w *Windows) DoSomething() {
    fmt.Println("do something with computer.")
}

func (w *Windows) GetBrand() string {
    return w.brand
}

// insertIntoUSBPort
func (w *Windows) insertIntoUSBPort() {
    fmt.Println("USB connector is plugged into windows machine.")
}


// WindowsAdapter 适配器对象
type WindowsAdapter struct {
    Windows // 使用嵌套struct，实现剩下不好实现的接口方法。
}

// WindowsAdapter 才能实现Computer接口
var _ Computer = &WindowsAdapter{}

// InsertIntoLightningPort 适配器模式的重点，另外用一个对象，采取"嵌套"的方式，并增加本来需要的方法。
func (w *WindowsAdapter) InsertIntoLightningPort() {
    fmt.Println("Adapter converts Lightning signal to USB.")
    // 这里可以进行对象适配。 不需要直接修改代码，因为有时候代码不是自己维护的
    // 或是用的是第三方包 不能修改。
    w.insertIntoUSBPort()
}
```
