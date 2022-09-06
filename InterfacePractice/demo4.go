package main

import "fmt"

type Inter1 interface {
	Ping1()
	Pang1()
}

type Anter1 interface {
	Inter1
	String1()
}

type example11 struct {
	Name string
}

type example21 struct {
	Name string
}

func (e1 *example11) Ping1() {
	fmt.Println("Ping!")
}

func (e1 *example11) Pang1() {
	fmt.Println("Pang!")
}


func main() {
	e1 := &example11{
		Name: "jiangjiang",
	}
	var inter interface{} = e1
	if t, ok := inter.(Inter1); ok {
		t.Pang1()
		t.Ping1()
	}

	if p, ok := inter.(Anter1); ok {
		p.String1()
	}
	if s, ok := inter.(*example11); ok {
		fmt.Println("name: ", s.Name)
	}

}

