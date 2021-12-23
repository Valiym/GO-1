package main

import "fmt"

func main() {
	var userInputLength float64
	var userInputWidth float64
	fmt.Println("Введите длину и ширину прямоугольника через пробел: ")
	fmt.Scanln(&userInputLength, &userInputWidth)
	fmt.Printf("Площадь прямоугольника равна: %.2f\n", userInputLength*userInputWidth)
}
