package ProxyPattern

// 服务
type Server interface {
	HandleRequest(string, string) (int, string)
}