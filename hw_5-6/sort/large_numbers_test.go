package sort

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func BenchmarkLargeNumbers(b *testing.B) {
	b.StopTimer()
	path := CreateLargeFileNumbers(10_000_000)
	file, _ := os.Open(path)
	var n, v int
	fmt.Fscan(file, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(file, &v)
		arr[i] = v
	}
	b.StartTimer()
	b.Run("counting", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			in := make([]int, n)
			copy(in, arr)
			b.StartTimer()
			Counting(in)
		}
	})
	b.Run("bucket", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			in := make([]int, n)
			copy(in, arr)
			b.StartTimer()
			Bucket(in)
		}
	})
	b.Run("radix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			in := make([]int, n)
			copy(in, arr)
			b.StartTimer()
			Radix(in)
		}
	})
}

func CreateLargeFileNumbers(n int) string {
	file, err := os.CreateTemp("", "otus_large_numbers_")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(strconv.Itoa(n) + "\n"))
	for i := 0; i < n; i++ {
		file.Write([]byte(strconv.Itoa(rand.Intn(65535)) + " "))
	}
	return file.Name()
}
