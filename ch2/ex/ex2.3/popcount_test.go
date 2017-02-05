//Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression.
//Compare the performance of the two versions. (Section 11.4 shows how to compare
//the performance of different implementations systematically.)

package popcount

import (
	"testing"
	"github.com/evilball/go_programming_language/ch2/ex/popcount"
)

//go test -bench .
//BenchmarkExpressionPopCount-4    2000000000               0.33 ns/op
//BenchmarkLoopPopCount-4    50000000                25.1 ns/op

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