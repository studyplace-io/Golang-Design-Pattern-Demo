package Singleton_pattern

import "sync"

/*
	单例模式：使用懒惰模式的单例模式，使用双重检查加锁保证线程安全。
	有分饿汉和懒汉：注意区别。
	如果用懒汉：需要注意并发安全 需要用锁。
 */

type Singleton struct {}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
	
}
