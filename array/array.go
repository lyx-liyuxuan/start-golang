package main

import "fmt"

func Array1() {
	var a1 [3]int
	var a2 [5]string
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
}

func Array2() {
	var a1 = [...]byte{'a', 'b', 'c', 'd'}
	fmt.Printf("a1: %c\n", a1)

}
