package sort

import (
	"github.com/stretchr/testify/require"
	"otus_algo/internal/factory"
	"sort"
	"testing"
)

func TestInsert(t *testing.T) {
	expected := make([]int, 100)
	actual := factory.InitListInt(100).Unsorted()
	copy(expected, actual)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Insert(actual)
	require.Equal(t, expected, actual)
}

func TestBinarySearch(t *testing.T) {
	expected := make([]int, 100)
	actual := factory.InitListInt(100).Unsorted()
	copy(expected, actual)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	InsertShift(actual)
	require.Equal(t, expected, actual)
}

func TestInsertShift(t *testing.T) {
	expected := make([]int, 100)
	actual := factory.InitListInt(100).Unsorted()
	copy(expected, actual)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	InsertShift(actual)
	require.Equal(t, expected, actual)
}

func TestShift(t *testing.T) {
	var l []int
	l = []int{2, 3, 4, 1}
	shift(l, 0, 3)
	require.Equal(t, []int{1, 2, 3, 4}, l)
}

func TestIterateSearch(t *testing.T) {
	require.Equal(t, 0, iterateSearch([]int{11, 22}, 00))
	require.Equal(t, 0, iterateSearch([]int{11, 22}, 11))
	require.Equal(t, 1, iterateSearch([]int{11, 44, 55}, 33))
	require.Equal(t, 2, iterateSearch([]int{11, 44, 55}, 55))
	require.Equal(t, 3, iterateSearch([]int{11, 44, 55}, 66))
}
