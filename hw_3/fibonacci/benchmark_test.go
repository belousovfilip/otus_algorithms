package fibonacci

import "testing"

func BenchmarkLow(b *testing.B) {
	var v int64 = 30
	runRecursionBenchmark(b, v)
	runIterateBenchmark(b, v)
	runBineBenchmark(b, v)
}

func BenchmarkLow1(b *testing.B) {
	var v int64 = 100
	runIterateBenchmark(b, v)
	runBineBenchmark(b, v)
}

func BenchmarkMiddle(b *testing.B) {
	var v int64 = 1_000_000
	runIterateBenchmark(b, v)
	runBineBenchmark(b, v)
}

func BenchmarkHigh(b *testing.B) {
	var v int64 = 10_000_000_000
	runIterateBenchmark(b, v)
	runBineBenchmark(b, v)
}

func runBineBenchmark(b *testing.B, v int64) {
	b.Run("Bine", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Bine(v)
		}
	})
}
func runRecursionBenchmark(b *testing.B, v int64) {
	b.Run("Recursion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Recursion(v)
		}
	})
}
func runIterateBenchmark(b *testing.B, v int64) {
	b.Run("Iterate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Iterate(v)
		}
	})
}
