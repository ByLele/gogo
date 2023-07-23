package main

import "fmt"

type Person struct {
	name string
	age  int
}

func Older(p1, p2 Person) (Person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var P Person
	P.name = "Alex"
	P.age = 23
	P1 := Person{"Tom", 12}
	P2 := Person{name: "auer", age: 10}
	p := new(Person)
	fmt.Println(P1, P2)
	fmt.Printf("name is", p.name)
}
