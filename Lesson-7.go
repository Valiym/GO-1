package main

import (
	"fmt"
	"os"
)

var i, j float32
var op string
var err error

type Sum struct {
	First  float32
	Second float32
}

func (s *Sum) CalculateArea() float32 {
	return s.First + s.Second
}

type Sub struct {
	First  float32
	Second float32
}

func (s *Sub) CalculateArea() float32 {
	return s.First - s.Second
}

type Mult struct {
	First  float32
	Second float32
}

func (s *Mult) CalculateArea() float32 {
	return s.First * s.Second
}

type Div struct {
	First  float32
	Second float32
}

func (s *Div) CalculateArea() float32 {
	if s.Second == 0 {
		fmt.Printf("Ошибка - Деление на ноль")
		os.Exit(1)
	}
	return s.First / s.Second
}

type AreaCalculator interface {
	CalculateArea() float32
}

//Как по мне, код получился из сплошных костылей, уверен есть решение проще.
//С выбором знаков операции возникли сложности, пришлось опять использовать switch.
//По-хорошему, надо эту задачу разобрать на уроке, хотя-бы посмотреть как она должна выглядеть в идеале.

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
	sum := &Sum{First: i, Second: j}
	sub := &Sub{First: i, Second: j}
	mult := &Mult{First: i, Second: j}
	div := &Div{First: i, Second: j}

	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		var shape AreaCalculator = sum
		fmt.Println(shape.CalculateArea())
	case "-":
		var shape AreaCalculator = sub
		fmt.Println(shape.CalculateArea())
	case "*":
		var shape AreaCalculator = mult
		fmt.Println(shape.CalculateArea())
	case "/":
		var shape AreaCalculator = div
		fmt.Println(shape.CalculateArea())
	default:
		fmt.Println("Ошибка - Операция выбрана неверно")
		os.Exit(1)
	}
}
