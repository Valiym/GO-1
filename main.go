package main

import "fmt"

func FibSlow(n int) uint64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return FibSlow(n-1) + FibSlow(n-2)
}

var cache = make(map[int]uint64)

func FibFast(n int) (result uint64) {
	if result, ok := cache[n]; ok {
		return result
	}
	defer func() {
		cache[n] = result
	}()
	if n <= 1 {
		return uint64(n)
	}
	return FibFast(n-1) + FibFast(n-2)
}

func main() {
	var v int
	fmt.Print("Введите число : ")
	fmt.Scan(&v)
	fmt.Println(FibFast(v))
	fmt.Println(FibSlow(v))
}
