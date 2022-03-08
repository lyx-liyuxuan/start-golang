package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b
		c
	)

	// 观察第二个iota会不会延续上一个迭代器的赋值
	const (
		apple = iota
		banana
		orange
	)

	// 100 101 102 103 104 505 506 507
	const (
		m = iota + 100
		n
		x
		y
		z
		o = iota + 500
		p
		q
	)
	fmt.Println(a, b, c)
	fmt.Println(apple, banana, orange)
	fmt.Println(m, n, x, y, z, o, p, q)
}
