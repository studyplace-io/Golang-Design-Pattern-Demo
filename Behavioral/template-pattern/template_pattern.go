package template_pattern

import "fmt"

// ICache 具体模版对象需要实现的接口
type ICache interface {
	Get(id int) interface{} 	// 获取缓存
	Output(id int) interface{}  // 输出缓存
}


// Cache 抽象模版对象，内部需要有接口对象
type Cache struct {
	c ICache  // 接口对象
}

// Output 抽象模版对象实现的方法，不一定要命名成Output，可以自定义，但也可以同时实现ICache接口
func (cache *Cache) Output(id int) interface{} {
	return map[string]interface{}{
		"code": 200,
		"message": "success",
		"result": cache.c.Get(id), // 调用接口中的方法
	}
}

var _ ICache = &RedisCache{}

// RedisCache 具体模版对象
type RedisCache struct {
	// 模版模式的重点：具体模版对象需要使用"嵌套"，套入抽象模版对象
	// 嵌套的对象实现Output(id int) interface{}方法，所以RedisCache也实现ICache接口
	*Cache
}

// NewRedisCache 重点。
func NewRedisCache() *RedisCache {
	// 这里嵌套的方式是重点
	redisCache := &RedisCache{}
	redisCache.Cache = &Cache{redisCache}
	return redisCache
}

func (rc *RedisCache) Get(id int) interface{} {
	fmt.Println("从redis获取缓存")

	return "从redis获取缓存"
}

var _ ICache = &EtcdCache{}

type EtcdCache struct {
	*Cache
}

func (ec *EtcdCache) Get(id int) interface{} {
	fmt.Println("从etcd获取缓存")

	return "从etcd获取缓存"
}

func NewEtcdCache() *EtcdCache {
	etcdCache := &EtcdCache{}
	etcdCache.Cache = &Cache{etcdCache}
	return etcdCache
}

