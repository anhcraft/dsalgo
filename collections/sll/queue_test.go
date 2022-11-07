package sll

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue1(t *testing.T) {
	queue := NewEmptyQueue[int]()
	queue.Offer(3)
	queue.Offer(8)
	queue.Offer(2)
	queue.Offer(9)
	queue.Offer(4)
	assert.Equal(t, []int{3, 8, 2, 9, 4}, queue.ToArray())

	ok, node := queue.Poll()
	assert.True(t, ok)
	assert.Equal(t, 3, node.Value)
	assert.Equal(t, []int{8, 2, 9, 4}, queue.ToArray())

	queue.Poll()
	queue.Poll()
	queue.Poll()
	ok, node = queue.Poll()
	assert.True(t, ok)
	assert.Equal(t, 4, node.Value)
	assert.Equal(t, []int{}, queue.ToArray())
	assert.Equal(t, 0, queue.Size)
}
