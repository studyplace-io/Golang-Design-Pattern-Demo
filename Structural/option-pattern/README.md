### 选项模式介绍：
选项模式适用场景：一个函数会有很多参数，为了方便函数的使用，
被调用方会希望给一些参数设定默认值，调用时只需要传与默认值不同的参数即可。
1. 参数多且复杂，会影响调用方使用。
2. 参数有明确的默认值。
3. 参数的扩展。

- 一句话概括：为复杂对象配置需要的参数。
- 实现方法：
    1. 创建一个XXXOption方法，并把WithXXX方法的输出设置成XXXOption。
- 使用方法：
    1. 调用NewXXX方法时，里面可以放入WithXXX的列表来设置参数。

### 示例：
```go
// ClientOption 选项模式的最重要的步骤。
// 必须要传入一个func，这个func可以操作要"创建的对象"，如此才能改变对象中的配置
type ClientOption func(client *HttpClient)
type ClientOptions []ClientOption

// apply 调用
func (cos ClientOptions) apply(c *HttpClient) {
    for _, opt := range cos {
        opt(c) // 调用
    }
}

// NewHttpClient 创建一个HttpClient对象，默认配置，但开放选项模式让用户自定义配置。
func NewHttpClient(opts ...ClientOption) *HttpClient {
    c := &HttpClient {
        Timeout: defaultTimeout,
        MaxIdle: defaultMaxIdle,
        ErrCallBack: defaultErrCallBack,
    }

    // 法一：直接遍历，并执行ClientOption方法
    //for _, opt := range opts {
    //	opt(c)
    //}

    // 法二：可做成一个切片，再执行apply()方法
    ClientOptions(opts).apply(c)

    return c
}

// WithXXX方法中还能做入参校验，过滤不符合入参的情况

func WithTimeout(timeout int) ClientOption {
    return func(client *HttpClient) {
        client.Timeout = timeout
    }
}

func WithMaxIdle(MaxIdle int) ClientOption {
    return func(client *HttpClient) {
        client.MaxIdle = MaxIdle
    }
}

func WithErrCallback(callback func(error)) ClientOption {
    return func(client *HttpClient) {
        client.ErrCallBack = callback
    }
}

```
