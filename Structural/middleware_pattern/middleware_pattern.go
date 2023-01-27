package middleware_pattern

type Context struct {
	// 想成洋葱模型。
	handlers []func(c *Context)
	index int
}

func (c *Context) Use(f func(c *Context)) {
	c.handlers = append(c.handlers, f)
}

func (c *Context) Next() {
	c.index++
	// 进入洋葱下一层，并调用
	Func := c.handlers[c.index]
	Func(c)
}

func (c *Context) Run() {
	// 进入洋葱的最外层。 并且调用
	firstFunc := c.handlers[0]
	firstFunc(c)
}

// 类似gin的GET方法 (PUT POST ... 等)
func (c *Context) GET(path string, f func(c *Context)) {
	c.handlers = append(c.handlers, f)
}


