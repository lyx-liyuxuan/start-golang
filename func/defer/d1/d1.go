package main

import "fmt"

func main() {
	fmt.Println("test defer")

	// 将defer放入延迟调用栈，先进后出
	// defer是在函数推出后执行的
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
}
