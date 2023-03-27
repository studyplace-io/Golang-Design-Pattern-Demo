package proxy_pattern

// NginxProxy 代理类对象
type NginxProxy struct {
	// 真正的对象访问的对象
	application *Application
	// 最大限流量
	maxAllowedRequest int
	// 当前限流量
	rateLimiter map[string]int
}

const (
	defaultName              = "default-application"
	defaultOther             = "default-other"
	defaultMaxAllowedRequest = 10
)

// NewNginxProxyServer 创建新的对象，加入选项模式
func NewNginxProxyServer(opts ...NginxProxyOption) *NginxProxy {

	instance := &NginxProxy{
		application: &Application{
			name:  defaultName,
			other: defaultOther,
		},
		maxAllowedRequest: defaultMaxAllowedRequest,
		rateLimiter:       make(map[string]int),
	}

	for _, opt := range opts {
		opt(instance)
	}

	return instance
}

type NginxProxyOption func(proxy *NginxProxy)

func WithMaxAllowedRequest(maxAllowedRequest int) NginxProxyOption {
	return func(proxy *NginxProxy) {
		proxy.maxAllowedRequest = maxAllowedRequest
	}
}

func WithName(name string) NginxProxyOption {
	return func(proxy *NginxProxy) {
		proxy.application.name = name
	}
}

func WithOther(other string) NginxProxyOption {
	return func(proxy *NginxProxy) {
		proxy.application.other = other
	}
}

// HandleRequest 调用代理类的HandleRequest方法
// 最重要的地方：需要实现与原来对象一样的方法。
func (n *NginxProxy) HandleRequest(url, method string) (int, string) {
	// 做一下限流功能
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	// 如果成功，就进行真正的application.handleRequest(url, method)！
	return n.application.HandleRequest(url, method)
}

// checkRateLimiting 限流方法
func (n *NginxProxy) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}