package primes

import (
	"math"
)

func Iterate(n int) int {
	isPrime := func(p int) bool {
		if p == 2 {
			return true
		}
		if p%2 == 0 {
			return false
		}
		sqrtP := int(math.Sqrt(float64(p)))
		for j := 3; j <= sqrtP; j += 2 {
			if p%j == 0 {
				return false
			}
		}
		return true
	}
	var count int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func Eratosthenes(n int) int {
	count := 0
	divs := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if divs[i] {
			continue
		}
		count++
		for j := i + i; j <= n; j += i {
			divs[j] = true
		}
	}
	return count
}
