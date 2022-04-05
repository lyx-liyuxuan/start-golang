package main

import (
	"errors"
	"fmt"
	"os"
)

func err1() {
	// os.Open 有2个返回值，第一个是 *File，第二个是 error
	_, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		//open file.txt: The system cannot find the file specified.
	}
}

// go的相对路径是相对于执行命令时的目录
func err2() {
	_, err := os.Open("../github/start-golang/error/error1/errtest.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("The file is exist.")
	}
}

func hello(name string) error {
	if len(name) == 0 {
		// errors.New()自定义返回的错误
		return errors.New("error: name is null")
	}
	fmt.Println("Hello,", name)
	return nil
}
func err3() {
	if err := hello("Alice"); err != nil {
		fmt.Println(err)
	}
}

func get(index int) int {
	arr := [3]int{2, 3, 4}
	return arr[index]
}
func err4() {
	fmt.Println(get(5))

}

func main() {
	// err1()
	// err2()
	// err3()
	err4()
}
