package loadbalance

type LoadBalance interface {
	SetBackendServer(backend []string)
	Run(addr string)
}
