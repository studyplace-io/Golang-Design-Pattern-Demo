### 代理模式介绍：
代理模式主要使用场景：业务系统的非功能性需求开发：如监控、统计、鉴权、限流、事务、幂等、日志等，这些和业务没有关系，所以可以放到Proxy模式中扩展。

- 一句话概括：执行一个任务或是一个请求时，经过主对象前，先经过Proxy对象，proxy对象背后再去调用
- 实现方法：主要分为两种：
    1. 定义一个Proxy对象，并实现"原对象"的关键方法(通常会把原对象抽象成接口，并原对象与Proxy对象均实现此接口)。
- 使用方法：
    1. 定义一个代理对象Proxy，并嵌入原对象，并定义需要的配置参数(使用选项模式)。
    2. 重点：Proxy对象"必须"实现与原对象相同的方法(接口)

### 示例：
```go
1. 代理模式+选项模式
// NginxProxy 代理类对象
type NginxProxy struct {
    // 真正的对象访问的对象
    application *Application
    // 最大限流量
    maxAllowedRequest int
    // 当前限流量
    rateLimiter map[string]int
}

const (
    defaultName              = "default-application"
    defaultOther             = "default-other"
    defaultMaxAllowedRequest = 10
)

// NewNginxProxyServer 创建新的对象，加入选项模式
func NewNginxProxyServer(opts ...NginxProxyOption) *NginxProxy {

      instance := &NginxProxy{
          application: &Application{
          name:  defaultName,
          other: defaultOther,
          },
          maxAllowedRequest: defaultMaxAllowedRequest,
          rateLimiter:       make(map[string]int),
      }
      
      for _, opt := range opts {
          opt(instance)
      }
    
      return instance
}

type NginxProxyOption func(proxy *NginxProxy)

func WithMaxAllowedRequest(maxAllowedRequest int) NginxProxyOption {
    return func(proxy *NginxProxy) {
            proxy.maxAllowedRequest = maxAllowedRequest
    }
}

func WithName(name string) NginxProxyOption {
    return func(proxy *NginxProxy) {
    	proxy.application.name = name
    }
}

func WithOther(other string) NginxProxyOption {
    return func(proxy *NginxProxy) {
            proxy.application.other = other
    }
}

// HandleRequest 调用代理类的HandleRequest方法
// 最重要的地方：需要实现与原来对象一样的方法。
func (n *NginxProxy) HandleRequest(url, method string) (int, string) {
    // 做一下限流功能
    allowed := n.checkRateLimiting(url)
    if !allowed {
        return 403, "Not Allowed"
    }
    // 如果成功，就进行真正的application.handleRequest(url, method)！
    return n.application.HandleRequest(url, method)
}

2. 代理模式搭配装饰器模式
// 代理对象
type UserProxy struct {
    svc *UserService
}

func NewUserProxy(svc *UserService) *UserProxy {
    return &UserProxy{
      svc: svc,
    }
}

func (u *UserProxy) Login(decorator LogDecorator) LoginFunc {
    fmt.Println("记录日志！！")
    // 使用代理模式抽象出来
    // 使用装饰器模式调用f方法
    return decorator(u.svc.Login)
}


// 结合装饰器模式。
// LoginFunc 装饰器模式的重点：需要一个函数签名。
type LoginFunc func(name string, pass string)
type LogDecorator func(LoginFunc) LoginFunc

// 装饰器重点：func中 传入特定func，且依然输出特定func。

// LogToRedis
func LogToRedis(f LoginFunc) LoginFunc {
    return func(name string, pass string) {
        // 业务逻辑
        fmt.Println("记录日志到redis")
        f(name, pass)
    }
}
```
