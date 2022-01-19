package main

import (
	"fmt"
	"math"
	"os"
)

var a, b float64
var err error
var op string

func sum(a, b float64, op string) (res float64) {
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
	return
}

func main() {
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

	fmt.Printf("Результат выполнения операции: %f\n", sum(a, b, op))
	return
}
