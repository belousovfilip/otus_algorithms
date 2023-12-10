package main

func LuckyTickets(n int64) int64 {
	if n == 1 {
		return 10
	}
	var next func(n int64) []int64
	next = func(n int64) []int64 {
		if n == 1 {
			out := make([]int64, 10)
			for i := 0; i < 10; i++ {
				out[i] = 1
			}
			return out
		}
		in := next(n - 1)
		out := make([]int64, (9*n)+1)
		for i1 := 0; i1 < 10; i1++ {
			for i2, v := range in {
				out[i1+i2] += v
			}
		}
		return out
	}
	sum := next(n)
	len := len(sum)
	last := sum[len/2]
	sum = sum[:len/2]
	var count int64
	for _, v := range sum {
		count += v * v
	}
	count = count * 2
	if len%2 != 0 {
		return count + last*last
	}
	return count
}
