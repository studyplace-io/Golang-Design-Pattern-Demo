package main

import (
	"fmt"
	"log"
	"net/http"
)

// 介绍装饰器在Go中的实现
// 装饰器在http中的使用
// 装饰器实现Pipeline

/*
	python语法的装饰器模式十分方便，直接用@语法糖即可


	from functools import wraps

	def decorator(func):
		@wraps(func) # @wraps接受一个函数来进行装饰，并加入了复制函数名称、注释文档、参数列表等等的功能。这可以让我们在装饰器里面访问在装饰之前的函数的属性。
		def decorated(*args, **kwargs):

			# 这里可以增加hook操作。
			print("start")
			##########

			# 调用原来函数
			result = func(*args, **kwargs) # 这里直接调用原函数 并没有修改他 只是在调用之前和之后增加了额外的处理逻辑

			# 这里可以增加hook操作。
			print("end")
			##########

			return result
		return decorated

	@decorator # 用装饰器装饰add 这种@语法糖确实很香啊
	def add(x):
	   return x + x

	print(add(4)) # 直接调用add

	>>> start
	>>> end
	>>> 8

 */


/*
	go语言实现装饰器模式：
	装饰器模式是遵循SOLID设计原则中的开放封闭原则的，即对扩展开放，对修改关闭。
 */



type add func(int) int

func addFunc(num int) int {
	return num + num
}

func decorator(f add) add {
	// 返回函数
	return func(i int) int {

		// 这里可以操作业务逻辑
		fmt.Println("start")

		// 调用真正的函数
		res := f(i)  // 直接调用

		// 这里可以调用hook业务逻辑
		fmt.Println("end")


		return res  // 返回计算结果

		
	}
}

func main() {
	//
	a := 4
	demo := decorator(addFunc)
	result := demo(a)
	fmt.Println(result)

}

// 事例二
// 扩展：装饰器模式在http Wed开发的实践
// http经常使用装饰器对Handler进行中间件插件化功能：ex 超时取消控制 请求认证 等功能


//装饰器
func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	// 返回一个ResponseWriter写入器，Request请求
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithServerHeader()")
		w.Header().Set("Server", "HelloServer")
		h(w, r)
	}
}

func WithAuthCookie(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithAuthCookie()")
		cookie := &http.Cookie{Name: "Auth", Value: "Pass", Path: "/"}
		http.SetCookie(w, cookie)
		h(w, r)
	}
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved Request %s from %s\n", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}

func HttpUseDecoratorBad() {
	// 装饰器装饰helloworld
	http.HandleFunc("/api/v1/hello", WithServerHeader(helloworld))
	// 多个装饰器装饰helloworld
	// 如果单用简单的装饰器模式，会导致装饰器需要层层嵌套。
	http.HandleFunc("/api/v1/hello", WithServerHeader(WithAuthCookie(helloworld)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 装饰器使用pipeline模式。 完善插件化中间件！

type HttpHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

func Handler(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		h = d(h)
	}
	return h
}

func HttpUseDecoratorGood()  {
	http.HandleFunc("/api/v1/hello", Handler(helloworld,
		WithServerHeader, WithAuthCookie))

}



