package main_test

import (
	"Lesson-10.2"
	"fmt"
	"testing"
)

func TestSumTable(t *testing.T) {
	type call struct {
		Op       string
		A, B     float64
		Expected float64
	}

	calls := []call{
		//Sum
		{"+", 1, 2, 3},
		{"+", 10, 50, 60},
		{"+", 1000, 500, 1500},
		{"+", 10000, 5000, 15000},

		//Sub
		{"-", 1, 2, -1},
		{"-", 10, 50, -40},
		{"-", 1000, 500, 500},
		{"-", 10000, 5000, 5000},

		//Mult
		{"*", 1, 2, 2},
		{"*", 10, 50, 500},
		{"*", 1000, 500, 500000},
		{"*", 10000, 5000, 50000000},

		//Div
		{"/", 1, 2, 0.5},
		{"/", 10, 50, 0.2},
		{"/", 1000, 500, 2},
		{"/", 10000, 5000, 2},
	}

	for _, c := range calls {
		fnCallStr := fmt.Sprintf("Sum(%.2f %s %.2f ) ", c.A, c.Op, c.B)

		result := main.Sum(c.A, c.B, c.Op)
		var err error
		if err != nil {
			t.Fatalf("%s expected to return no error, but received %s", fnCallStr, err)
		}
		if result != c.Expected {
			t.Fatalf("%s expected to return %.2f, but received %.2f", fnCallStr, c.Expected, result)
		}
	}

}
