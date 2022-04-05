package main

import (
	"fmt"
	"time"
)

//使用channel信道，可以在协程之间传递消息。阻塞等待并发协程返回消息。
var ch = make(chan string, 10) // 创建大小为10的缓冲信道

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url // 将 url 发送给信道
}

func main() {
	for i := 0; i < 10; i++ {
		go download("a.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}