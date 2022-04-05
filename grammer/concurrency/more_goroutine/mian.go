package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 30; i++ {
		// time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'z'; i++ {
		// time.Sleep(250 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	go numbers()
	go alphabets()
	// 主程序休息一会，等go创建协程
	time.Sleep(1 * time.Millisecond)
	fmt.Println("main terminated")
}
