package main

import "fmt"

func febonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go febonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
