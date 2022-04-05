package main

import "fmt"

// go函数中参数是按值传递的,函数内部会拷贝一份参数
// 对参数的修改不会影响外部变量
func add(num int) {
	num += 1
}

// 函数参数使用指针,对参数的变化会传递到外部变量
func realAdd(num *int) {
	*num += 1
}

func main() {
	num := 100
	add(num)
	fmt.Println(num) // 100，num 没有变化

	realAdd(&num)
	fmt.Println(num) // 101，指针传递，num 被修改
}
