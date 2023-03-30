package subject_object_pattern

// Subject 被观察者对象要实现的接口
type Subject interface {
	// Register 注册方法
	Register(observer Observer)
	// Deregister 注销方法
	Deregister(observer Observer)
	// NotifyAll 通知所有观察者方法
	notifyAll()
	// NotifySingle 通知单一观察者方法
	notifySingle(id string)
}

// Observer 观察者的对象要实现的接口
type Observer interface {
	// 更新方法
	update(string)
	// 取得id方法
	getID() string
}
