package main

import (
	"fmt"
)

func main() {
	arr := []int{8, 5, 6, 4, 3, 8, 9, 1, 15, 51, 10}
	Sort(arr)
	fmt.Println(arr)
}

func Sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		a := arr[i]
		b := i
		for ; b >= 1 && arr[b-1] > a; b-- {
			arr[b] = arr[b-1]
		}
		arr[b] = a
	}
}
