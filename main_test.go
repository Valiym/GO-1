package main_test

import (
	"Lesson-10.3"
	"fmt"
)

func ExampleSort() {
	s := []int64{6, 0, 2, 4, 5, 3, 8, 9, 1, 7}
	fmt.Println(main.Sort(s))
	// Output: [0 1 2 3 4 5 6 7 8 9]
}
