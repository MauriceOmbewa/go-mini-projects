package main

import (
	"fmt"
)

func main() {
	num := Atoi("-54764123")
	fmt.Println(num)
}

func Atoi(a string) int {
	sign := 1
	result := 1

	if a[0] == '-' {
		sign *= -1
		a = a[1:]
	}

	for i, char := range a {
		if i == 0 {
			result *= int(char - '0')
		} else {
			result = (result * 10) + int(char-'0')
		}
	}

	return result * sign
}


