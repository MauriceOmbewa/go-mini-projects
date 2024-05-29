package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(10))
}

func FindPrevPrime(nb int) int {
	for i:=nb; i > 0; i--{
		if prime(i){
			return i
		}
	}
	return 0
}

func prime(n int) bool{
	for i := 2; i < n; i++{
		if n%i == 0{
			return false
		}
	}
	return true
}