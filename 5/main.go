package main

import "fmt"

func main() {
	var num int = 21
	for true {
		var divisible bool = true
		for i := 1; i < 21; i++ {
			if num%i != 0 {
				divisible = false
				num++
				break
			}
		}
		if divisible {
			fmt.Println(num)
			break
		}
	}
}
