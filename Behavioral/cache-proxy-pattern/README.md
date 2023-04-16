### 缓存代理模式介绍：
缓存代理主要使用场景：可以存储某些耗时的操作结果，当再次使用时，可以快速调用。

- 一句话概括：当访问到请求成本比较高的服务时，可以考虑使用此方式。
- 实现方法：
    1. 实现抽象接口，并在缓存代理对象中初始化缓存map。
- 使用方法：
    1. 先初始化抽象代理与具体代理。
    2. 调用方法。

### 示例：
```go
// RequestProvider 请求抽象接口
type RequestProvider interface {
    DoRequest(requestId string) string
}

// CacheRequestProvider 缓存代理对象
type CacheRequestProvider struct {
    provider RequestProvider	// 接口对象
    // TODO: 这里可以实现一个LRU的缓存淘汰策略
    cache    map[string]string 	// 缓存map
}

// DoRequest 调用请求: 先从缓存里查，没有就调用具体对象远程请求
func (c *CacheRequestProvider) DoRequest(requestId string) string {
    // 先查，如果没有就调用，且存入缓存
    request, ok := c.cache[requestId]
    if !ok {
        request = c.provider.DoRequest(requestId)
        c.cache[requestId] = request
    }
    
    return request
}
```
