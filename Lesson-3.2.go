package main

import "fmt"

func main() {
	var number, j, i, flag int

	fmt.Print("Введите число : ")
	fmt.Scan(&number)
	for j = 2; j <= number; j++ {
		flag = 0
		for i = 2; i <= j/i; i++ {
			if j%i == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 {
			fmt.Println(j)
		}
	}
}
