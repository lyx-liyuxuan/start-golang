package main

import (
	"fmt"
)

/*
   break        default      func         interface    select
   case         defer        go           map          struct
   chan         else         goto         package      switch
   const        fallthrough  if           range        type
   continue     for          import       return       var

    Constants:    true  false  iota  nil

        Types:    int  int8  int16  int32  int64
                  uint  uint8  uint16  uint32  uint64  uintptr
                  float32  float64  complex128  complex64
                  bool  byte  rune  string  error

    Functions:   make  len  cap  new  append  copy  close  delete
                 complex  real  imag
                 panic  recover
*/

var name string
var age int
var isOk bool

var (
	a string
	b int64
	c bool
	d float32
)

const (
	n1 = 100
	n2
	n3
)

func foo() (int, string) {
	return 10, "Q1mi"
}
func unname() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}

func array() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]

	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}

func slice() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10

	b := []int{}
	b = append(b, 1, 2, 3, 4)
	fmt.Println(len(b), cap(b))
	b = append(b, 5, 6, 7)
	fmt.Println(len(b), cap(b))

	c := b
	fmt.Println(c)
	fmt.Println(b)
	c[0] = 1000
	fmt.Println(c)
	fmt.Println(b)

	d := make([]int, len(c), cap(c))
	copy(d, c)
	fmt.Println(d)

	a = []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}

func TestMap() {
	HashMap := make(map[int]int) // 可以定义容量
	a := []int{0, 1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4}
	for _, v := range a {
		if _, ok := HashMap[v]; ok {
			HashMap[v]++
		} else {
			HashMap[v] = 1
		}
	}
	for key, _ := range HashMap {
		fmt.Println(key, HashMap[key])
	}
}

func main() {
	//fmt.Println(name, age, isOk, a, b, c, d)
	//
	//m, n := 10, 100
	//fmt.Println(m, n)

	//unname() // 测试匿名遍量
	//
	//array() // 测试数组

	//slice() // 测试切片
	//
	TestMap() // 测试map
}
