package factory_pattern

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
