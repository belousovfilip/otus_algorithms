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

func TestInsert(t *testing.T) {
	expected := make([]int, 100)
	actual := factory.InitListInt(100).Unsorted()
	copy(expected, actual)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Insert(actual)
	require.Equal(t, expected, actual)
}

func TestComplexInsert(t *testing.T) {
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
		for j := 0; j <= 7; j++ {
			var n int
			f, _ := os.Open(p + strconv.Itoa(j) + ".in")
			fmt.Fscan(f, &n)
			actual := getArrFromFile(f, n)
			f, _ = os.Open(p + strconv.Itoa(j) + ".out")
			expected := getArrFromFile(f, n)
			Insert(actual)
			require.Equal(t, expected, actual)
		}
	}
}

func TestInsertBinary(t *testing.T) {
	expected := make([]int, 100)
	actual := factory.InitListInt(100).Unsorted()
	copy(expected, actual)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	InsertBinary(actual)
	require.Equal(t, expected, actual)
}

func TestComplexInsertBinary(t *testing.T) {
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
		for j := 0; j <= 7; j++ {
			var n int
			f, _ := os.Open(p + strconv.Itoa(j) + ".in")
			fmt.Fscan(f, &n)
			actual := getArrFromFile(f, n)
			f, _ = os.Open(p + strconv.Itoa(j) + ".out")
			expected := getArrFromFile(f, n)
			InsertBinary(actual)
			require.Equal(t, expected, actual)
		}
	}
}

func TestInsertShift(t *testing.T) {
	expected := make([]int, 100)
	actual := factory.InitListInt(100).Unsorted()
	copy(expected, actual)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	InsertShift(actual)
	require.Equal(t, expected, actual)
}

func TestComplexInsertShift(t *testing.T) {
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
		for j := 0; j <= 7; j++ {
			var n int
			f, _ := os.Open(p + strconv.Itoa(j) + ".in")
			fmt.Fscan(f, &n)
			actual := getArrFromFile(f, n)
			f, _ = os.Open(p + strconv.Itoa(j) + ".out")
			expected := getArrFromFile(f, n)
			InsertShift(actual)
			require.Equal(t, expected, actual)
		}
	}
}
