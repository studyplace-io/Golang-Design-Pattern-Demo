package main

import (
	"github.com/gin-gonic/gin"
	round "github.com/practice/Design-Patterns-practice/other/loadbalance-pattern/round_robin"
)

func main() {

	// 模拟后端存储
	go server1()
	go server2()
	go server3()

	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	// 随机模式
	//rl := random.NewLoadBalancerForRandomMode()
	//rl.SetBackendServer(backends)

	// 轮询模式
	rl := round.NewLoadBalancerForRoundRobinMode()
	rl.SetBackendServer(backends)
	rl.Run(":8080")

}

func server1() {

	router := gin.Default()

	// 定义路由和处理函数
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!, server1",
		})
	})

	// 启动服务器
	router.Run(":8081")
}

func server2() {
	// 创建Gin引擎
	router := gin.Default()

	// 定义路由和处理函数
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!, server2",
		})
	})

	// 启动服务器
	router.Run(":8082")
}

func server3() {
	// 创建Gin引擎
	router := gin.Default()

	// 定义路由和处理函数
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!, server3",
		})
	})

	// 启动服务器
	router.Run(":8083")
}
