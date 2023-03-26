package middleware_pattern

import (
	"log"
	"net/http"
)

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



