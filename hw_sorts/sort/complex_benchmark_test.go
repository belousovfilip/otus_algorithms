package sort

import (
	"otus_algo/internal/factory"
	"strconv"
	"testing"
)

func Sort(arr []int) {
	//Bubble(arr)
	//Insert(arr)
	//InsertBinary(arr)
	//InsertShift(arr)
	//Shell(arr)
	//Selection(arr)
	//Heap(arr)
	//Quick(arr)
	//Merge(arr)
	//Bucket(arr)
	//Counting(arr)
	Radix(arr)
}

var l = []int{
	100,
	1000,
	10_000,
	100_000,
	1_000_000,
}

func BenchmarkUnsorted(b *testing.B) {
	for _, v := range l {
		b.Run(strconv.Itoa(v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				list := factory.InitListInt(v).Unsorted()
				b.StartTimer()
				Sort(list)
			}
		})
	}
}

func BenchmarkSorted(b *testing.B) {
	for _, v := range l {
		b.Run(strconv.Itoa(v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				list := factory.InitListInt(v).Sorted()
				b.StartTimer()
				Sort(list)
			}
		})
	}
}

func BenchmarkSortedWithUnsortedItemInMiddle(b *testing.B) {
	for _, v := range l {
		b.Run(strconv.Itoa(v), func(b *testing.B) {
			b.StopTimer()
			list := factory.InitListInt(v).SortedWithUnsortedItemInMiddle()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				Sort(list)
			}
		})
	}
}

func BenchmarkUnsortedRevers(b *testing.B) {
	for _, v := range l {
		b.Run(strconv.Itoa(v), func(b *testing.B) {
			b.StopTimer()
			list := factory.InitListInt(v).Revers()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				Sort(list)
			}
		})
	}
}