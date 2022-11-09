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

func (t *GenericTree[T]) BFS(f func(val *GenericTreeNode[T])) {
	queue := sll.NewQueue[*GenericTreeNode[T]](sll.NewQueueNode[*GenericTreeNode[T]](t.Root, nil))
	for true {
		ok, node := queue.Poll()
		if !ok || node == nil {
			break
		}
		f(node.Value)
		for _, v := range node.Value.Children {
			queue.Offer(v)
		}
	}
}

func (t *GenericTree[T]) DFSPreOrder(f func(val *GenericTreeNode[T])) {
	stack := sll.NewStack[*GenericTreeNode[T]](sll.NewStackNode[*GenericTreeNode[T]](t.Root, nil))
	for true {
		ok, node := stack.Poll()
		if !ok || node == nil {
			break
		}
		f(node.Value)
		for i := len(node.Value.Children) - 1; i >= 0; i-- {
			stack.Offer(node.Value.Children[i])
		}
	}
}

func (t *GenericTree[T]) DFSPreOrderReversed(f func(val *GenericTreeNode[T])) {
	stack := sll.NewStack[*GenericTreeNode[T]](sll.NewStackNode[*GenericTreeNode[T]](t.Root, nil))
	for true {
		ok, node := stack.Poll()
		if !ok || node == nil {
			break
		}
		f(node.Value)
		for _, v := range node.Value.Children {
			stack.Offer(v)
		}
	}
}

func (t *GenericTree[T]) DFSPostOrder(f func(val *GenericTreeNode[T])) {
	ordered := sll.NewStack[*GenericTreeNode[T]](nil)
	stack := sll.NewStack[*GenericTreeNode[T]](sll.NewStackNode[*GenericTreeNode[T]](t.Root, nil))
	for true {
		ok, node := stack.Poll()
		if !ok || node == nil {
			break
		}
		for _, v := range node.Value.Children {
			stack.Offer(v)
		}
		ordered.Offer(node.Value)
	}
	for true {
		ok, node := ordered.Poll()
		if !ok || node == nil {
			break
		}
		f(node.Value)
	}
}

func (t *GenericTree[T]) DFSPostOrderReversed(f func(val *GenericTreeNode[T])) {
	ordered := sll.NewStack[*GenericTreeNode[T]](nil)
	stack := sll.NewStack[*GenericTreeNode[T]](sll.NewStackNode[*GenericTreeNode[T]](t.Root, nil))
	for true {
		ok, node := stack.Poll()
		if !ok || node == nil {
			break
		}
		for i := len(node.Value.Children) - 1; i >= 0; i-- {
			stack.Offer(node.Value.Children[i])
		}
		ordered.Offer(node.Value)
	}
	for true {
		ok, node := ordered.Poll()
		if !ok || node == nil {
			break
		}
		f(node.Value)
	}
}
