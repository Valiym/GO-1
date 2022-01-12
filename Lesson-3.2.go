package main

import "fmt"

func main() {
	var number, a, b, c int

	fmt.Print("Введите число : ")
	fmt.Scan(&number)
	for a = 2; a <= number; a++ {
		c = 0
		for b = 1; b <= a; b++ {
			if a%b == 0 {
				c++
			}
		}
		if c == 2 {
			fmt.Println(a)
		}
	}
}
