package main

import (
	"fmt"
	"reflect"
)

// type RunMyString struct {
// }

func Run() {
	// byte = uint8 一个字节的ascii
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
	fmt.Println("\n你好")
}

func string1() {
	str1 := "Golang"
	str2 := "Go语言"

	// reflect.TypeOf().Kind()可以知道某个变量的类型
	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("str2[2]: %d  %c\n", str2[2], str2[2])
	fmt.Printf("len(str2): %v\n", len(str2)) // 字符串以byte数组类型存储
}

func main() {
	// Run()
	string1()
}
