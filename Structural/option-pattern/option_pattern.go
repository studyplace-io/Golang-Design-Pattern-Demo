package option_pattern

import (
	"fmt"
	"net/http"
)

type HttpClient struct {
	Client *http.Client
	// 下列字段都是配置
	Timeout     int
	MaxIdle     int
	ErrCallBack func(error)
}

const (
	defaultTimeout = 10
	defaultMaxIdle = 10
)

var (
	defaultErrCallBack = func(err error) {
		fmt.Printf("call back error: %s\n", err)
	}
)

// ClientOption 选项模式的最重要的步骤。
// 必须要传入一个func，这个func可以操作要"创建的对象"，如此才能改变对象中的配置
type ClientOption func(client *HttpClient)
type ClientOptions []ClientOption

// apply 调用
func (cos ClientOptions) apply(c *HttpClient) {
	for _, opt := range cos {
		opt(c) // 调用
	}
}

// NewHttpClient 创建一个HttpClient对象，默认配置，但开放选项模式让用户自定义配置。
func NewHttpClient(opts ...ClientOption) *HttpClient {
	c := &HttpClient{
		Timeout:     defaultTimeout,
		MaxIdle:     defaultMaxIdle,
		ErrCallBack: defaultErrCallBack,
	}

	// 法一：直接遍历，并执行ClientOption方法
	//for _, opt := range opts {
	//	opt(c)
	//}

	// 法二：可做成一个切片，再执行apply()方法
	ClientOptions(opts).apply(c)

	return c
}

// WithXXX方法中还能做入参校验，过滤不符合入参的情况

func WithTimeout(timeout int) ClientOption {
	return func(client *HttpClient) {
		client.Timeout = timeout
	}
}

func WithMaxIdle(MaxIdle int) ClientOption {
	return func(client *HttpClient) {
		client.MaxIdle = MaxIdle
	}
}

func WithErrCallback(callback func(error)) ClientOption {
	return func(client *HttpClient) {
		client.ErrCallBack = callback
	}
}
