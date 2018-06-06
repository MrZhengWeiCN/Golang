package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	//无缓存是同步阻塞的，有缓存就是异步的
	//缓存就是chan的容量
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		Go(c)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
	//Select()
}

func Go(c chan bool) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(a)
	c <- true
}

//造成CPU死机
func Select() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	//利用select作为消息的发送者
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
