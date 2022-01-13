package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		//Отвечу почему сделал так.
		//Сначала вводим несколько чисел (число и Enter).
		//Затем, чтобы вывести отсортированный слайс, нажимаем еще раз Enter.
		//Не смог разобраться как прервать набор массива и вывести результат в консоль, сделал вот так.
		//Понимаю что нет вызова ошибок.
		if err != nil {
			Sort(inputNums)
			fmt.Println(inputNums)
			break
		}
		inputNums = append(inputNums, num)
	}
}
func Sort(arr []int64) {
	for i := 1; i < len(arr); i++ {
		a := arr[i]
		b := i
		for ; b >= 1 && arr[b-1] > a; b-- {
			arr[b] = arr[b-1]
		}
		arr[b] = a
	}
}
