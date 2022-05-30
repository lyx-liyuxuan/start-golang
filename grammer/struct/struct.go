package main

import "fmt"

// 结构体
type Student struct {
	name string
	age  int
}

// 方法，和函数不同，在func和函数名之间加上调用方法的实例名及其类型
func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

func main() {
	stu := &Student{
		name: "Tom",
	}
	// 调用方法：实例名.方法名（参数）
	msg := stu.hello("Jack")
	fmt.Println(msg)
}
