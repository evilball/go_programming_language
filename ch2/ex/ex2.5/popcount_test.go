//Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x. Write
//a version of PopCount that counts bits by using this fact, and assess its
//performance.

package popcount

import (
	"testing"
	"github.com/evilball/go_programming_language/ch2/ex/popcount"
)

//go test -bench .
//BenchmarkExpressionPopCount-4                    2000000000               0.34 ns/o
//BenchmarkLoopPopCount-4                    50000000                25.8 ns/op
//BenchmarkShiftPopCount-4                10000000               146 ns/op
//BenchmarkClearRightPopCount-4           100000000               11.8 ns/op

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

func BenchmarkClearRightPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.ClearRightPopCount(uint64(i))
	}
}