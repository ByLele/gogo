package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //cpu让出时间片，下个时间继续恢复执行goroutine
		fmt.Println(s)
	}
}
func main() {
	go say("hello")
	say("hello")
}
