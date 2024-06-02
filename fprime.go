package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2{
		return
	}

	input := os.Args[1]
	var final []int

	inpconv, err := strconv.Atoi(input)
	if err != nil{
		return
	}

	for i := 2; i <= inpconv; i++{
		if inpconv%i == 0{
			inpconv /= i
			final = append(final, i)
			if inpconv%i == 0{
				inpconv /= i
				final = append(final, i)
			}
		}
	}
	if len(final) < 1{
		return
	}

	for i := 0; i < len(final); i++{
		if Prime(final[i]){
			fmt.Print(final[i])
			if i != len(final)-1 && Prime(final[i+1]) == true{
				fmt.Print("*")
			}
		}
	}
	fmt.Println()
}

func Prime(a int) bool{
	if a == 1{
		return false
	} else if a == 2 {
		return true
	}
	for i := 2; i < a; i++{
		if a%i == 0{
			return false
		}
	}
	return true
}