package main

import (
	"fmt"
)

var cache = make(map[uint]uint64)

func fib(n uint) (result uint64) {
	if result, ok := cache[n]; ok {
		return result
	}
	defer func() {
		cache[n] = result
	}()
	if n <= 1 {
		return uint64(n)
	}
	return fib(n-1) + fib(n-2)
}
func main() {
	var v uint
	fmt.Print("Введите число : ")
	fmt.Scan(&v)
	fmt.Println(fib(v))

}
