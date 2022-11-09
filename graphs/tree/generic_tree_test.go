package tree

import (
	"dsalgo/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenericTree1(t *testing.T) {
	tree := NewGenericTree(NewGenericTreeNode("A"))
	tree.Root.InsertAndGet(NewGenericTreeNode("B")).Insert(NewGenericTreeNode("E"), NewGenericTreeNode("F"))
	tree.Root.InsertAndGet(NewGenericTreeNode("C")).Insert(NewGenericTreeNode("G"), NewGenericTreeNode("H"), NewGenericTreeNode("I"))
	tree.Root.InsertAndGet(NewGenericTreeNode("D")).InsertAndGet(
		NewGenericTreeNode("J").Insert(
			NewGenericTreeNode("K"),
			NewGenericTreeNode("L"),
		),
	)
	list := collections.NewArrayList[string]()
	tree.BFS(func(e *GenericTreeNode[string]) {
		list.Append(e.Value)
	})
	assert.Equal(t, []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}, list.Data)
}

func TestGenericTree2(t *testing.T) {
	tree := NewGenericTree(NewGenericTreeNode(1))
	tree.Root.InsertAndGet(NewGenericTreeNode(2)).Insert(NewGenericTreeNode(4), NewGenericTreeNode(5), NewGenericTreeNode(6))
	tree.Root.InsertAndGet(NewGenericTreeNode(3)).InsertAndGet(NewGenericTreeNode(7)).Insert(NewGenericTreeNode(8), NewGenericTreeNode(9))
	list := collections.NewArrayList[int]()
	tree.DFSPreOrder(func(e *GenericTreeNode[int]) {
		list.Append(e.Value)
	})
	assert.Equal(t, []int{1, 2, 4, 5, 6, 3, 7, 8, 9}, list.Data)
}

func TestGenericTree3(t *testing.T) {
	tree := NewGenericTree(NewGenericTreeNode(1))
	tree.Root.InsertAndGet(NewGenericTreeNode(2)).Insert(NewGenericTreeNode(4), NewGenericTreeNode(5), NewGenericTreeNode(6))
	tree.Root.InsertAndGet(NewGenericTreeNode(3)).InsertAndGet(NewGenericTreeNode(7)).Insert(NewGenericTreeNode(8), NewGenericTreeNode(9))
	list := collections.NewArrayList[int]()
	tree.DFSPreOrderReversed(func(e *GenericTreeNode[int]) {
		list.Append(e.Value)
	})
	assert.Equal(t, []int{1, 3, 7, 9, 8, 2, 6, 5, 4}, list.Data)
}

func TestGenericTree4(t *testing.T) {
	tree := NewGenericTree(NewGenericTreeNode(1))
	tree.Root.InsertAndGet(NewGenericTreeNode(2)).Insert(NewGenericTreeNode(4), NewGenericTreeNode(5), NewGenericTreeNode(6))
	tree.Root.InsertAndGet(NewGenericTreeNode(3)).InsertAndGet(NewGenericTreeNode(7)).Insert(NewGenericTreeNode(8), NewGenericTreeNode(9))
	list := collections.NewArrayList[int]()
	tree.DFSPostOrder(func(e *GenericTreeNode[int]) {
		list.Append(e.Value)
	})
	assert.Equal(t, []int{4, 5, 6, 2, 8, 9, 7, 3, 1}, list.Data)
}

func TestGenericTree5(t *testing.T) {
	tree := NewGenericTree(NewGenericTreeNode(1))
	tree.Root.InsertAndGet(NewGenericTreeNode(2)).Insert(NewGenericTreeNode(4), NewGenericTreeNode(5), NewGenericTreeNode(6))
	tree.Root.InsertAndGet(NewGenericTreeNode(3)).InsertAndGet(NewGenericTreeNode(7)).Insert(NewGenericTreeNode(8), NewGenericTreeNode(9))
	list := collections.NewArrayList[int]()
	tree.DFSPostOrderReversed(func(e *GenericTreeNode[int]) {
		list.Append(e.Value)
	})
	assert.Equal(t, []int{9, 8, 7, 3, 6, 5, 4, 2, 1}, list.Data)
}
