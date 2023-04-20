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

	// 附注：调用核心，这个方法不加Next()
	exampleContext.GET("/", func(c *Context) {
		fmt.Println("handler func!!")
	})

	// 执行
	exampleContext.Run() // 会从第一个输入的中间件开始执行。

}

// middleware1PreHook 执行中间件具体逻辑前的准备工作。
func middleware1PreHook() {
	// 可以在这里准备需要的配置or资源等等
	fmt.Println("middleware1即将执行，调用PreHook做准备工作")
}

// middleware1PostHook 执行此中间件逻辑后的善后工作
func middleware1PostHook() {
	fmt.Println("middleware1执行完毕，调用PostHook做善后工作")
}

func middleware1() MainFunc {
	return func(c *Context) {
		// 准备
		middleware1PreHook()
		// 这里就是可以写的中间件逻辑。。。。
		fmt.Println("执行中间件1")
		c.Next() // 执行下一个中间件
		fmt.Println("执行完毕中间件1")
		// 善后
		middleware1PostHook()
	}
}

func middleware2() MainFunc {
	return func(c *Context) {
		fmt.Println("执行中间件2")
		c.Next()
		fmt.Println("执行完毕中间件2")
	}
}

func middleware3() MainFunc {
	return func(c *Context) {
		fmt.Println("执行中间件3")
		c.Next()
		fmt.Println("执行完毕中间件3")
	}
}
