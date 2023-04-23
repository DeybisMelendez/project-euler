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
	var n int = 2
	var total int
	for {
		if isPrime(n) {
			total += n
		}
		if n >= 2000000 {
			break
		}
		n++
	}
	fmt.Println(total)
}
