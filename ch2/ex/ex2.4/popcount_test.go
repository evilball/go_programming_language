//Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument
//through 64 bit positions, testing the rightmost bit each time. Compare its performance
//to the table-lookup version.

package popcount

import (
	"testing"
	"github.com/evilball/go_programming_language/ch2/ex/popcount"
)

//go test -bench .
//BenchmarkExpressionPopCount-4            2000000000               0.40 ns/
//BenchmarkLoopPopCount-4            50000000                27.4 ns/o
//BenchmarkShiftPopCount-4        10000000               180 ns/op

func BenchmarkExpressionPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.ExpressionPopCount(uint64(i))
	}
}

func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.LoopPopCount(uint64(i))
	}
}

func BenchmarkShiftPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.ShiftPopCount(uint64(i))
	}
}