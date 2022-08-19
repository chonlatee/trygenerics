package bst

import "github.com/chonlatee/trygenerics/dequeue"

type number interface {
	float64 | int64
}

type node[T number] struct {
	val   T
	left  *node[T]
	right *node[T]
}

func newNode[T number](val T) *node[T] {
	return &node[T]{
		val:   val,
		left:  nil,
		right: nil,
	}
}

type bst[T number] struct {
	root *node[T]
}

func New[T number]() *bst[T] {
	return &bst[T]{
		root: nil,
	}
}

func (b *bst[T]) Insert(val T) {
	n := newNode(val)

	if b.root == nil {
		b.root = n
		return
	}

	root := b.root
	var prev *node[T]
	for root != nil {
		prev = root
		if n.val <= root.val {
			root = root.left
		} else {
			root = root.right
		}
	}

	if n.val <= prev.val {
		prev.left = n
	} else {
		prev.right = n
	}
}

func (b *bst[T]) Search(val T) bool {
	root := b.root
	for root != nil {
		if root.val == val {
			return true
		} else if val < root.val {
			root = root.left
		} else {
			root = root.right
		}
	}

	return false
}

func (b *bst[T]) Traversal() []T {
	root := b.root
	q := dequeue.New[*node[T]]()
	q.EnqueueRear(root)

	var res []T

	for !q.IsEmpty() {
		v := q.DequeueFront()
		res = append(res, v.val)
		if v.left != nil {
			q.EnqueueRear(v.left)
		}
		if v.right != nil {
			q.EnqueueRear(v.right)
		}
	}

	return res
}
