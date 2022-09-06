package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
	canWrite() bool
}

type file struct {
	CanWrite bool

}

func (f *file) WriteData(data interface{}) error {

	fmt.Println("把data写入file中", data)
	return nil

}

func (f *file) canWrite() bool {
	if !f.CanWrite {
		return true
	}
	return false

}

func main() {
	f := &file{
		CanWrite: true,
	}
	var i DataWriter
	i = f
	// 调用方法
	if err := i.WriteData("data"); err != nil {
		panic(err)
	}


}
