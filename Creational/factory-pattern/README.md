### 工厂方法介绍：
工厂方法有包括：简单工厂、工厂方法与抽象工厂等模式。(本文介绍常用的前两种)


- 一句话概括：用来创建"不同"但是关连类型的struct(如：都是人，但是男人女人，白人黑人等)。
- 实现方法：实现某IFactory接口类，并使用传入的参数决定生成何种类型对象。
- 使用方法：
    1. 当创建的对象逻辑简单，可以直接使用new或是简单工厂。
    2. 当创建的对象逻辑复杂，可以将创建的逻辑拆分到不同工厂中，基于一个大工厂下细分多个小工厂。
  
### 示例：
1.使用一个UserCreateFunc方法，创建对象。(法二：也可以实现UserFactory，且让每个user都实现此接口)
```go
// 法一：可以实现此接口
// UserFactory 用户工厂
//type UserFactory interface {
//
//}

// 法二：实现此方法，返回interface{}
type UserCreateFunc func(id int, name string) interface{}

type User struct {
	Id int
	UserName string
}

func NewUser() UserCreateFunc {
	return func(id int, name string) interface{} {
		return &User{
			Id: id,
			UserName: name,
		}
	}
}



type Admin struct {
	Id int
	AdminName string
	Role string
}

func NewAdmin() UserCreateFunc {
	return func(id int, name string) interface{} {
		return &Admin{
			Id: id,
			AdminName: name,
			Role: "admin-role",
		}
	}
}

```
2. 当需要创建的总类较多且逻辑复杂，可使用工厂方法。需要实现
```go
/*
	当有不同关连种类的对象需要创建，可以用工厂方法
 */

// IProductFactory 总工厂
type IProductFactory interface {
	// 注意：这里返回接口对象，可以看成工厂的"对象"
	CreateProduct(t ProductType) IProduct
}

// IProduct 产品接口，"所有要被创建的对象，都要实现此接口"
type IProduct interface {
	// 可以抽象出不同的方法
	GetInfo() string
	SetInfo()
}

// ProductType 区分标示
type ProductType int

```
![](https://github.com/StudyPlace-io/Golang-Design-Pattern-Demo/blob/feat_factory_method/image/simple_factory.jpg?raw=true)
![](https://github.com/StudyPlace-io/Golang-Design-Pattern-Demo/blob/feat_factory_method/image/factory_method.jpg?raw=true)

