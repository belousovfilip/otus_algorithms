package primes

import "testing"

func BenchmarkLow(b *testing.B) {
	v := 10_000
	runIterateBenchmark(b, v)
	runEratosthenesBenchmark(b, v)
}

func BenchmarkMiddle(b *testing.B) {
	v := 1_000_000
	runIterateBenchmark(b, v)
	runEratosthenesBenchmark(b, v)
}

func BenchmarkHigh(b *testing.B) {
	v := 100_000_000
	runIterateBenchmark(b, v)
	runEratosthenesBenchmark(b, v)
}

func runIterateBenchmark(b *testing.B, v int) {
	b.Run("Iterate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Iterate(v)
		}
	})
}

func runEratosthenesBenchmark(b *testing.B, v int) {
	b.Run("Eratosthenes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Eratosthenes(v)
		}
	})
}
