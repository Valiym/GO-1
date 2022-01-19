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
		if err != nil {
			panic(err)
		}
		inputNums = append(inputNums, num)
	}
	Sort(inputNums)
	fmt.Println(inputNums)
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
