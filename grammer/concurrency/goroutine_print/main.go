package main

import (
	"fmt"
)

/*
goroutine
- goroutine 是一种非常轻量级的实现，可在单个进程里执行成千上万的并发任务,它是Go语言并发设计的核心。
- goroutine 其实就是线程，但是它比线程更小，十几个 goroutine 可能体现在底层就是五六个线程
- 使用 go 关键字就可以创建 goroutine
- 开启一个goroutine的消耗非常小（大约2KB的内存）,启动时间比线程快

channel
使用通道发送数据
1) 通道发送数据的格式	 通道变量 <- 值
2）发送将持续阻塞直到数据被接收

使用通道接收数据
1) 阻塞接收数据      data := <-ch
2) 非阻塞接收数据    data, ok := <-ch
3) 接收任意数据，忽略接收的数据    <-ch
*/

func printer(c chan int) {

	// 开始无限循环等待数据
	for {

		// 从channel中获取一个数据
		data := <-c

		// 将0视为数据结束
		if data == 0 {
			break
		}

		// 打印数据
		fmt.Println(data)
	}

	// 通知main已经结束循环(我搞定了!)
	c <- 0

}

func main() {

	// 创建一个channel
	c := make(chan int)

	// 并发执行printer, 传入channel
	go printer(c)

	for i := 1; i <= 10; i++ {

		// 将数据通过channel投送给printer
		c <- i
	}

	// 通知并发的printer结束循环(没数据啦!)
	c <- 0

	// 等待printer结束(搞定喊我!)
	<-c

}
