package graph

import (
	"github.com/chonlatee/trygenerics/dequeue"
	"golang.org/x/exp/constraints"
)

type vertex[T constraints.Ordered] struct {
	key  T
	edge []*vertex[T]
}

type graph[T constraints.Ordered] struct {
	vertices map[T]*vertex[T]
}

func newVertex[T constraints.Ordered](key T) *vertex[T] {
	return &vertex[T]{
		key:  key,
		edge: []*vertex[T]{},
	}
}

func New[T constraints.Ordered]() *graph[T] {
	return &graph[T]{
		vertices: make(map[T]*vertex[T]),
	}
}

func (g *graph[T]) AddVertex(key T) {
	v := newVertex(key)
	g.vertices[key] = v
}

func (g *graph[T]) AddEdge(key1, key2 T) {
	// TODO: should check vertices contain key1 and key2
	k1 := g.vertices[key1]
	k2 := g.vertices[key2]
	var list []T
	for _, v := range k1.edge {
		list = append(list, v.key)
	}
	if !isContain[T](list, k2.key) {
		k1.edge = append(k1.edge, k2)
		k2.edge = append(k2.edge, k1)
	}
}

func (g *graph[T]) TraversalDFS(start T) []T {
	visited := make(map[T]bool)
	stack := dequeue.New[T]()
	stack.EnqueueRear(start)
	var res []T

	for !stack.IsEmpty() {
		n := stack.DequeueRear()
		visited[n] = true
		res = append(res, n)
		for _, v := range g.vertices[n].edge {
			if _, ok := visited[v.key]; !ok {
				visited[v.key] = true
				stack.EnqueueRear(v.key)
			}
		}
	}

	return res
}

func (g *graph[T]) TraversalBFS(start T) []T {
	visited := make(map[T]bool)
	queue := dequeue.New[T]()
	queue.EnqueueRear(start)
	var res []T
	for !queue.IsEmpty() {
		n := queue.DequeueFront()
		res = append(res, n)
		visited[n] = true
		for _, v := range g.vertices[n].edge {
			if _, ok := visited[v.key]; !ok {
				visited[v.key] = true
				queue.EnqueueRear(v.key)
			}
		}
	}

	return res
}

func isContain[T constraints.Ordered](l []T, val T) bool {
	for _, v := range l {
		if v == val {
			return true
		}
	}
	return false
}
