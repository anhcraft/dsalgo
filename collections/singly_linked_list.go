package collections

type SinglyLinkedListNode[E any] struct {
	Value E
	Next  *SinglyLinkedListNode[E]
}

func NewSinglyLinkedListNode[E any](val E, next *SinglyLinkedListNode[E]) *SinglyLinkedListNode[E] {
	return &SinglyLinkedListNode[E]{
		Value: val,
		Next:  next,
	}
}

type SinglyLinkedList[E any] struct {
	Size int
	Head *SinglyLinkedListNode[E]
}

func NewEmptySinglyLinkedList[E any]() *SinglyLinkedList[E] {
	return &SinglyLinkedList[E]{
		Size: 0,
		Head: nil,
	}
}

func NewSinglyLinkedList[E any](head *SinglyLinkedListNode[E]) *SinglyLinkedList[E] {
	n := 0
	node := head
	for node != nil {
		n++
		node = node.Next
	}
	return &SinglyLinkedList[E]{
		Size: n,
		Head: head,
	}
}

func (s *SinglyLinkedList[E]) Get(i int) *SinglyLinkedListNode[E] {
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

func (s *SinglyLinkedList[E]) Add(i int, elem E) {
	if i == 0 {
		s.Head = NewSinglyLinkedListNode(elem, s.Head)
		s.Size++
		return
	}
	k := 0
	node := s.Head
	for k < i && node != nil {
		k++
		if k == i {
			node.Next = NewSinglyLinkedListNode(elem, node.Next)
			s.Size++
			break
		}
		node = node.Next
	}
}

func (s *SinglyLinkedList[E]) Offer(elem E) {
	s.Add(0, elem)
}

func (s *SinglyLinkedList[E]) DeleteTail(i int) {
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

func (s *SinglyLinkedList[E]) DeleteAt(i int) (bool, *SinglyLinkedListNode[E]) {
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

func (s *SinglyLinkedList[E]) ToArray() []E {
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

func (s *SinglyLinkedList[E]) Poll() (bool, *SinglyLinkedListNode[E]) {
	return s.DeleteAt(s.Size - 1)
}

// By adding to a linked list, the order is opposite to the given comparator
func mergeOrderedLinked[E Ordered](a *SinglyLinkedListNode[E], n int, b *SinglyLinkedListNode[E], m int) *SinglyLinkedList[E] {
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
	list := NewEmptySinglyLinkedList[E]()
	for q := len(arr) - 1; q >= 0; q-- {
		list.Add(0, arr[q])
	}
	return list
}

func mergeSortLinked[E Ordered](a *SinglyLinkedListNode[E], n int) *SinglyLinkedListNode[E] {
	if n < 2 {
		return a
	}
	mid := n >> 1
	left := a
	right := a
	for i := 0; i < n-mid && right.Next != nil; i++ {
		right = right.Next
	}
	return mergeOrderedLinked(mergeSortLinked(left, n-mid), n-mid, mergeSortLinked(right, mid), mid).Head
}

func MergeSortLinked[E Ordered](a *SinglyLinkedList[E]) *SinglyLinkedList[E] {
	return NewSinglyLinkedList(mergeSortLinked(a.Head, a.Size))
}
