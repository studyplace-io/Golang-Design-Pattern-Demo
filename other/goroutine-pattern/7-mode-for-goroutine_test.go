package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// https://mp.weixin.qq.com/s/9Y6DsciTeZEn2RpTRysvow

func TestExample(t *testing.T) {
	//UseChanToStopGoroutine()
	//UseChanToBroadcast()

}

// 1. 使用Map实现Set
// 使用空结构体struct{}，意味着你的map的value部分并不占用空间。

type Set struct {
	MapToSet map[string]struct{}
}

func (s *Set) UseMapForSet(name string) error {

	if _, ok := s.MapToSet[name]; ok {
		return nil
	}
	s.MapToSet[name] = struct{}{}

	// ..... 省略

	return errors.New("aaaaa")

}

// 2. 使用chan struct{}实现Goroutine的同步

func helloworld(stopChan chan struct{}) {

	for {

		select {
		case <-stopChan:
			fmt.Println("收到停止消息！")
			return
		default:
			// 执行业务逻辑
			time.Sleep(time.Second)
			fmt.Println("hello world!!")
		}
	}
}

func UseChanToStopGoroutine() {
	stopC := make(chan struct{})
	go helloworld(stopC)

	time.Sleep(time.Second * 10)
	stopC <- struct{}{}
	fmt.Println("主goroutine结束！")
}

// 3. 使用Close广播
// 如果我们运行多个go helloworld(stopC)，那么不需要发送多个struct{}{}stopC，我们只需要关闭stopC通道来广播信号:
func UseChanToBroadcast() {
	stopC := make(chan struct{})
	go helloworld(stopC)
	go helloworld(stopC)

	time.Sleep(time.Second * 10)

	close(stopC)
	fmt.Println("主goroutine结束！")
}

// 后面几个模式没有太搞懂
