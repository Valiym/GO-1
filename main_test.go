package main

import (
	"testing"
)

var resultSlow uint64

func BenchmarkFibSlow(b *testing.B) {
	var res uint64
	for i := 0; i < b.N; i++ {
		res = FibSlow(43)
	}
	resultSlow = res
}

var resultFast uint64

func BenchmarkFibFast(b *testing.B) {
	var res uint64
	for i := 0; i < b.N; i++ {
		res = FibFast(43)
	}
	resultFast = res
}

//Результат теста:
//BenchmarkFibSlow
//BenchmarkFibSlow-6             1        1963052400 ns/op
//BenchmarkFibFast
//BenchmarkFibFast-6      137092140                8.721 ns/op
//PASS
