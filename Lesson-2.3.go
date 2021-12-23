package main

import "fmt"

func main() {
	var userNumber int64
	fmt.Println("Введите число: ")
	fmt.Scanln(&userNumber)
	fmt.Printf("Количество сотен: %d\n", (userNumber/100)%10)
	fmt.Printf("Количество десятков: %d\n", (userNumber/10)%10)
	fmt.Printf("Количество единиц: %d\n", userNumber%10)
}
