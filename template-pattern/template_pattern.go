package template_pattern

import "fmt"

type ICache interface {
	Get(id int) interface{} // 获取缓存
	Output(id int) interface{} // 输出缓存
}

// 模版类
type Cache struct {
	c ICache // 需要
}

func (c *Cache) Get(id int) interface{} {
	return nil
}

func (cache *Cache) Output(id int) interface{} {
	return map[string]interface{}{
		"code": 200,
		"message": "success",
		"result": cache.c.Get(id),

	}
}


type RedisCache struct {
	*Cache
}

// 这里容易搞混
func NewRedisCache() *RedisCache {
	redisCache := &RedisCache{}
	redisCache.Cache = &Cache{redisCache}
	return redisCache
}

func (rc *RedisCache) Get(id int) interface{} {
	fmt.Println("从redis获取缓存")

	return "从redis获取缓存"
}

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

