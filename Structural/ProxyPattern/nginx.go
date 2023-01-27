package ProxyPattern

// 代理类对象
type Nginx struct {
	// 真正的对象
	application       *Application
	// 最大限流量
	maxAllowedRequest int
	// 当前限流量
	rateLimiter       map[string]int
}

// 创建新的对象
func NewNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

// 调用代理类的HandleRequest方法
func (n *Nginx) HandleRequest(url, method string) (int, string) {
	// 做一下限流功能
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	// 如果成功，就进行真正的application.handleRequest(url, method)！
	return n.application.HandleRequest(url, method)
}

// 限流方法
func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
