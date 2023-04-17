package main

import "fmt"

func main() {
	var sumSquare int
	var squareSum int
	for i := 1; i < 101; i++ {
		sumSquare += i * i
		squareSum += i
	}
	squareSum *= squareSum
	var result int = squareSum - sumSquare
	fmt.Println(result)
}
