package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 并发下载N个资源
func download(url string) {
	fmt.Println("start to download", url)
	// time.Sleep(time.Second) // 模拟耗时操作
	wg.Done()
}

func main1() {
	for i := 0; i < 100; i++ {
		// 为wg添加一个计数，wg.Done()减去一个计数
		wg.Add(1)
		// 启动新的协程并发执行download函数
		go download("a.com/" + string(i+'0'))
	}
	// 等待所有协程执行结束
	wg.Wait()
	fmt.Println("Done!")
}

func download2(url string) {
	fmt.Println("start to download", url)
}

func main2() {
	for i := 0; i < 100; i++ {
		download2("b.com/" + string(i+'0'))
	}
}

func main() {
	// main1()
	main2()
}
