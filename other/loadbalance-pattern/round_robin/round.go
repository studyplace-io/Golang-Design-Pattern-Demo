package round

import (
	"github.com/practice/Design-Patterns-practice/other/loadbalance-pattern/loadbalance"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type roundLoadBalance struct {
	healthyCheck    time.Duration
	proxy           *httputil.ReverseProxy
	backends        []*Server
	healthyBackends []*Server
	mu              sync.Mutex
}

const (
	defaultHealthyCheck = time.Second * 10
)

type Server struct {
	URL *url.URL
	// Healthy 标示健康状态
	Healthy bool
}

// NewLoadBalancerForRoundRobinMode 使用轮询模式
func NewLoadBalancerForRoundRobinMode() loadbalance.LoadBalance {
	// 创建负载均衡器实例
	lb := &roundLoadBalance{}

	// 创建反向代理器, 也可以用
	// ServeHTTP(w http.ResponseWriter, r *http.Request) 方法实现
	lb.proxy = &httputil.ReverseProxy{
		Director: lb.director,
	}

	return lb
}

func (lb *roundLoadBalance) SetBackendServer(backends []string) {
	// 启动负载均衡器监听的端口
	for _, u := range backends {
		serverURL, err := url.Parse(u)
		if err != nil {
			log.Fatal(err)
		}
		lb.backends = append(lb.backends, &Server{
			URL:     serverURL,
			Healthy: true,
		})
	}
	lb.healthyBackends = lb.backends
}

func (lb *roundLoadBalance) Run(addr string) {

	if len(lb.backends) == 0 {
		panic("backend server can't be nil, use SetBackendServer method to add backend server")
	}

	go lb.healthCheck()
	// 启动负载均衡器监听的端口
	err := http.ListenAndServe(addr, lb.proxy)
	if err != nil {
		log.Fatal(err)
	}
}

// director 反向代理方法，选择目标服务器
func (lb *roundLoadBalance) director(req *http.Request) {
	targetServer := lb.randomHealthyServer()
	// 设置目标服务器的主机
	req.URL.Scheme = targetServer.URL.Scheme
	req.URL.Host = targetServer.URL.Host
	req.URL.Path = singleJoiningSlash(targetServer.URL.Path, req.URL.Path)
	req.Host = targetServer.URL.Host
}

func (lb *roundLoadBalance) healthCheck() {
	if lb.healthyCheck == time.Duration(0) {
		lb.healthyCheck = defaultHealthyCheck
	}

	ticker1 := time.NewTicker(lb.healthyCheck)

	for {
		select {
		case <-ticker1.C:
			lb.mu.Lock()
			for _, server := range lb.backends {
				// 发送健康检查请求, 修改状态
				resp, err := http.Get(server.URL.String())
				if err != nil || resp.StatusCode != http.StatusOK {
					server.Healthy = false
				} else {
					server.Healthy = true
				}
			}
			lb.mu.Unlock()
		}
	}
}

func (lb *roundLoadBalance) randomHealthyServer() *Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	healthyServers := make([]*Server, 0)

	// 选择健康的服务器
	for _, server := range lb.healthyBackends {
		if server.Healthy {
			healthyServers = append(healthyServers, server)
		}
	}

	if len(healthyServers) == 0 {
		return nil
	}

	server := healthyServers[0]
	lb.healthyBackends = append(lb.healthyBackends[1:], server)

	return server
}

func singleJoiningSlash(a, b string) string {
	aslash := len(a) > 0 && a[len(a)-1] == '/'
	bslash := len(b) > 0 && b[0] == '/'
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
