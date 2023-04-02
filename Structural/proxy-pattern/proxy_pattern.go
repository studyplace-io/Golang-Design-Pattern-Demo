package proxy_pattern

import "fmt"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Login(name string, pass string) {
	fmt.Println("登入成功！！")
}


type UserProxy struct {
	svc *UserService
}

func NewUserProxy(svc *UserService) *UserProxy {
	return &UserProxy{
		svc: svc,
	}
}

func (u *UserProxy) Login(decorator LogDecorator) LoginFunc {
	fmt.Println("记录日志！！")
	// 使用代理模式抽象出来
	// 使用装饰器模式调用f方法
	return decorator(u.svc.Login)
}


// 结合装饰器模式。
// LoginFunc 装饰器模式的重点：需要一个函数签名。
type LoginFunc func(name string, pass string)
type LogDecorator func(LoginFunc) LoginFunc

// 装饰器重点：func中 传入特定func，且依然输出特定func。

// LogToRedis
func LogToRedis(f LoginFunc) LoginFunc {
	return func(name string, pass string) {
		// 业务逻辑
		fmt.Println("记录日志到redis")
		f(name, pass)
	}
}

// LogToEtcd
func LogToEtcd(f LoginFunc) LoginFunc {
	return func(name string, pass string) {
		// 业务逻辑
		fmt.Println("记录日志到etcd")
		f(name, pass)
	}
}