package cache_proxy_pattern

// RequestProvider 请求抽象接口
type RequestProvider interface {
	DoRequest(requestId string) string
}

// CacheRequestProvider 缓存代理对象
type CacheRequestProvider struct {
	provider RequestProvider	// 接口对象
	// TODO: 这里可以实现一个LRU的缓存淘汰策略
	cache    map[string]string 	// 缓存map
}

func NewCacheRequestProvider(provider RequestProvider) *CacheRequestProvider {
	return &CacheRequestProvider{
		provider: provider,
		cache: make(map[string]string),
	}
}

// DoRequest 调用请求: 先从缓存里查，没有就调用具体对象远程请求
func (c *CacheRequestProvider) DoRequest(requestId string) string {
	// 先查，如果没有就调用，且存入缓存
	request, ok := c.cache[requestId]
	if !ok {
		request = c.provider.DoRequest(requestId)
		c.cache[requestId] = request
	}

	return request
}

