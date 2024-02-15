package sort

import (
	"github.com/stretchr/testify/require"
	"otus_algo/internal/factory"
	"sort"
	"testing"
)

func TestShell(t *testing.T) {
	expected := make([]int, 10)
	actual := factory.InitListInt(10).Unsorted()
	copy(expected, actual)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Shell(actual)
	require.Equal(t, expected, actual)
}
