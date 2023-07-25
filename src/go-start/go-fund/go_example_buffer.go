package main

import "fmt"

func sum2(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func main() {
	a := []int{3, 1, 3, 5, 6, 7}
	c := make(chan int, 12)
	go sum2(a[:len(a)/2], c)
	go sum2(a[len(a)/2:], c)
	x, y := <-c, <-c // 从 c 接收
	fmt.Println(x, y, x+y)
}
