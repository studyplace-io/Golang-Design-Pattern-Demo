package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("xxx.txt", os.O_RDWR | os.O_CREATE, 0775)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var i io.ReadWriter
	i = f
	// 用switch case 看一下接口对象类型
	switch v := i.(type) {
	case *os.File:
		_, _ = v.Write([]byte("io.ReadWriter\n"))
		_ = v.Sync()
	case io.ReadWriter:
		_, _ = v.Write([]byte("io.ReadWriter\n"))
	default:
		fmt.Println("接口类型错误")
		return
	}

}
