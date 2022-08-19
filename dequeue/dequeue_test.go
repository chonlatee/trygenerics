package dequeue_test

import (
	"testing"

	"github.com/chonlatee/trygenerics/dequeue"
	"github.com/stretchr/testify/require"
)

func Test_Dequeue(t *testing.T) {
	t.Run("insert front", func(t *testing.T) {
		q := dequeue.New[int]()
		q.EnqueueFront(30)
		q.EnqueueFront(20)

		require.Equal(t, 20, q.DequeueFront())
		require.Equal(t, 30, q.DequeueFront())
	})

	t.Run("insert rear", func(t *testing.T) {
		q := dequeue.New[int]()
		q.EnqueueRear(30)
		q.EnqueueRear(20)

		require.Equal(t, 30, q.DequeueFront())
		require.Equal(t, 20, q.DequeueFront())

	})

	t.Run("dequeue front", func(t *testing.T) {
		q := dequeue.New[int]()
		q.EnqueueRear(40)
		q.EnqueueRear(50)

		require.Equal(t, 40, q.DequeueFront())
		require.Equal(t, 50, q.DequeueFront())
	})

	t.Run("dequeue rear", func(t *testing.T) {
		q := dequeue.New[int]()
		q.EnqueueRear(50)
		q.EnqueueRear(60)

		require.Equal(t, 60, q.DequeueRear())
		require.Equal(t, 50, q.DequeueRear())
	})

	t.Run("front value", func(t *testing.T) {
		q := dequeue.New[int]()
		q.EnqueueRear(50)
		q.EnqueueRear(60)

		require.Equal(t, 50, q.Front())
	})

	t.Run("rear value", func(t *testing.T) {
		q := dequeue.New[int]()
		q.EnqueueRear(50)
		q.EnqueueRear(60)

		require.Equal(t, 60, q.Rear())
	})

	t.Run("is empty", func(t *testing.T) {
		q := dequeue.New[int]()
		q.EnqueueRear(50)
		q.EnqueueRear(60)

		q.DequeueFront()
		q.DequeueFront()
		
		require.Equal(t, true, q.IsEmpty())
	})
}
