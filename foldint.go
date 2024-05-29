package main

import "fmt"

// FoldInt applies the function f to each element in the slice a, starting with the accumulator n.
// MOST IMPORTANT(func FoldInt)
func FoldInt(f func(int, int) int, a []int, n int) {
	acc := n

	for _, val := range a {
		acc = f(acc, val)
	}
	fmt.Println(acc)
}

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Mul(x, y int) int {
	return x * y
}

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac) // Expected output: 99
	FoldInt(Mul, table, ac) // Expected output: 558
	FoldInt(Sub, table, ac) // Expected output: 87
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac) // Expected output: 93
	FoldInt(Mul, table, ac) // Expected output: 0
	FoldInt(Sub, table, ac) // Expected output: 93
}
