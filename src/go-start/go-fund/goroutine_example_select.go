package main

import "fmt"

func febonaci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			//监听的channel没有准备好时默认执行(select 不再阻塞等待channel)
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	//TODO:具体执行流程
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 100
	}()
	febonaci2(c, quit)
}
