package main

import ()

func main() {
	m := map[int]func(op int)int{}
	m1 := func(op int) int { return op }
	m2 := func(op int) int { return op * op }
	m3 := func(op int) int { return op * op * op}
}