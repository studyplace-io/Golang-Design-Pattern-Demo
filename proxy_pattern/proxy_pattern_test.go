package proxy_pattern

import "testing"

func TestProxyPattern(t *testing.T) {

	user := new(UserService)
	userProxy := NewUserProxy(user)

	userProxy.Login(LogToEtcd)("test", "test-pass")

}
