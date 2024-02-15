package sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertS1(t *testing.T) {
	type in struct {
		arr   []int
		value int
	}
	type out struct{ index int }
	type item struct {
		in
		out
	}
	items := []item{
		{in{arr: []int{11, 22, 44, 55, 66, 77, 88, 99}, value: 0}, out{index: 0}},
		{in{arr: []int{11, 22, 44, 55, 66, 77, 88, 99}, value: 11}, out{index: 0}},
		{in{arr: []int{00, 11, 22, 33, 44, 55, 66, 77}, value: 88}, out{index: 8}},
		{in{arr: []int{11, 22, 33, 44, 66, 77, 88, 99}, value: 55}, out{index: 4}},
		{in{arr: []int{11, 22, 33, 44, 66, 77, 88, 99}, value: 44}, out{index: 4}},
	}
	for _, item := range items {
		require.Equal(t, item.out.index, binarySearch(item.in.arr, item.in.value))
	}

}
