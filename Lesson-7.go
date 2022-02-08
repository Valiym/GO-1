package main

import (
	"fmt"
	"os"
)

var i, j float32
var op string
var err error

type Sum struct {
	FirstNumber  float32
	SecondNumber float32
}

func (s *Sum) Calculate() float32 {
	return s.FirstNumber + s.SecondNumber
}

type Sub struct {
	FirstNumber  float32
	SecondNumber float32
}

func (s *Sub) Calculate() float32 {
	return s.FirstNumber - s.SecondNumber
}

type Mult struct {
	FirstNumber  float32
	SecondNumber float32
}

func (s *Mult) Calculate() float32 {
	return s.FirstNumber * s.SecondNumber
}

type Div struct {
	FirstNumber  float32
	SecondNumber float32
}

func (s *Div) Calculate() float32 {
	if s.SecondNumber == 0 {
		fmt.Printf("Ошибка - Деление на ноль")
	}
	return s.FirstNumber / s.SecondNumber
}

type Calculator interface {
	Calculate() float32
}

func main() {
	fmt.Print("Введите первое число: ")
	_, err = fmt.Scanln(&i)
	if err != nil {
		fmt.Printf("Ошибка - Введите число")
		return
	}
	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&j)
	if err != nil {
		fmt.Printf("Ошибка - Введите число")
		return
	}
	sum := &Sum{FirstNumber: i, SecondNumber: j}
	sub := &Sub{FirstNumber: i, SecondNumber: j}
	mult := &Mult{FirstNumber: i, SecondNumber: j}
	div := &Div{FirstNumber: i, SecondNumber: j}

	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	_, err = fmt.Scanln(&op)
	if err != nil {
		return
	}

	switch op {
	case "+":
		var shape Calculator = sum
		fmt.Println(shape.Calculate())
	case "-":
		var shape Calculator = sub
		fmt.Println(shape.Calculate())
	case "*":
		var shape Calculator = mult
		fmt.Println(shape.Calculate())
	case "/":
		var shape Calculator = div
		fmt.Println(shape.Calculate())
	default:
		fmt.Println("Ошибка - Операция выбрана неверно")
		os.Exit(1)
	}
}
