package SubjectObjectPattern

// 观察者的对象 接口
type Observer interface {
	// 更新方法
	update(string)
	// 取得id方法
	getID() string
}
