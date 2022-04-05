package main

import (
	"fmt"
	"time"
)

func main() {
	/* 匿名函数创建goroutine
		go func( 参数列表 ){
	    函数体
		}( 调用参数列表 )
	*/
	go func() {
		var times int
		for i := 0; i <= 10; i++ {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()
	var input string
	fmt.Scanln(&input)
}
