package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	var err error
	fmt.Print("Введите первое число: ")
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Printf("Ошибка - Введите число")
		return
	}
	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Printf("Ошибка - Введите число")
		return
	}
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
		if b == 0 {
			fmt.Println("Ошибка - Деление на ноль")
			os.Exit(1)
		}
		res = a / b

	case "%":
		res = math.Mod(a, b)

	default:
		fmt.Println("Ошибка - Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %f\n", res)
	return
}
