### 单例模式介绍：
简介：单例对象必须保证只有一个实例存在。许多时候整个系统只需要拥有一个的全局对象

- 一句话概括：整个系统只需要一个全局对象。
- 实现方法：golang语法中使用sync.Do()实现
  
### 示例：

```go
// Singleton 全局只能初始化一次的对象
type Singleton struct {
	info string
}

// 只要初始化后就赋值给singleton变量。
var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
```
