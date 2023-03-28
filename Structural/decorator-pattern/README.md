### 装饰器模式介绍：
装饰器模式主要使用场景：在原来的对象中需要增加一些"额外"的功能时使用。

- 一句话概括：不改变原对象的情况下，动态地给该对象增加一些职责（即增加其额外功能）的模式。
- 实现方法：分为两种：
    1. 采用函数签名实现
    2. 采用实现接口实现
- 使用方法：
    1. 调用时需要动态传入原对象给装饰器方法或装饰器对象。

### 示例：
```go
1. 定义一个方法签名(原本需要装饰的func) ，并实现装饰的功能。

type User struct {
    Id 	 int
    Name string
}

// UserInfoFunc 装饰器模式最重要的地方，需要搞一个type func
type UserInfoFunc func(id int) *User

func GetInfo(id int) *User {
    return &User{
        Id:   id,
        Name: "test-user",
    }
}

// GetInfoWithRole 装饰一下。
// 装饰器功能：入参是一个UserInfoFunc  出参也是一个UserInfoFunc，
// 并在返回的func中做一些功能的增强。
func GetInfoWithRole(fn UserInfoFunc) UserInfoFunc {
    return func(id int) *User {
        user := fn(id)
        user.Name = "decorator-pattern" + user.Name
    
        return user
    }
}

2.使用接口的方式实现，所有的装饰器都需要实现接口对象的方法。

// IPizza 披萨对象接口
type IPizza interface {
    GetPrice() int
}

type Pizza struct {
    pizzaprice int
}

func NewPizza(price int) *Pizza {
    return &Pizza{
        pizzaprice: price,
    }
}

func (p *Pizza) GetPrice() int {
    return p.pizzaprice
}
```
