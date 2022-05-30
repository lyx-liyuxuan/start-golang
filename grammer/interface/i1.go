package main

import "fmt"

// 声明接口
type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func test1() {
	// Student实例转换为Person接口
	var p Person = &Student{
		name: "Bob",
		age:  18,
	}

	fmt.Println(p.getName())
	fmt.Printf("p: %v\n", p)
}

func test2() {
	var p Person = &Student{
		name: "Tom",
		age:  25,
	}

	// 接口转为示例
	stu := p.(*Student)
	fmt.Println(stu.getName())
}

func main() {
	// test1()
	test2()
}
