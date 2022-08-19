package dequeue

import "container/list"

type deQueue[T any] struct {
	dl *list.List
}

func New[T any]() *deQueue[T] {
	return &deQueue[T]{
		dl: list.New(),
	}
}

func (dq *deQueue[T]) EnqueueFront(val T) {
	dq.dl.PushFront(val)
}

func (dq *deQueue[T]) EnqueueRear(val T) {
	dq.dl.PushBack(val)
}

func (dq *deQueue[T]) DequeueFront() T {
	e := dq.dl.Front()
	dq.dl.Remove(e)
	val := e.Value
	return val.(T)
}

func (dq *deQueue[T]) DequeueRear() T {
	e := dq.dl.Back()
	val := e.Value
	dq.dl.Remove(e)
	return val.(T)
}

func (dq *deQueue[T]) Front() T {
	e := dq.dl.Front()
	val := e.Value
	return val.(T)
}

func (dq *deQueue[T]) Rear() T {
	e := dq.dl.Back()
	val := e.Value
	return val.(T)
}

func (dq *deQueue[T]) IsEmpty() bool {
	return dq.dl.Len() == 0
}
