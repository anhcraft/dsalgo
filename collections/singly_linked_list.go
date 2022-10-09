package collections

type SinglyLinkedListNode[E any] struct {
	value E
	next  *SinglyLinkedListNode[E]
}

func NewSinglyLinkedListNode[E any](val E, next *SinglyLinkedListNode[E]) *SinglyLinkedListNode[E] {
	return &SinglyLinkedListNode[E]{
		value: val,
		next:  next,
	}
}

type SinglyLinkedList[E any] struct {
	len  int
	head *SinglyLinkedListNode[E]
}

func NewEmptySinglyLinkedList[E any]() *SinglyLinkedList[E] {
	return &SinglyLinkedList[E]{
		len:  0,
		head: nil,
	}
}

func NewSinglyLinkedList[E any](head *SinglyLinkedListNode[E]) *SinglyLinkedList[E] {
	n := 0
	node := head
	for node != nil {
		n++
		node = node.next
	}
	return &SinglyLinkedList[E]{
		len:  n,
		head: head,
	}
}

func (s *SinglyLinkedList[E]) Get(i int) *SinglyLinkedListNode[E] {
	k := 0
	node := s.head
	for k <= i && node != nil {
		if k == i {
			return node
		}
		k++
		node = node.next
	}
	return nil
}

func (s *SinglyLinkedList[E]) Add(i int, elem E) {
	if i == 0 {
		s.head = NewSinglyLinkedListNode(elem, s.head)
		s.len++
		return
	}
	k := 0
	node := s.head
	for k < i && node != nil {
		k++
		if k == i {
			node.next = NewSinglyLinkedListNode(elem, node.next)
			s.len++
			break
		}
		node = node.next
	}
}

func (s *SinglyLinkedList[E]) Delete(i int) {
	if i == 0 {
		s.head = nil
		s.len = 0
		return
	}
	k := 0
	node := s.head
	for k < i && node != nil {
		k++
		if k == i {
			node.next = nil
			s.len = k
			break
		}
		node = node.next
	}
}

func (s *SinglyLinkedList[E]) ToArray() []E {
	arr := make([]E, s.len)
	next := s.head
	i := 0
	for next != nil {
		arr[i] = next.value
		next = next.next
		i++
	}
	return arr
}

// By adding to a linked list, the order is opposite to the given comparator
func mergeOrderedLinked[E Ordered](a *SinglyLinkedListNode[E], n int, b *SinglyLinkedListNode[E], m int) *SinglyLinkedList[E] {
	i := 0
	j := 0
	k := 0
	arr := make([]E, n+m)
	for a != nil && b != nil && i < n && j < m {
		if a.value <= b.value {
			arr[k] = a.value
			k++
			a = a.next
			i++
		} else {
			arr[k] = b.value
			k++
			b = b.next
			j++
		}
	}
	for a != nil && i < n {
		arr[k] = a.value
		k++
		a = a.next
		i++
	}
	for b != nil && j < m {
		arr[k] = b.value
		k++
		b = b.next
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
	for i := 0; i < n-mid && right.next != nil; i++ {
		right = right.next
	}
	return mergeOrderedLinked(mergeSortLinked(left, n-mid), n-mid, mergeSortLinked(right, mid), mid).head
}

func MergeSortLinked[E Ordered](a *SinglyLinkedList[E]) *SinglyLinkedList[E] {
	return NewSinglyLinkedList(mergeSortLinked(a.head, a.len))
}
