package main

import "fmt"

func slice1() {
	// 数组的定义，长度不能改变
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr1: %v\n", arr1)

	// 切片的定义，可以扩容
	s1 := make([]int, 5, 10)
	fmt.Printf("s1: %v\n", s1)
	fmt.Println(len(s1), cap(s1))

	// 添加超过切片容量cap，自动扩容且为2倍扩容
	s1 = append(s1, 1, 2, 3, 4, 5, 6)
	fmt.Printf("s1: %v\n", s1)
	fmt.Println(len(s1), cap(s1))

	// 子切片
	s2 := s1[2:]
	s3 := s1[1:6]
	// 合并切片
	s4 := append(s2, s3...)
	fmt.Printf("s4: %v\n", s4)
}

func main() {
	slice1()
}
