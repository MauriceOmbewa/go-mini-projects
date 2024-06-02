package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args) != 2{
		return
	}

	input := os.Args[1]

	if len(input) == 0{
		z01.PrintRune('\n')
		return
	}

	for k, char := range input{
		if (char >= 'a' && char <= 'z'){
			value := char-'a'
			for i :=0; i < int(value);i++{
				z01.PrintRune(char)
			}
		} else if (char >= 'A' && char <= 'Z'){
			value := char-'A'
			for i :=0; i < int(value);i++{
				z01.PrintRune(char)
			}
		}
		z01.PrintRune(char)
		if k == len(input)-1{
			z01.PrintRune('\n')
		}
	}
}