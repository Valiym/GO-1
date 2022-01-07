package main

import "fmt"

func main() {
	var number, a, b int
	var flag bool

	fmt.Print("Введите число : ")
	fmt.Scan(&number)
	for a = 2; a <= number; a++ {
		flag = false
		for b = 2; b <= a/2; b++ {
			if a%b == 0 {
				flag = true
			}
		}
		if flag == false {
			fmt.Println(a)
		}
	}
}
