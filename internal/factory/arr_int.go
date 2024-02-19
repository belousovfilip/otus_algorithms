package factory

import (
	"math/rand"
)

type ArrInt struct {
	count int
}

func (l *ArrInt) Sorted() []int {
	list := make([]int, l.count)
	for i := 0; i < l.count; i++ {
		list[i] = i
	}
	return list
}

func (l *ArrInt) SortedWithUnsortedItemInMiddle() []int {

	list := l.Sorted()
	list[l.count/2] = l.count + 1
	return list
}

func (l *ArrInt) Unsorted() []int {

	list := make([]int, l.count)
	for i := 0; i < l.count; i++ {
		list[i] = rand.Intn(l.count)
	}
	return list
}

func (l *ArrInt) Revers() []int {
	list := make([]int, l.count)
	for i := 0; i < l.count; i++ {
		list[i] = l.count - i - 1
	}
	list[0] = l.count + 1
	return list
}

func InitListInt(count int) *ArrInt {
	return &ArrInt{count: count}
}
