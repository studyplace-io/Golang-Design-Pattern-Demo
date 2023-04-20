package decorator_pattern

type User struct {
	Id   int
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
