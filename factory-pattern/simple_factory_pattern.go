package factory_pattern

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
