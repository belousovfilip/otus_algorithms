package array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSingle(t *testing.T) {
	RunTestAdd(t, InitSingle())
	RunTestRemove(t, InitSingle())
	RunTestAddByIndex(t, InitSingle())
}
func TestFactor(t *testing.T) {
	RunTestAdd(t, InitFactor())
	RunTestRemove(t, InitFactor())
	RunTestAddByIndex(t, InitFactor())
}

func TestVector(t *testing.T) {
	RunTestAdd(t, InitVector())
	RunTestRemove(t, InitVector())
	RunTestAddByIndex(t, InitVector())
}

func TestSlice(t *testing.T) {
	RunTestAdd(t, InitSlice())
	RunTestRemove(t, InitSlice())
	RunTestAddByIndex(t, InitSlice())
}

func RunTestAdd(t *testing.T, arr Array) {
	for i := 0; i < 1000; i++ {
		arr.Add(i)
		require.Equal(t, i+1, arr.Size())
		require.Equal(t, i, arr.Get(i))
	}
}

func RunTestRemove(t *testing.T, arr Array) {
	for i := 0; i < 1000; i++ {
		arr.Add(i)
	}
	require.Equal(t, 1000, arr.Size())
	for i := 500; i < 550; i++ {
		arr.Remove(i)
		require.Equal(t, 1000+500-1-i, arr.Size())
	}
	arr.Add(888)
	arr.Add(999)
	require.Equal(t, 999, arr.Remove(arr.Size()-1))
	require.Equal(t, 888, arr.Get(arr.Size()-1))
}

func RunTestAddByIndex(t *testing.T, arr Array) {
	arr.Add(33)
	arr.AddByIndex(0, 22)
	require.Equal(t, 22, arr.Get(0))
	arr.Add(33)
	for i := 1; i < 1000; i++ {
		arr.AddByIndex(i, i)
		require.Equal(t, i, arr.Size()-3)
		require.Equal(t, i, arr.Get(i))
	}
}
