package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	first := os.Args[1]
	operator := os.Args[2]
	second := os.Args[3]
	var final int


	if (operator != "+" && operator != "-" && operator != "/" && operator != "*" && operator != "%") {
		fmt.Println("Error")
		return
	}

	firstconv, err := strconv.Atoi(first)
	if err != nil {
		return
	}
	secondconv, err := strconv.Atoi(second)
	if err != nil {
		return
	}

	if operator == "+" {
		final = firstconv + secondconv
	} else if operator == "-" {
		final = firstconv - secondconv
	} else if operator == "/" {
		if secondconv == 0 || firstconv == 0{
			fmt.Println("No division by 0")
			return
		}
		final = firstconv / secondconv
	} else if operator == "%" {
		if secondconv == 0 || firstconv == 0{
			fmt.Println("No modulo by 0")
			return
		}
		final = firstconv % secondconv
	}else if operator == "*"{
		final = firstconv * secondconv
	}
	fmt.Println(final)

}
