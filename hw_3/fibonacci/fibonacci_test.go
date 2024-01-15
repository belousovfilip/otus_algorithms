package fibonacci

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"otus_algo/helpers"
	"testing"
)

type item struct {
	in  int64
	out int64
}

func TestRecursion(t *testing.T) {
	items, err := getItems(6)
	if err != nil {
		panic(err)
	}
	for _, v := range items {
		require.Equal(t, v.out, Recursion(v.in))
	}
}

func TestIterate(t *testing.T) {
	items, err := getItems(6)
	if err != nil {
		panic(err)
	}
	for _, v := range items {
		require.Equal(t, v.out, Iterate(v.in))
	}
}

func TestBine(t *testing.T) {
	items, err := getItems(6)
	if err != nil {
		panic(err)
	}
	for _, v := range items {
		require.Equal(t, v.out, Bine(v.in))
	}
}

func getItems(n int) ([]item, error) {
	var data []item
	for i := 0; i <= n; i++ {
		in, err := helpers.ReadInt64FromFile(fmt.Sprintf("./test/test.%d.in", i))
		if err != nil {
			return nil, err
		}
		out, err := helpers.ReadInt64FromFile(fmt.Sprintf("./test/test.%d.out", i))
		if err != nil {
			return nil, err
		}
		data = append(data, item{in, out})
	}
	return data, nil
}
