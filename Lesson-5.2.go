package main

import "fmt"

func fib(n int) int {
	cache := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = cache[i-1] + cache[i-2]
		}
		cache[i] = f
	}
	return cache[n]
}

func main() {
	var v int
	fmt.Print("Введите число : ")
	fmt.Scan(&v)
	fmt.Println(fib(v))
}
