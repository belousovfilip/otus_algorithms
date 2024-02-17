package sort

import (
	"github.com/stretchr/testify/require"
	"otus_algo/internal/factory"
	"sort"
	"testing"
)

func TestBubble(t *testing.T) {
	expected := make([]int, 100)
	actual := factory.InitListInt(100).Unsorted()
	copy(expected, actual)
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
	Bubble(actual)
	require.Equal(t, expected, actual)
}
