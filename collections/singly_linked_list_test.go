package collections

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSinglyLinkedList1(t *testing.T) {
	list := NewEmptySinglyLinkedList[int]()
	list.Add(0, 3)
	list.Add(0, 5)
	list.Add(0, 7)
	assert.Equal(t, []int{7, 5, 3}, list.ToArray())

	list.Add(1, 8)
	assert.Equal(t, []int{7, 8, 5, 3}, list.ToArray())

	list.Add(3, 9)
	assert.Equal(t, []int{7, 8, 5, 9, 3}, list.ToArray())

	list.Add(5, 11)
	assert.Equal(t, []int{7, 8, 5, 9, 3, 11}, list.ToArray())

	assert.Equal(t, 7, list.Get(0).Value)
	assert.Equal(t, 5, list.Get(2).Value)

	list.DeleteTail(3)
	assert.Equal(t, []int{7, 8, 5}, list.ToArray())

	list.DeleteTail(0)
	assert.Equal(t, []int{}, list.ToArray())

	list.Add(0, 9)
	list.Add(0, 4)
	list.Add(0, 1)
	assert.Equal(t, []int{1, 4, 9}, list.ToArray())

	list.DeleteAt(1)
	assert.Equal(t, []int{1, 9}, list.ToArray())

	list.DeleteAt(0)
	assert.Equal(t, []int{9}, list.ToArray())
}

func TestSinglyLinkedList2(t *testing.T) {
	list := NewEmptySinglyLinkedList[int]()
	list.Add(0, 3)
	list.Add(0, 5)
	list.Add(0, 7)
	list.Add(0, 1)
	list.Add(0, 4)
	list.Add(0, 2)
	list.Add(0, 9)
	assert.Equal(t, []int{9, 2, 4, 1, 7, 5, 3}, list.ToArray())
	assert.Equal(t, []int{1, 2, 3, 4, 5, 7, 9}, MergeSortLinked(list).ToArray())

	list.Add(0, -2)
	list.Add(0, -4)
	list.Add(0, -3)
	list.Add(0, -5)
	assert.Equal(t, []int{-5, -3, -4, -2, 9, 2, 4, 1, 7, 5, 3}, list.ToArray())
	assert.Equal(t, []int{-5, -4, -3, -2, 1, 2, 3, 4, 5, 7, 9}, MergeSortLinked(list).ToArray())
}

func TestSinglyLinkedList3(t *testing.T) {
	list := NewEmptySinglyLinkedList[int]()
	n := 1_000_000
	for i := 0; i < n; i++ {
		list.Add(0, rand.Intn(n<<1))
	}
	last := 0
	node := MergeSortLinked(list).Head
	for node != nil {
		if node.Value < last {
			t.FailNow()
		}
		node = node.Next
	}
}
