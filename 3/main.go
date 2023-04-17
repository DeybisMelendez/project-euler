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
	var input int = 600851475143
	var maxFactor int
	var prime int = 3
	for true {
		if input%prime == 0 {
			if prime > maxFactor {
				maxFactor = prime
			}
			input /= prime
		}
		prime += 2
		for true {
			if isPrime(prime) {
				break
			} else {
				prime++
			}
		}
		if input == 1 {
			break
		}
	}
	fmt.Println(maxFactor)
}
