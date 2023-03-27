package proxy_pattern

import "testing"

func TestProxyPattern(t *testing.T) {

	// 一般调用方式
	user := NewUserService()
	user.Login("tt", "test")

	// 使用proxy代理模式的调用方式
	userProxy := NewUserProxy(user)
	userProxy.Login(LogToEtcd)("test", "test-pass")

}
