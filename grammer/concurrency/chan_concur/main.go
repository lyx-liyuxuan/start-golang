package main

import "fmt"

func main() {
	// 创建一个通道
	ch := make(chan int)

	// 创建一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		//通过通道发送数据
		ch <- 0

		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")
	// 等待匿名goroutine，阻塞接收数据
	<-ch

	fmt.Println("all done")
}
