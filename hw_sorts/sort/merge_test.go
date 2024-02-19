package sort

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"otus_algo/internal/factory"
	"sort"
	"strconv"
	"testing"
)

func TestMerge(t *testing.T) {
	expected := make([]int, 110)
	actual := factory.InitListInt(110).Unsorted()
	copy(expected, actual)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Merge(actual)
	require.Equal(t, expected, actual)
}

func TestComplexMerge(t *testing.T) {
	getArrFromFile := func(r io.Reader, n int) []int {
		var v int
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &v)
			arr[i] = v
		}
		return arr
	}
	folders := []string{"0.random", "1.digits", "2.sorted", "3.revers"}
	for _, folder := range folders {
		p := "tests/" + folder + "/test."
		for j := 0; j <= 6; j++ {
			var n int
			f, _ := os.Open(p + strconv.Itoa(j) + ".in")
			fmt.Fscan(f, &n)
			actual := getArrFromFile(f, n)
			f, _ = os.Open(p + strconv.Itoa(j) + ".out")
			expected := getArrFromFile(f, n)
			Merge(actual)
			require.Equal(t, expected, actual)
		}
	}
}
