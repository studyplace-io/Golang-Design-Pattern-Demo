package Singleton_pattern

import "sync"

/*
 单例模式：使用懒惰模式的单例模式，使用双重检查加锁保证线程安全。
 有分饿汉和懒汉：注意区别。
 如果用懒汉：需要注意并发安全 需要用锁。
*/

// Singleton 全局只能初始化一次的对象
type Singleton struct {
	info string
}

// 只要初始化后就赋值给singleton变量。
var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}

func (s *Singleton) GetInstanceInfo() string {
	return s.info
}

func (s *Singleton) SetInstance(info string) {
	s.info = info
}