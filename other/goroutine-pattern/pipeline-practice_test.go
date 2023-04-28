package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

// https://mp.weixin.qq.com/s/IuGaKbQvW7z1KsoMhB8mgg
// pipeline 将数据流分为多个环节，channel 用于数据流动，goroutine 用于处理数据。
// fan-out 用于分发任务，fan-in 用于数据整合，通过 FAN 模式可以让流水线更好地并发。

func TestExample1(t *testing.T) {
	TryPipeline()
	//TryFanINFanOut()
}

/*
	工作模式并不高效，会让整个流水线的效率取决于最慢的环节。
	因为每个环节中的任务量是不同的，这意味着我们需要的机器资源是存在差异的。
	任务量小的环节，尽量占有少量的机器资源，任务量重的环节，需要更多线程并行处理。
*/

// generate 函数：它充当生产者角色，将数据写入 channel，并把该 channel 返回。当所有数据写入完毕，关闭 channel。
func generate(nums ...int) <-chan int {
	// 非阻塞chan
	out := make(chan int)
	// 启一个gorouitne 把data放入chan，当做生产者
	go func() {
		for _, num := range nums {
			out <- num
		}
		// 只有一个生产者时，执行完毕记得要关闭chan
		close(out)
	}()

	return out

}

// square 函数：它是数据处理的角色，从开始环节中的 channel  取出数据，计算平方，将结果写入新的 channel ，并把该新的 channel 返回。当所有数据计算完毕，关闭该新 channel。
func square(inputChan <-chan int) <-chan int {
	// 返回结果的chan
	out := make(chan int)

	// 启一个goroutine，从inputChan拿出数据进行处理，再放入out中
	go func() {

		for num := range inputChan {
			out <- num * num
		}
		// 生产者只有一个时，使用完毕需要关闭通道
		close(out)

	}()

	return out

}

// TryPipeline 负责编排整个 pipeline ，并充当消费者角色：读取square函数的 channel 数据，打印出来。
func TryPipeline() {

	// 串式调用
	c := generate(2, 3, 4, 5, 6, 3, 6)
	resultChan := square(c)

	// 主goroutine当做消费者。
	for value := range resultChan {
		fmt.Printf("value:%d \n", value)
	}

	fmt.Println("主goroutine结束")

}

///////////////////////////////////////////////////

// 负责调用FanINFanOut，使用两个generate1函数按照不同interval执行，
// 实现了接收系统系统中断信号（终端上执行 CTRL+C 即可发送中断信号）的优雅的关闭机制。
func TryFanINFanOut() {

	// 分发出两个goroutine
	mc1, stop1 := generate1("message from generator 1", time.Millisecond*200)
	mc2, stop2 := generate1("message from generator 2", time.Millisecond*300)

	// 整合多个chan的信息
	mmc, wg1 := multiplex(mc1, mc2)

	// 接收错误的chan
	errChan := make(chan error)

	// 启一个goroutine进行优雅关闭
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("\n TryFanINFanOut%s signal received", <-sc)
	}()

	// 启一个goroutine 当做消费者
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()

		for m := range mmc {
			fmt.Println(m)
		}
	}()

	// 如果有错误，要打印错误信息，如果没有，就会在这一直阻塞。
	if err := <-errChan; err != nil {
		fmt.Println(err.Error())
	}

	// 调用停止的函数
	stopGenerating(mc1, stop1)
	stopGenerating(mc2, stop2)
	// 阻塞再这里
	wg1.Wait()

	// 使用完毕记得要关闭
	close(mmc)
	wg2.Wait()

	fmt.Println("主goroutine退出")

}

func generate1(message string, interval time.Duration) (chan string, chan struct{}) {
	messageChan := make(chan string)
	stopChan := make(chan struct{})

	go func() {
		defer func() {
			close(stopChan)
		}()

		for {
			select {
			case <-stopChan:
				fmt.Println("generate goroutine stopped!")
				return
			default:
				time.Sleep(interval)
				messageChan <- message
			}
		}

	}()

	return messageChan, stopChan

}

// stopGenerating 函数通过通过向 sc 中传入空结构体，通知 generate退出，调用 close(messageChan) 关闭消息 channel
func stopGenerating(messageChan chan string, stopChan chan struct{}) {
	stopChan <- struct{}{}

	close(messageChan)
}

// 多路复用函数 multiplex 创建并返回整合消息 channel 和控制并发的 wg
func multiplex(messageChan ...chan string) (chan string, *sync.WaitGroup) {
	mChan := make(chan string)
	wg := &sync.WaitGroup{} // 这里可以注意跟var wg sync.WaitGroup的区别！

	for _, message := range messageChan {
		wg.Add(1)

		go func(m chan string, wg *sync.WaitGroup) {
			defer wg.Done()

			for message := range m {
				mChan <- message

			}

		}(message, wg)
	}

	return mChan, wg

}
