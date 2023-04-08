### 迭代器模式介绍：
迭代器模式主要使用场景：顺序访问自定义对象的元素。

- 一句话概括：遍历整合的自定义对象
- 实现方法：
    1. 实现所需要的Iterator接口与Collection接口
- 使用方法：
    1. 实例化对象与迭代器，并把所需要的对象注入迭代器。
    2. 调用迭代器的HasNext() GetNext()方法。

### 示例：
```go
// Iterator 迭代器接口
type Iterator interface {
    // 是否有下一个的方法
    HasNext() bool
    // 拿到下一个的方法
    GetNext() *User
}

// Collection 收集器接口对象：负责注入需要迭代的对象，返回迭代器对象
type Collection interface {
      CreateIterator() Iterator
}

```
