package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	var (
		a int = 1
		b int = 1
	)

	t.Log(a)

	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		temp := a
		a = b
		b = a + temp
	}
	fmt.Println()
}
