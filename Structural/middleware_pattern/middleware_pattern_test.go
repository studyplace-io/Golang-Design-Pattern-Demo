package middleware_pattern

import (
	"fmt"
	"testing"
)

func TestMiddlewarePattern(t *testing.T) {

	exampleContext := &Context{}
	// 加入中间件
	exampleContext.Use(middleware1())
	exampleContext.Use(middleware2())
	exampleContext.Use(middleware3())

	// 调用核心，这个方法不加Next.()
	exampleContext.GET("/", func(c *Context) {
		fmt.Println("handler func!!")
	})


	exampleContext.Run()


}


func middleware1() func(c *Context) {
	return func(c *Context) {
		// 实现业务逻辑
		fmt.Println("执行中间件1")
		c.Next()
		fmt.Println("执行完毕中间件1")
	}
}

func middleware2() func(c *Context) {
	return func(c *Context) {
		fmt.Println("执行中间件2")
		c.Next()
		fmt.Println("执行完毕中间件2")
	}
}

func middleware3() func(c *Context) {
	return func(c *Context) {
		fmt.Println("执行中间件3")
		c.Next()
		fmt.Println("执行完毕中间件3")
	}
}


