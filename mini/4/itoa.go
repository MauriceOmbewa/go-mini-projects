package main

import "fmt"

func main() {
	fmt.Println(Itoa(-1654383))
}

func Itoa(a int) string {
	var final string
	isNegative := false
	if a < 0 {
		final += "-"
		a = a * -1
		isNegative = true
	}

	for a/10 != 0 {
		final = string(a%10+'0') + final
		a = a / 10
	}
	final = string(a+'0') + final
	if isNegative{
		final = final[len(final)-1:] + final[:len(final)-1]
	}
	return final
}
