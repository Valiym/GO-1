package main

import "fmt"

func main() {
	var number, j, i int
	var flag bool

	fmt.Print("Введите число : ")
	fmt.Scan(&number)
	for j = 2; j <= number; j++ {
		flag = false
		for i = 2; i <= j/i; i++ {
			if j%i == 0 {
				flag = true
				break
			}
		}
		if flag == false {
			fmt.Println(j)
		}
	}
}
