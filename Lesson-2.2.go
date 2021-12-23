package main

import (
	"fmt"
	"math"
)

func main() {
	var userInputSquare float64
	fmt.Println("Введите площадь круга: ")
	fmt.Scanln(&userInputSquare)
	radiusCircle := math.Sqrt(userInputSquare / math.Pi)
	fmt.Printf("Диаметр окружности равен: %.2f\n", radiusCircle*2)
	fmt.Printf("Длина окружности равна: %.2f\n", 2*math.Pi*radiusCircle)
}
