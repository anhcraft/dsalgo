package tree

import (
	"dsalgo/collections/sll"
)

type GenericTreeNode[T any] struct {
	Value    T
	Children []*GenericTreeNode[T]
}

func NewGenericTreeNode[T any](val T) *GenericTreeNode[T] {
	return &GenericTreeNode[T]{
		Value: val,
	}
}

func (n *GenericTreeNode[T]) Insert(nodes ...*GenericTreeNode[T]) *GenericTreeNode[T] {
	n.Children = append(n.Children, nodes...)
	return n
}

func (n *GenericTreeNode[T]) InsertAndGet(node *GenericTreeNode[T]) *GenericTreeNode[T] {
	n.Children = append(n.Children, node)
	return node
}

type GenericTree[T any] struct {
	Root *GenericTreeNode[T]
}

func NewGenericTree[T any](root *GenericTreeNode[T]) *GenericTree[T] {
	return &GenericTree[T]{
		Root: root,
	}
}

// BFS Breadth-first search
func (t *GenericTree[T]) BFS(f func(val *GenericTreeNode[T])) {
	queue := sll.NewStack[*GenericTreeNode[T]](sll.NewStackNode[*GenericTreeNode[T]](t.Root, nil))
	for true {
		ok, head := queue.Poll()
		if !ok || head == nil {
			break
		}
		f(head.Value)
		for _, v := range head.Value.Children {
			queue.Offer(v)
		}
	}
}
