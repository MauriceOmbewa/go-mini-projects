package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	first := os.Args[1]

	firstconv, err := strconv.Atoi(first)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i := 0; i < 99999999999; i++ {
		if firstconv == int(math.Pow(2, float64(i))) {
			fmt.Println("true")
			return
		} else {
			fmt.Println("false")
			return
		}
	}
}
