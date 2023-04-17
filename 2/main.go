package main

import "fmt"

var fiboMap = make(map[uint64]uint64)

func fibo(n uint64) uint64 {
	if num, ok := fiboMap[n]; ok {
		return num
	}
	if n <= 1 {
		return n
	}
	var result uint64 = fibo(n-1) + fibo(n-2)
	fiboMap[n] = result
	return result
}

func main() {
	var num uint64
	var total uint64
	for true {
		var fibo uint64 = fibo(num)
		if fibo <= 4000000 {
			if fibo%2 == 0 {
				total += fibo
			}
			num++
		} else {
			break
		}
	}
	fmt.Println(total)
}
