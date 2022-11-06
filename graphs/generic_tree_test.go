package graphs

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
	assert.Equal(t, list.Data, []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"})
}
