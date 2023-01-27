package StrategyPatterns

// 缓存的地方
type Cache struct {
	// 用map缓存
	storage      map[string]string
	// 接口对象
	strategyinterface StrategyInterface
	// 元data
	capacity     int
	maxCapacity  int
}

func InitCache(e StrategyInterface) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		strategyinterface: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

// input一个接口对象  "就是实现此接口的方法"。
func (c *Cache) SetStrategyInterface(e StrategyInterface) {
	c.strategyinterface = e
}

// 加入map中
func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.method()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) method() {
	c.strategyinterface.method(c)
	c.capacity--
}
