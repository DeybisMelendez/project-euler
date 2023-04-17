package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n == 0 || n == 4 || n == 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var i int
	var num int = 2
	for true {
		if isPrime(num) {
			i++
		}
		if i == 10001 {
			fmt.Println(num)
			break
		}
		num++
	}
}
