package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type example1 struct {
	Name string
}

type example2 struct {
	Name string
}

func (e1 *example1) Ping() {
	fmt.Println("Ping!")
}

func (e1 *example1) Pang() {
	fmt.Println("Pang")
}


func main() {
	var ii Inter
	ii = &example1{
		Name: "jiang",
	}
	fmt.Println(ii)

	var inter interface{} = ii
	t := inter.(Inter)
	t.Ping()
	t.Pang()

	s := t.(*example1)
	fmt.Println(s.Name)
}

