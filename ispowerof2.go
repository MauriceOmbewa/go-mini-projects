package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args) != 2{
		return
	}

	input := os.Args[1]
	inpconv, _:= strconv.Atoi(input)

	var new bool

	for inpconv > 0{
		if inpconv%2 == 1{
			new = false
			break
		} else if inpconv/ 2 == 1 {
			new = true
			break
		}
		inpconv = inpconv / 2
	}

	if new == true{
		z01.PrintRune('t')
		z01.PrintRune('r')
		z01.PrintRune('u')
		z01.PrintRune('e')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('f')
		z01.PrintRune('a')
		z01.PrintRune('l')
		z01.PrintRune('s')
		z01.PrintRune('e')
		z01.PrintRune('\n')
	}
}