package main

import (
	"fmt"
	"math"
)

type Square struct {
	long float32
	short float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() (float32, error)
}

func (sq *Square) Area() (float32, error) {
	return sq.long * sq.short, nil
}

func (c *Circle) Area() (float32, error) {
	return c.radius * c.radius * math.Pi, nil
}


func main() {
	// 接口变量
	var areaIntface Shaper


	sq := &Square{
		long: 5,
		short: 5,
	}
	cr := &Circle{
		radius: 3,
	}
	res, _ := sq.Area()

	areaIntface = sq
	fmt.Println("Square面积：", res)
	res2, _ := cr.Area()
	fmt.Println("Circle面积：", res2)

	// 类型断言！
	if t, ok := areaIntface.(*Square); ok {
		fmt.Printf("areaIntface类型：%v", t)
	} else if u, ok := areaIntface.(*Circle); ok {
		fmt.Printf("areaIntface类型：%v", u)
	} else {
		fmt.Println("areaIntface不能")
	}

}
