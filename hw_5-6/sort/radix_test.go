package sort

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strconv"
	"testing"
)

func TestRadix(t *testing.T) {
	arr := []int{1, 523, 111, 224, 305, 3, 130}
	Radix(arr)
	fmt.Println(arr)
}

func TestComplexRadix(t *testing.T) {
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
			t.Run(p+strconv.Itoa(j), func(t *testing.T) {
				var n int
				path := p + strconv.Itoa(j)
				f, _ := os.Open(path + ".in")
				fmt.Fscan(f, &n)
				actual := getArrFromFile(f, n)
				f, _ = os.Open(path + ".out")
				expected := getArrFromFile(f, n)
				Radix(actual)
				require.Equal(t, expected, actual)
			})
		}
	}
}
