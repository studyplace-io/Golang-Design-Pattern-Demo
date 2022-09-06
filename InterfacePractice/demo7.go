package main

type Socket struct {

}

func (s *Socket) Write(data interface{}) (n int, err error) {
	return 0, nil
}

func (s *Socket) Close() error {
	return nil
}

// 分别实现三个接口。

type IWrite interface {
	Write(interface{}) (int, error)
}

type IClose interface {
	Close() error
}

type IWriteClose interface {
	Write(interface{}) (int, error)
	Close() error
}

func usingWriter(writer IWriteClose) {
	writer.Write(nil)
}

func usingCloser(closer IWriteClose) {
	closer.Close()
}

func main() {
	// 对象
	s := &Socket{}
	// 接口对象复值
	var iw IWrite
	iw = s
	iw.Write("")
	var ic IClose
	ic = s
	ic.Close()
	var iwc IWriteClose
	iwc = s
	iwc.Write("")
	iwc.Close()

	usingWriter(s)
	usingCloser(s)

}