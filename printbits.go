package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	inputconv, err := strconv.Atoi(input)
	for err != nil{
		for i := 0; i < 8; i++{
			z01.PrintRune('0')
		}
		z01.PrintRune('\n')
		return
	}

	var new []int
	s := 0

	for inputconv > 0 {
		s = inputconv % 2
		new = append(new, s)
		inputconv = inputconv / 2
	}

	k := 0
	for i := len(new) - 1; i >= 0; i-- {
		if i == 0{
			k += new[i]
		} else{
			k += new[i]
			k = k * 10
		}
	}

	out := strconv.Itoa(k)
	for _, char := range out {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}


