//Exercise 1.3: Experiment to measure the difference in running time between our
//potentially inefficient versions and the one that uses strings.Join. (Section 1.6
//illustrates part of the time package, and Section 11.4 shows how to write
//benchmark tests for systematic performance evaluation.)

package concat

import (
	"testing"
)

//go test -bench .
//BenchmarkEcho1-4        300000000                4.27 ns/op
//BenchmarkEcho2-4        200000000                6.23 ns/op
//BenchmarkEcho3-4        200000000                8.21 ns/op

var (
	args = []string{"go arg1 arg2 arg3 arg4 arg5 arg6"}
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(args)
	}
}

