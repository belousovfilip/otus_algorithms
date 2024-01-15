package fibonacci

import "math"

func Recursion(n int64) int64 {
	if n <= 1 {
		return n
	}
	return Recursion(n-1) + Recursion(n-2)
}

func Iterate(n int64) int64 {
	if n <= 1 {
		return n
	}
	var a, b int64 = 1, 1
	for i := int64(2); i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func Bine(in int64) int64 {
	n := float64(in)
	r := math.Pow((1+math.Sqrt(5))/2, n)
	r -= math.Pow((1-math.Sqrt(5))/2, n)
	r /= math.Sqrt(5)
	return int64(math.Round(r))
}
