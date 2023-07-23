package main

import (
	"fmt"
	"math"
)

/*
定义： func (r ReceiverType) funcName(parameters) (results)
虽然 method 名字一样，但是接受者不一样，method 就不一样
method 里可以访问接受者的字段
调用 method  通过 . 访问就像 stuct 里访问字段一样
*/

/* method: 带有接受者的函数。method 附属在一个给定的类型上，语法和函数的声明语法几乎一样，只是在函数后加上一个 receiver

A method is a funciton with am inplicit first argument, called a reciver - Rob Pike*/
type Rectangle struct {
	width, high float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 { //调用者 (r Rectangle) Receiver
	return r.width * r.high
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) setHeight(height float64) {
	r.high = height
}
func main() {
	r1 := Rectangle{12, 1}
	c1 := Circle{10}
	fmt.Println(r1, c1)

}
