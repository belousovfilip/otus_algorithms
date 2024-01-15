package power

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func equal(t *testing.T, out, actual float64) {
	require.Equal(t, int(out*10_000_000), int(actual*10_000_000))
}

func TestIterate(t *testing.T) {
	l := getItems()
	for _, v := range l {
		equal(t, v.out, Iterate(v.in.a, v.in.b))
	}
}
func TestMultiplication(t *testing.T) {
	l := getItems()
	for _, v := range l {
		equal(t, v.out, Multiplication(v.in.a, v.in.b))
	}
}
func TestBinary(t *testing.T) {
	l := getItems()
	for _, v := range l {
		equal(t, v.out, Binary(v.in.a, v.in.b))
	}
}

type in struct {
	a float64
	b float64
}

type item struct {
	in
	out float64
}

func getItems() []item {
	return []item{
		{in{2, 2}, 4.0},
		{in{3, 3}, 27.0},
		{in{2, 10}, 1024.0},
		{in{123456789, 0}, 1.0},
		{in{1.001, 1000}, 2.71692393224},
		{in{1.0001, 10000}, 2.71814592682},
		{in{1.00001, 100000}, 2.71826823719},
		{in{1.000001, 1000000}, 2.7182804691},
		{in{1.0000001, 10000000}, 2.71828169413},
		{in{1.00000001, 100000000}, 2.71828179835},
		{in{1.000000001, 1000000000}, 2.71828205201},
		{in{1.0000000001, 10000000000}, 2.71828205323},
	}
}
