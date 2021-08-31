package sample

import "testing"

func Sum(a, b int) int {
	return a + b
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(123, 123)
	}
}

func BenchmarkSumSub(b *testing.B) {
	var bms = []struct {
		name string
		a    int
		b    int
	}{
		{"s", 1, 1},
		{"m", 123, 123},
		{"l", 1000000, 1000000},
		{"x", 9223372036854775806, 1},
	}
	for _, bm := range bms {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sum(bm.a, bm.b)
			}
		})
	}
}

// $ go test -bench .
// goos: linux
// goarch: amd64
// pkg: github.com/miku/goexp/Lifecycle/Benchmarks
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkSum-8          1000000000               0.2662 ns/op
// PASS
// ok      github.com/miku/goexp/Lifecycle/Benchmarks      0.302s
//
// goos: linux
// goarch: amd64
// pkg: github.com/miku/goexp/Lifecycle/Benchmarks
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkSum-8          1000000000               0.4200 ns/op
// BenchmarkSumSub/s-8     1000000000               0.3953 ns/op
// BenchmarkSumSub/m-8     1000000000               0.3993 ns/op
// BenchmarkSumSub/l-8     1000000000               0.3898 ns/op
// BenchmarkSumSub/x-8     1000000000               0.3706 ns/op
// PASS
// ok      github.com/miku/goexp/Lifecycle/Benchmarks      2.254s
