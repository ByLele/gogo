package main

import "fmt"

//TODO:
/*interface 是一组 method 签名的组合，通过 interface 来定义对象的一组行为*/
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Emploee struct {
	Human
	company string
	money   float32
}

func (h *Human) SayHi() {
	fmt.Printf("Hi %s %s", h.name, h.phone)
}

func (h *Human) Sing(lyrics string) {
	fmt.Print("la la", lyrics)
}

func (h *Human) Guzzle(beerstein string) {
	fmt.Println("Guzzle", beerstein)
}
func (e *Emploee) SayHi() {
	fmt.Printf("Hi %s %s %s", e.name, e.company, e.phone)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}
func (e *Emploee) SpendSalary(amount float32) {
	e.money -= amount
}

/*
interface 可以被任意对象实现，一个对象也可以实现任意多个 interface。 任意类型都是现了空 interface (interface{})，也就是包含了0个 method 的interface。
interface 值：如果定义一个interface的变量，这个变量里可以存实现这个 interface 的任意类型的对象。 比如定义了一个 Men interface 类型的变量 m，
m 可以存 Human, Student, Employee。 因为 m 可以持有三种对象，可以定一个包含 Men 类型的 slice，它可以倍赋予实现了 Men 接口的任意结构的对象。 通过这种方式实现了鸭子类型
*/
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}
type ElderyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

/*
空 interface: 可以存储任意类型数值，类似于 c void*。
一个函数把 interface{} 作为参数，可以接受任意类型。同理也可以返回任意类型
*/
func func_void_interface() {
	var a interface{}
	var i int = 5
	s := "hello"
	a = i
	a = s
	fmt.Println(a)
}
func main() {
	func_void_interface()
}
