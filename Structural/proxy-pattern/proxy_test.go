package proxy_pattern

import (
	"fmt"
	"testing"
)

func TestProxyPattern(t *testing.T) {

	// 两个url请求
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	// 原来服务调用方式
	application := NewApplication(
		WithApplicationName("aaa"), WithApplicationOther("otheraaa"))

	httpCode, body := application.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = application.HandleRequest(createUserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = application.HandleRequest(createUserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	// 加入proxy模式后的调用方式
	applicationProxy := NewNginxProxyServer(WithMaxAllowedRequest(5), WithName("proxy-test"),
		WithOther("proxy-other"))
	httpCode, body = applicationProxy.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

}
