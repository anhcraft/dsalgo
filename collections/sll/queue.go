package sll

type QueueNode[E any] struct {
	Value E
	Next  *QueueNode[E]
}

func NewQueueNode[E any](val E, next *QueueNode[E]) *QueueNode[E] {
	return &QueueNode[E]{
		Value: val,
		Next:  next,
	}
}

type Queue[E any] struct {
	Size int
	Head *QueueNode[E]
	Tail *QueueNode[E]
}

func NewEmptyQueue[E any]() *Queue[E] {
	return &Queue[E]{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
}

func NewQueue[E any](head *QueueNode[E]) *Queue[E] {
	if head == nil {
		return NewEmptyQueue[E]()
	}
	n := 0
	node := head
	tail := head
	for node != nil {
		tail = node
		node = node.Next
		n++
	}
	return &Queue[E]{
		Size: n,
		Head: head,
		Tail: tail,
	}
}

func (s *Queue[E]) Offer(elem E) {
	node := NewQueueNode[E](elem, nil)
	if s.Head == nil {
		s.Head = node
		s.Tail = node
	} else {
		s.Tail.Next = node
		s.Tail = node
	}
	s.Size++
}

func (s *Queue[E]) Insert(elem E) {
	node := NewQueueNode[E](elem, nil)
	if s.Head == nil {
		s.Head = node
		s.Tail = node
	} else {
		t := s.Head
		s.Head = node
		node.Next = t
	}
	s.Size++
}

func (s *Queue[E]) Poll() (bool, *QueueNode[E]) {
	if s.Head == nil {
		return false, nil
	}
	s.Size--
	node := s.Head
	s.Head = node.Next
	if s.Head == nil {
		s.Tail = nil
	}
	return true, node
}

func (s *Queue[E]) ToArray() []E {
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
