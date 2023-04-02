### 观察者模式介绍：
观察者模式主要使用场景：定义一种订阅机制，可在对象事件发生时通知多个"观察"该对象的其他对象。

- 一句话概括：简易订阅模式，分为观察者与被观察者。
- 实现方法：
    1. 分别定义Subject接口与Object接口，并让对象分别实现该接口。
- 使用方法：
    1. 创建观察者与被观察者，并且把"观察者"加入"被观察者"维护的切片。
    2. 当"被观察者"的状态改变时，使用通知方法来通知观察者。

### 示例：
```go

// Subject 被观察者对象要实现的接口
type Subject interface {
    // Register 注册方法
    Register(observer Observer)
    // Deregister 注销方法
    Deregister(observer Observer)
    // NotifyAll 通知所有观察者方法
    notifyAll()
    // NotifySingle 通知单一观察者方法
    notifySingle(id string)
}

// Observer 观察者的对象要实现的接口
type Observer interface {
    // 更新方法
    update(string)
    // 取得id方法
    getID() string
}



```
