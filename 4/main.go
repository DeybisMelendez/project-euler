package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
func isPalindrome(str string) bool {
	return str == reverse(str)
}

func main() {
	var palindrome int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			var val int = i * j
			var number string = strconv.Itoa(val)
			if isPalindrome(number) {
				if val > palindrome {
					palindrome = val
				}
			}
		}
	}
	fmt.Println(palindrome)
}
