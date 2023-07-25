package main

import "fmt"

/*
多个 goroutine 运行在同一个进程里边，共享内存数据。
不过设计上需要遵循：不要通过共享来通信，通过通信来共享。
*/

func chan_tuto() {
	/*channels：goroutine 共享相同的地址空间，访问共享内存需要做好同步。
	TODO:必须用 make 创建
	默认channel接收和发送数据都是阻塞的，
	除非另一端准备好，这样使得 goroutine 同步变得简单，无需显示 lock。*/
	//ci := make(chan int)
	//cs := make(chan string)
	//cf := make(chan interface{})
	//ch <- v   //发送数据v到 channels ch
	//v := <-ch //从ch中接受数据，赋值给v
	//fmt.Println(ci, cs, cf, ch, v)
}

func sum(a []int, c chan int) { // TODO:c Chan int 输出值的注意, 如何同步阻塞发生时机
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func main() {
	a := []int{3, 1, 3, 5, 6, 7}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
