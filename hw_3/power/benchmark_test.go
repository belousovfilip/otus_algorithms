package power

import "testing"

func BenchmarkLow(b *testing.B) {
	var v float64 = 10_000
	runIterateBenchmark(b, 1.0000000001, v)
	runMultiplicationBenchmark(b, 1.0000000001, v)
	runBinaryBenchmark(b, 1.0000000001, v)
}

func BenchmarkLow1(b *testing.B) {
	var v float64 = 6_800
	runIterateBenchmark(b, 1.0000000001, v)
	runMultiplicationBenchmark(b, 1.0000000001, v)
	runBinaryBenchmark(b, 1.0000000001, v)
}

func BenchmarkMiddle(b *testing.B) {
	var v float64 = 10_000_000
	runIterateBenchmark(b, 1.0000000001, v)
	runMultiplicationBenchmark(b, 1.0000000001, v)
	runBinaryBenchmark(b, 1.0000000001, v)
}

func BenchmarkHigh(b *testing.B) {
	var v float64 = 10_000_000_000
	runIterateBenchmark(b, 1.0000000001, v)
	runMultiplicationBenchmark(b, 1.0000000001, v)
	runBinaryBenchmark(b, 1.0000000001, v)
}

func runIterateBenchmark(b *testing.B, a, n float64) {
	b.Run("Iterate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Iterate(a, n)
		}
	})
}
func runMultiplicationBenchmark(b *testing.B, a, n float64) {
	b.Run("Multiplication", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Multiplication(a, n)
		}
	})
}
func runBinaryBenchmark(b *testing.B, a, n float64) {
	b.Run("Binary", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Binary(a, n)
		}
	})
}
