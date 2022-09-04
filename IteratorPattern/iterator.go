package IteratorPattern

// 迭代器接口
type Iterator interface {
	// 是否有下一个的方法
	HasNext() bool
	// 拿到下一个的方法
	GetNext() *User
}
