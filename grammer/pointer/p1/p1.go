package main

import "fmt"

// 修改指针的值
func swap1(x, y *int) {
	temp := *x
	*x = *y
	*y = temp

}

func main() {
	var (
		a = 10
		b = 20
	)
	swap1(&a, &b)

	fmt.Println(a, b)
}
