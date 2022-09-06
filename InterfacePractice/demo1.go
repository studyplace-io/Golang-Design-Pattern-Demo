package main

import (
	"fmt"
	"io"
	"os"
)

// 接口声明方法
// 第一种：直接声明
type Reader interface {
	Read(p [] byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Read(p [] byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 第二种：里面嵌套接口
type ReadWriter1 interface {
	Reader
	Writer
}

// 第三种：有嵌套接口也有方法
type ReadWriter2 interface {
	Reader
	Write(p []byte) (n int, err error)
}


func main() {
	// 例一：接口初始化
	var i io.Reader // 只有声明 接口变量为nil
	fmt.Printf("%T\n", i)

	file, _ := os.OpenFile("xxx.txt", os.O_RDWR | os.O_CREATE, 0755)
	var rw io.ReadWriter = file
	var w io.Writer = rw
	fmt.Printf("%T\n", w)


}


