package SubjectObjectPattern

// 被观察者的接口
type Subject interface {
	// 注册方法
	register(observer Observer)
	// 注销方法
	deregister(observer Observer)
	// 通知方法
	notifyAll()
}