package main

func main() {

}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//type Element interface{}
//
//type List []Element
//
//type Person struct {
//	name string
//	age  int
//}
//
//func (p Person) String() string {
//	return "(name:" + p.name + "-age:" + strconv.Itoa(p.age) + "years)"
//}
//
//func main() {
//	list := make(List, 3)
//	list[0] = 1
//	list[1] = "hello"
//	list[2] = Person{"Deniss", 30}
//	for index, element := range list {
//		switch value := element.(type) { // element.(type) 只能在 switch 里边使用，外边还是要用 comma-ok
//		case int:
//			fmt.Printf("int %d %d", index, value)
//		case string:
//			fmt.Printf("str %d %d", index, value)
//		case Person:
//			fmt.Printf("person %d %d", index, value)
//		default:
//			fmt.Printf("list different type %d", index)
//		}
//	}
//}
