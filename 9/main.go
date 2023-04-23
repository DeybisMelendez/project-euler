package main

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Pythagorean_triple

// Euclid's formula
func main() {
	var a, b, c int
	var found bool
	for m := (1); m < 1000; m++ {
		for n := (1); n < m; n++ {

			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			if a+b+c == 1000 && a*a+b*b == c*c {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if found {
		fmt.Printf("found: a: %d, b: %d, c: %d, a+b+c: %d, solution: %d.\n", a, b, c, a+b+c, a*b*c)
	} else {
		fmt.Println("not found")
	}
}
