package option_pattern

type HttpClient struct {
	Timeout int
	MaxIdle int
	ErrCallBack func(error)
}


type ClientOption func(client *HttpClient)
type ClientOptions []ClientOption

// apply 调用
func (cos ClientOptions) apply(c *HttpClient) {
	for _, opt := range cos {
		opt(c) // 调用
	}
}

// NewHttpClient
func NewHttpClient(opts ...ClientOption) *HttpClient {
	c := &HttpClient{}

	//for _, opt := range opts {
	//	opt(c)
	//}

	ClientOptions(opts).apply(c)

	return c
}

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
