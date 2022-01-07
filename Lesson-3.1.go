package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, %): ")
	fmt.Scanln(&op)
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if a == 0 {
			fmt.Println("Деление на ноль")
			os.Exit(1)
		} else if b == 0 {
			fmt.Println("Деление на ноль")
			os.Exit(1)
		} else {
			res = a / b
		}
	case "%":
		res = math.Mod(a, b)

	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %f\n", res)
}
