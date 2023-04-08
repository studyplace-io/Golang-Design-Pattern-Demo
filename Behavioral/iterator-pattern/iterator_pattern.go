package iterator_pattern

// Iterator 迭代器接口
type Iterator interface {
	// 是否有下一个的方法
	HasNext() bool
	// 拿到下一个的方法
	GetNext() *User
}

// Collection 收集器接口对象：负责注入需要迭代的对象，返回迭代器对象
type Collection interface {
	CreateIterator() Iterator
}


// User 对象
type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

// UserIterator 对象迭代器
type UserIterator struct {
	index int  // 维护一个索引标示
	users []*User
}

func (u *UserIterator) HasNext() bool {
	// 判断维护的索引是否已经到底了
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) GetNext() *User {
	// 如果有下一个就取下一个
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

// UserCollection 收集器对象，负责保存所有迭代的对象
type UserCollection struct {
	Users []*User
}


func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		users: u.Users,
	}
}