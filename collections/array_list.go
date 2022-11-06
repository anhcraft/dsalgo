package collections

type ArrayList[E any] struct {
	Data []E
}

func NewArrayList[E any]() *ArrayList[E] {
	return &ArrayList[E]{
		Data: make([]E, 0),
	}
}

func (a *ArrayList[E]) Append(elem E) {
	a.Data = append(a.Data, elem)
}

func (a *ArrayList[E]) Prepend(elem E) {
	a.Insert(0, elem)
}

func (a *ArrayList[E]) Insert(index int, elem E) {
	n := make([]E, len(a.Data)+1)
	for i := 0; i < index; i++ {
		n[i] = a.Data[i]
	}
	for i := index; i < len(a.Data); i++ {
		n[i+1] = a.Data[i]
	}
	n[index] = elem
	a.Data = n
}

func (a *ArrayList[E]) Delete(index int) {
	n := make([]E, len(a.Data)-1)
	for i := 0; i < index; i++ {
		n[i] = a.Data[i]
	}
	for i := index + 1; i < len(a.Data); i++ {
		n[i-1] = a.Data[i]
	}
	a.Data = n
}

func (a *ArrayList[E]) CopySlice(begin int, end int) *ArrayList[E] {
	if end < 0 {
		end += len(a.Data)
	}
	n := make([]E, end-begin+1)
	for i := begin; i <= end; i++ {
		n[i-begin] = a.Data[i]
	}
	return &ArrayList[E]{
		Data: n,
	}
}

func (a *ArrayList[E]) Rotate(offset int) *ArrayList[E] {
	length := len(a.Data)
	offset %= length
	if offset == 0 {
		return a
	}
	mid := -offset
	if offset > 0 {
		mid += length
	}
	head := a.Data[0:mid]
	tail := a.Data[mid:length]
	return &ArrayList[E]{
		Data: append(tail, head...),
	}
}

func (a *ArrayList[E]) SearchAfterRotation(index int, offset int) int {
	i := (index + (offset % len(a.Data))) % len(a.Data)
	if i < 0 {
		i += len(a.Data)
	}
	return i
}

func (a *ArrayList[E]) Iterate(f func(e E)) {
	for _, v := range a.Data {
		f(v)
	}
}

func BinarySearch[E Ordered](a *ArrayList[E], elem E) int {
	lower := 0
	upper := len(a.Data) - 1
	for lower <= upper && elem >= a.Data[lower] && elem <= a.Data[upper] {
		mid := (lower + upper) >> 1
		if elem < a.Data[mid] {
			upper = mid - 1
		} else if elem > a.Data[mid] {
			lower = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func BinarySearchFirstOccurrence[E Ordered](a *ArrayList[E], elem E) int {
	lower := 0
	upper := len(a.Data) - 1
	for lower <= upper && elem >= a.Data[lower] && elem <= a.Data[upper] {
		mid := (lower + upper) >> 1
		if elem < a.Data[mid] {
			upper = mid - 1
		} else if elem > a.Data[mid] {
			lower = mid + 1
		} else {
			for i := mid; i > 0; i-- {
				if a.Data[i-1] != elem {
					return i
				}
			}
		}
	}
	return -1
}

func BinarySearchLastOccurrence[E Ordered](a *ArrayList[E], elem E) int {
	lower := 0
	upper := len(a.Data) - 1
	for lower <= upper && elem >= a.Data[lower] && elem <= a.Data[upper] {
		mid := (lower + upper) >> 1
		if elem < a.Data[mid] {
			upper = mid - 1
		} else if elem > a.Data[mid] {
			lower = mid + 1
		} else {
			for i := mid; i < len(a.Data); i++ {
				if a.Data[i+1] != elem {
					return i
				}
			}
		}
	}
	return -1
}

func InterpolationSearch[E Integer](a *ArrayList[E], elem E) int {
	lower := 0
	upper := len(a.Data) - 1
	for lower <= upper && elem >= a.Data[lower] && elem <= a.Data[upper] {
		ratio := (elem - a.Data[lower]) / (a.Data[upper] - a.Data[lower])
		mid := lower + (upper-lower)*int(ratio)
		if elem < a.Data[mid] {
			upper = mid - 1
		} else if elem > a.Data[mid] {
			lower = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func MergeOrdered[E Ordered](a []E, b []E) []E {
	i := 0
	j := 0
	k := 0
	c := make([]E, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c[k] = a[i]
			k++
			i++
		} else {
			c[k] = b[j]
			k++
			j++
		}
	}
	for i < len(a) {
		c[k] = a[i]
		k++
		i++
	}
	for j < len(b) {
		c[k] = b[j]
		k++
		j++
	}
	return c
}

func MergeSort[E Ordered](a []E) []E {
	if len(a) < 2 {
		return a
	}
	mid := len(a) >> 1
	left := MergeSort(a[:mid])
	right := MergeSort(a[mid:])
	return MergeOrdered(left, right)
}

func QuickSortPartition[E Ordered](a []E, low int, high int) int {
	pivot := a[(low+high)>>1]
	i := low - 1
	j := high + 1
	for true {
		for true {
			i++
			if a[i] >= pivot {
				break
			}
		}
		for true {
			j--
			if a[j] <= pivot {
				break
			}
		}
		if i < j {
			temp := a[i]
			a[i] = a[j]
			a[j] = temp
			continue
		} else {
			break
		}
	}
	return j
}

func quickSort[E Ordered](a []E, low int, high int) {
	if low >= high || len(a) < 2 {
		return
	}
	p := QuickSortPartition(a, low, high)
	quickSort(a, low, p)
	quickSort(a, p+1, high)
}

func QuickSort[E Ordered](a []E) []E {
	quickSort(a, 0, len(a)-1)
	return a
}

func InsertionSort[E Ordered](a []E) []E {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
	return a
}
