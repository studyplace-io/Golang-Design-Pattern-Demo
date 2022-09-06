package main

import "fmt"

// 多个对象实现相同接口


// 定义对外的接口，提供两种方法
// Start Log两种
type Service interface {
	Start() error
	Log(string)

}

// Logger对象
type Logger struct {
	// 可以自己加入字段
}

// Logger的方法Log
func (l *Logger) Log(name string) {
	fmt.Println(name)
	fmt.Println("调用Log方法")

}

// GameService对象
type GameService struct {
	*Logger
	// 还可以自己增加字段 ...
}

// GameService的方法Start
func (g *GameService) Start() error {

	fmt.Println("调用Start方法")
	return nil
}

func main() {
	// GameService同时实现 Log Start方法。
	// 相当于实现Service接口
	var s Service = &GameService{
		&Logger{},
	}
	_ = s.Start()
	s.Log("jiang")
}