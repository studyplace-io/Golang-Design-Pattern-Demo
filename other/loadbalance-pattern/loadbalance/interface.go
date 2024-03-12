package loadbalance

// LoadBalance 负载均衡代理接口
type LoadBalance interface {
	// SetBackendServer 设置后端 server
	SetBackendServer(backend []string)
	// Run 启动反向代理
	Run(addr string)
}
