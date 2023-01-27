package proxy_pattern

import "fmt"

type UserService struct {

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
	// 使用装饰器模式调用
	return decorator(u.svc.Login)
}


// 结合装饰器模式。

type LoginFunc func(name string, pass string)
type LogDecorator func(LoginFunc) LoginFunc

func LogToRedis(f LoginFunc) LoginFunc {
	return func(name string, pass string) {
		// 业务逻辑
		fmt.Println("记录日志到redis")
		f(name, pass)
	}
}

func LogToEtcd(f LoginFunc) LoginFunc {
	return func(name string, pass string) {
		// 业务逻辑
		fmt.Println("记录日志到etcd")
		f(name, pass)
	}
}