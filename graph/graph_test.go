package graph_test

import (
	"testing"

	"github.com/chonlatee/trygenerics/graph"
	"github.com/stretchr/testify/require"
)

func Test_Graph(t *testing.T) {

	g := graph.New[string]()

	g.AddVertex("a")
	g.AddVertex("b")
	g.AddVertex("c")
	g.AddVertex("d")
	g.AddVertex("e")

	g.AddEdge("a", "b")
	g.AddEdge("a", "d")
	g.AddEdge("a", "e")

	g.AddEdge("b", "e")
	g.AddEdge("b", "c")

	g.AddEdge("c", "d")
	g.AddEdge("c", "e")

	g.AddEdge("d", "e")

	t.Run("depth-first search", func(t *testing.T) {
		res := g.TraversalDFS("a")
		require.Equal(t, []string{"a", "e", "c", "d", "b"}, res)
	})

	t.Run("breadth-first search", func(t *testing.T) {
		res := g.TraversalBFS("a")
		require.Equal(t, []string{"a", "b", "d", "e", "c"}, res)
	})

}
