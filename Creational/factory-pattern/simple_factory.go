package factory_pattern

const (
	NormalUser = iota
	AdminUser
)

type UserType int

// CreateUser 通过一个传参，分辨要创建何种类型
func CreateUser(t UserType) UserCreateFunc {
	switch t {
	case NormalUser:
		return NewUser()
	case AdminUser:
		return NewAdmin()
	}

	return nil
}
