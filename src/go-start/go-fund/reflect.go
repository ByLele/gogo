package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := "hello"
	var l interface{}
	t := reflect.TypeOf(i)
	v := reflect.TypeOf(l)
	fmt.Println(t, v)
}
