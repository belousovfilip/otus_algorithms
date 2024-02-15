package factory

import (
	"math/rand"
)

type ListInt struct {
	count int
}

func (l *ListInt) Sorted() []int {
	list := make([]int, l.count)
	for i := 0; i < l.count; i++ {
		list[i] = i
	}
	return list
}

func (l *ListInt) SortedWithUnsortedItemInMiddle() []int {

	list := l.Sorted()
	list[l.count/2] = l.count + 1
	return list
}

func (l *ListInt) Unsorted() []int {

	list := make([]int, l.count)
	for i := 0; i < l.count; i++ {
		list[i] = rand.Intn(l.count)
	}
	return list
}

func (l *ListInt) Revers() []int {
	list := make([]int, l.count)
	for i := 0; i < l.count; i++ {
		list[i] = l.count - i - 1
	}
	list[0] = l.count + 1
	return list
}

func InitListInt(count int) *ListInt {
	return &ListInt{count: count}
}
