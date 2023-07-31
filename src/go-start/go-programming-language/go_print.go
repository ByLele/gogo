package main

import (
	"fmt"
	"os"
	"strings"
)

func _func_print_() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
}

func _func_print2_() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}
}
func _func_print3_() {
	fmt.Println(strings.Join(os.Args[1:], ""))
}
func main() {
	_func_print_()
}
