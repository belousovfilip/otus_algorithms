package queue

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPriority(t *testing.T) {
	t.Run("", func(t *testing.T) {
		queue := InitPriority()
		require.Nil(t, queue.dequeue())
	})
	t.Run("", func(t *testing.T) {
		queue := InitPriority()
		queue.enqueue(1, "-1-")
		require.Equal(t, "-1-", queue.dequeue().value)
		require.Nil(t, queue.dequeue())
	})
	t.Run("", func(t *testing.T) {
		queue := InitPriority()
		queue.enqueue(2, "-2-1")
		queue.enqueue(2, "-2-2")
		queue.enqueue(1, "-1-1-")
		queue.enqueue(3, "-3-1-")
		queue.enqueue(3, "-3-2-")
		require.Equal(t, "-1-1-", queue.dequeue().value)
		require.Equal(t, "-2-1", queue.dequeue().value)
		require.Equal(t, "-2-2", queue.dequeue().value)
		require.Equal(t, "-3-1-", queue.dequeue().value)
		require.Equal(t, "-3-2-", queue.dequeue().value)
		require.Nil(t, queue.dequeue())
	})
}
