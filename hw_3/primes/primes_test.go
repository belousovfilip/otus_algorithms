package primes

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"otus_algo/internal"
	"testing"
)

type item struct {
	in  int
	out int
}

func TestIterate(t *testing.T) {
	items, err := getItems(7)
	if err != nil {
		panic(err)
	}
	for _, v := range items {
		require.Equal(t, v.out, Iterate(v.in))
	}
}

func TestEratosthenes(t *testing.T) {
	items, err := getItems(14)
	if err != nil {
		panic(err)
	}
	for _, v := range items {
		require.Equal(t, v.out, Eratosthenes(v.in), fmt.Sprintf("in:%d ", v.in))
	}
}
func getItems(n int) ([]item, error) {
	var data []item
	for i := 0; i <= n; i++ {
		in, err := internal.ReadIntFromFile(fmt.Sprintf("./test/test.%d.in", i))
		if err != nil {
			return nil, err
		}
		out, err := internal.ReadIntFromFile(fmt.Sprintf("./test/test.%d.out", i))
		if err != nil {
			return nil, err
		}
		data = append(data, item{in, out})
	}
	return data, nil
}
