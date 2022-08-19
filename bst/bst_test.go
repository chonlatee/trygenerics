package bst_test

import (
	"testing"

	"github.com/chonlatee/trygenerics/bst"
	"github.com/stretchr/testify/require"
)

func Test_BST(t *testing.T) {
	b := bst.New[int64]()
	b.Insert(10)
	b.Insert(8)
	b.Insert(7)
	b.Insert(9)
	b.Insert(20)
	b.Insert(30)
	b.Insert(40)
	b.Insert(50)

	t.Run("traversal", func(t *testing.T) {
		res := b.Traversal()
		require.Equal(t, []int64{10, 8, 20, 7, 9, 30, 40, 50}, res)
	})

	t.Run("search found", func(t *testing.T) {
		require.Equal(t, true, b.Search(20))
	})

	t.Run("search not found", func(t *testing.T) {
		require.Equal(t, false, b.Search(60))
	})
}
