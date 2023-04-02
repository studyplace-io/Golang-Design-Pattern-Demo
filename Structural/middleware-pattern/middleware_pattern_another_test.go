package middleware_pattern

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMiddlewarePattern1(t *testing.T) {

	// 加入中间件，第一个入参：主handler，后面可传入多个中间件。
	handler := ApplyMiddleware(MyHandler, LoggerMiddleware)
	fmt.Println("httpserver start!!")
	http.ListenAndServe(":8080", handler)

}
