package main

import (
	"fmt"
	"time"
)

//使用channel信道，可以在协程之间传递消息。阻塞等待并发协程返回消息。
var ch = make(chan string, 10) // 创建大小为10的缓冲信道

/*
使用通道发送数据
1) 通道发送数据的格式	 通道变量 <- 值
2）发送将持续阻塞直到数据被接收

使用通道接收数据
1) 阻塞接收数据      data := <-ch
2) 非阻塞接收数据    data, ok := <-ch
3) 接收任意数据，忽略接收的数据    <-ch
*/

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url // 将 url 发送给信道
}

func main() {
	for i := 0; i < 10; i++ {
		go download("a.com/" + string(i+'0'))
	}
	for i := 0; i < 10; i++ {
		msg := <-ch // 等待信道返回消息
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}
