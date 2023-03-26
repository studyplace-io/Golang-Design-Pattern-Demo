### 中间件模式介绍：
中间件模式主要使用在不同系统间"传递请求"、"处理响应"的设计模式。

- 一句话概括：可以在client端-server端中实现某些特定功能(ex: 验证、授权、准入或是限流等功能)。
- 实现方法：主要分为两种：
  1. 自定义实现类似gin框架的中间件模式。
  2. 根据http Func 实现。
- 使用方法：
    1. 在中间件模式中，可以自定义扩展需要的中间件插件。
    2. 调用时需要特别注意："注册"(需要在执行流程前，注册) "执行后处理"(需要记得调用Next()方法)

![](https://github.com/StudyPlace-io/Golang-Design-Pattern-Demo/blob/main/image/middleware.jpg?raw=true)
![](https://github.com/StudyPlace-io/Golang-Design-Pattern-Demo/blob/main/image/middleware-4.png?raw=true)

### 示例：
```go
1. 模拟gin框架的中间件模式
// MainFunc 中间件框架上主要实现的方法，不论是中间件方法或是业务逻辑方法，都要实现。
type MainFunc func(c *Context)

// Context 模仿gin的中间件模式
type Context struct {
	// 想成洋葱模型。
	handlers []MainFunc  // 里面装的除了中间件方法，还有真正的业务逻辑方法
	index int // 标示，记录下一个要执行的中间件
}

// Use 把自己实现的中间件Func 加入到维护的切片中。
func (c *Context) Use(f MainFunc) {
	c.handlers = append(c.handlers, f)
}

// Next 中间件Func中需要调用此方法
func (c *Context) Next() {
	c.index++
	// 进入洋葱下一层，并调用
	Func := c.handlers[c.index]
	Func(c)
}

// Run 执行流程
func (c *Context) Run() {
	// 进入洋葱的最外层。 并且调用
	firstFunc := c.handlers[0]
	firstFunc(c)
}

2. 基于http Func实现中间件
// Middleware 中间件Func，仔细看就能发现，此Func 接受一个与"主方法"一样的方法，并在把它输出
type Middleware func(http.HandlerFunc) http.HandlerFunc

// LoggerMiddleware 日志中间件
func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 这里可以做准备工作
        log.Printf("%s %s", r.Method, r.URL.Path)
        next(w, r)  // 这里会调用到下一个中间件
        // 执行后可以进行善后工作
        log.Println("log middleware completed!")
    }
}

// MyHandler 主要Handler
func MyHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}

// ApplyMiddleware 载入中间件方法，第一个入参是主要handler 后续接受一个切片，用来执行中间件Func
func ApplyMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
  for _, middleware := range middlewares {
    handler = middleware(handler)
  }
  return handler
}

```
