package decorator_pattern

type User struct {
	Id int
	Name string
}

// 装饰器模式最重要的地方，需要搞一个type func
type UserInfoFunc func(id int) *User

func GetInfo(id int) *User {
	return &User{
		Id: id,
		Name: "test-user",
	}
}

// 装饰一下。
func GetInfoWithRole(fn UserInfoFunc) UserInfoFunc {
	return func(id int) *User {
		user := fn(id)
		user.Name = "decorator-pattern" + user.Name

		return user
	}
}

