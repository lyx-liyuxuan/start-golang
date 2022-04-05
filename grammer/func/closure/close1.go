package main

import "fmt"

func main1() {
	// 准备一个字符串
	str := "hello world"
	// 创建一个匿名函数
	foo := func() {

		// 匿名函数中访问str
		str = "hello dude"
	}
	// 调用匿名函数
	foo()
	fmt.Printf("str: %v\n", str)
}

func main() {
	main1()
}
