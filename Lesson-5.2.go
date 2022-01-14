package main

import "fmt"

func fib(n int) int {
	cash := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = cash[i-1] + cash[i-2]
		}
		cash[i] = f
	}
	return cash[n]
}

func main() {
	var v int
	fmt.Print("Введите число : ")
	fmt.Scan(&v)
	fmt.Println(fib(v))
}
