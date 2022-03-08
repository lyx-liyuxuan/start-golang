package array

import "fmt"

func MyByte() {
	// 定义byte数组
	var a [3]byte = [3]byte{'a', 'b', 'c'}

	// for...range 循环遍历数组
	for _, val := range a {
		fmt.Printf("%c\n", val)
	}

	var m []string
	m = append(m, "你好")
	fmt.Println(len(m))
	m = append(m, "Golang") //[]不是切片类型，不能自动扩容
	m = append(m, "yeah")
	fmt.Printf("%s %s\n", m[0], m[1])
	fmt.Print(len(m))
}
