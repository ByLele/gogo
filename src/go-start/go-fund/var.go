package main

import (
	"debug/pe"
	"errors"
	"fmt"
)
//go 基础变量

func _func_var_() {
	var var_name type//type类型可以为 ...
	var name1,name2,name3,type
	var a,b,c= 1,2,3
	A,B,C := 1,2,3
	_,b = 1,2

}
func _func_const_() {
	const name = "aaa"
	const Pi float32 = 3.14
	const i = 100
}

func _func_default_var_() {
	var isActive = bool
	var enabe,disable = true,false
	vaild := false
	var hello string
	var emptyString = ""
	s := "hello"
	c := []byte(s) //把字符串转换为[]byte 类型#TODO: 转换数组类型  byte 是什么
	c[0] = 'c'
	s2 := string(c) //转换为string
	fmt.Printf("%s\n",s2)
	m :=  `hello
			world` //括号里的是 Raw string #TODO: raw string是什么类型
	fmt.Printf(m)
	//错误类型 error type
	err := errors.New("emit macho swarf: elf header corrupted")
	if err != nil {
		fmt.print(err)
	}

}
func _func_arr_ (){
	/*import (
		"fmt"
		"os"
	)*/
	const(
		i = 100
		pi = 3.14
	)
	var(
		i int
		pi float32
	)
	const(
		x = itoa //x == 0
		y = ioto //y == 1
	)
}

func main() {
	fmt.Println("1234")
}
