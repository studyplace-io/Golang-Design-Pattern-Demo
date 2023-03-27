package middleware_pattern

import "fmt"

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

// 类似gin的GET方法 (PUT POST ... 等)
func (c *Context) GET(path string, f MainFunc) {
	// 记得需要加入的中间件后面！！！
	c.handlers = append(c.handlers, f)
}

func (c *Context) DoSomething(somethingInput string, f MainFunc) {
	// 记得需要加入的中间件后面！！！
	c.handlers = append(c.handlers, f)

	// 这里需要执行具体逻辑
	fmt.Println("do something important !!", somethingInput)
}

