package sll

import (
	"dsalgo/utils"
)

type StackNode[E any] struct {
	Value E
	Next  *StackNode[E]
}

func NewStackNode[E any](val E, next *StackNode[E]) *StackNode[E] {
	return &StackNode[E]{
		Value: val,
		Next:  next,
	}
}

type Stack[E any] struct {
	Size int
	Head *StackNode[E]
}

func NewEmptyStack[E any]() *Stack[E] {
	return &Stack[E]{
		Size: 0,
		Head: nil,
	}
}

func NewStack[E any](head *StackNode[E]) *Stack[E] {
	n := 0
	node := head
	for node != nil {
		n++
		node = node.Next
	}
	return &Stack[E]{
		Size: n,
		Head: head,
	}
}

func (s *Stack[E]) Get(i int) *StackNode[E] {
	k := 0
	node := s.Head
	for k <= i && node != nil {
		if k == i {
			return node
		}
		k++
		node = node.Next
	}
	return nil
}

func (s *Stack[E]) Add(i int, elem E) {
	if i == 0 {
		s.Head = NewStackNode(elem, s.Head)
		s.Size++
		return
	}
	k := 0
	node := s.Head
	for k < i && node != nil {
		k++
		if k == i {
			node.Next = NewStackNode(elem, node.Next)
			s.Size++
			break
		}
		node = node.Next
	}
}

func (s *Stack[E]) Offer(elem E) {
	s.Add(0, elem)
}

func (s *Stack[E]) DeleteTail(i int) {
	if i == 0 {
		s.Head = nil
		s.Size = 0
		return
	}
	k := 0
	node := s.Head
	for k < i && node != nil {
		k++
		if k == i {
			node.Next = nil
			s.Size = k
			break
		}
		node = node.Next
	}
}

func (s *Stack[E]) DeleteAt(i int) (bool, *StackNode[E]) {
	if i == 0 {
		curr := s.Head
		s.Head = curr.Next
		s.Size--
		return true, curr
	}
	k := 0
	node := s.Head
	for k < i && node != nil {
		k++
		if k == i {
			curr := node.Next
			if curr != nil {
				node.Next = curr.Next
				s.Size--
			}
			return true, curr
		}
		node = node.Next
	}
	return false, nil
}

func (s *Stack[E]) ToArray() []E {
	arr := make([]E, s.Size)
	next := s.Head
	i := 0
	for next != nil {
		arr[i] = next.Value
		next = next.Next
		i++
	}
	return arr
}

func (s *Stack[E]) Poll() (bool, *StackNode[E]) {
	return s.DeleteAt(0)
}

// By adding to a linked list, the order is opposite to the given comparator
func mergeOrderedStack[E utils.Ordered](a *StackNode[E], n int, b *StackNode[E], m int) *Stack[E] {
	i := 0
	j := 0
	k := 0
	arr := make([]E, n+m)
	for a != nil && b != nil && i < n && j < m {
		if a.Value <= b.Value {
			arr[k] = a.Value
			k++
			a = a.Next
			i++
		} else {
			arr[k] = b.Value
			k++
			b = b.Next
			j++
		}
	}
	for a != nil && i < n {
		arr[k] = a.Value
		k++
		a = a.Next
		i++
	}
	for b != nil && j < m {
		arr[k] = b.Value
		k++
		b = b.Next
		j++
	}
	list := NewEmptyStack[E]()
	for q := len(arr) - 1; q >= 0; q-- {
		list.Add(0, arr[q])
	}
	return list
}

func mergeSortStack[E utils.Ordered](a *StackNode[E], n int) *StackNode[E] {
	if n < 2 {
		return a
	}
	mid := n >> 1
	left := a
	right := a
	for i := 0; i < n-mid && right.Next != nil; i++ {
		right = right.Next
	}
	return mergeOrderedStack(mergeSortStack(left, n-mid), n-mid, mergeSortStack(right, mid), mid).Head
}

func MergeSortStack[E utils.Ordered](a *Stack[E]) *Stack[E] {
	return NewStack(mergeSortStack(a.Head, a.Size))
}
