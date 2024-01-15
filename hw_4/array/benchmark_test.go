package array

import "testing"

func BenchmarkAdd(b *testing.B) {
	RunBenchmarkAdd(b, "single", func() Array { return InitSingle() })
	RunBenchmarkAdd(b, "factor", func() Array { return InitFactor() })
	RunBenchmarkAdd(b, "vector", func() Array { return InitVector() })
	RunBenchmarkAdd(b, "slice", func() Array { return InitSlice() })
}

func BenchmarkRemove(b *testing.B) {
	RunBenchmarkRemoveMiddle(b, "single", func() Array { return InitSingle() })
	RunBenchmarkRemoveMiddle(b, "factor", func() Array { return InitFactor() })
	RunBenchmarkRemoveMiddle(b, "vector", func() Array { return InitVector() })
	RunBenchmarkRemoveMiddle(b, "slice", func() Array { return InitSlice() })
}

func RunBenchmarkAdd(b *testing.B, name string, init func() Array) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			arr := init()
			for j := 0; j < 100_000_000; j++ {
				arr.Add(j)
			}
		}
	})
}

func RunBenchmarkRemoveMiddle(b *testing.B, name string, init func() Array) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			arr := init()
			b.StopTimer()
			for j := 0; j < 200_000; j++ {
				arr.Add(j)
			}
			b.StartTimer()
			for j := 10; j < 100_000; j++ {
				arr.Remove(j)
			}
		}
	})
}
