package collections

type ArrayList[E any] struct {
	data []E
}

func NewArrayList[E any]() *ArrayList[E] {
	return &ArrayList[E]{
		data: make([]E, 0),
	}
}

func (a *ArrayList[E]) Append(elem E) {
	a.data = append(a.data, elem)
}

func (a *ArrayList[E]) Prepend(elem E) {
	a.Insert(0, elem)
}

func (a *ArrayList[E]) Insert(index int, elem E) {
	n := make([]E, len(a.data)+1)
	for i := 0; i < index; i++ {
		n[i] = a.data[i]
	}
	for i := index; i < len(a.data); i++ {
		n[i+1] = a.data[i]
	}
	n[index] = elem
	a.data = n
}

func (a *ArrayList[E]) Delete(index int) {
	n := make([]E, len(a.data)-1)
	for i := 0; i < index; i++ {
		n[i] = a.data[i]
	}
	for i := index + 1; i < len(a.data); i++ {
		n[i-1] = a.data[i]
	}
	a.data = n
}

func (a *ArrayList[E]) CopySlice(begin int, end int) *ArrayList[E] {
	if end < 0 {
		end += len(a.data)
	}
	n := make([]E, end-begin+1)
	for i := begin; i <= end; i++ {
		n[i-begin] = a.data[i]
	}
	return &ArrayList[E]{
		data: n,
	}
}

func (a *ArrayList[E]) Rotate(offset int) *ArrayList[E] {
	length := len(a.data)
	offset %= length
	if offset == 0 {
		return a
	}
	mid := -offset
	if offset > 0 {
		mid += length
	}
	head := a.data[0:mid]
	tail := a.data[mid:length]
	return &ArrayList[E]{
		data: append(tail, head...),
	}
}

func (a *ArrayList[E]) SearchAfterRotation(index int, offset int) int {
	i := (index + (offset % len(a.data))) % len(a.data)
	if i < 0 {
		i += len(a.data)
	}
	return i
}

func BinarySearch[E Ordered](a *ArrayList[E], elem E) int {
	lower := 0
	upper := len(a.data) - 1
	for lower <= upper && elem >= a.data[lower] && elem <= a.data[upper] {
		mid := (lower + upper) >> 1
		if elem < a.data[mid] {
			upper = mid - 1
		} else if elem > a.data[mid] {
			lower = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func BinarySearchFirstOccurrence[E Ordered](a *ArrayList[E], elem E) int {
	lower := 0
	upper := len(a.data) - 1
	for lower <= upper && elem >= a.data[lower] && elem <= a.data[upper] {
		mid := (lower + upper) >> 1
		if elem < a.data[mid] {
			upper = mid - 1
		} else if elem > a.data[mid] {
			lower = mid + 1
		} else {
			for i := mid; i > 0; i-- {
				if a.data[i-1] != elem {
					return i
				}
			}
		}
	}
	return -1
}

func BinarySearchLastOccurrence[E Ordered](a *ArrayList[E], elem E) int {
	lower := 0
	upper := len(a.data) - 1
	for lower <= upper && elem >= a.data[lower] && elem <= a.data[upper] {
		mid := (lower + upper) >> 1
		if elem < a.data[mid] {
			upper = mid - 1
		} else if elem > a.data[mid] {
			lower = mid + 1
		} else {
			for i := mid; i < len(a.data); i++ {
				if a.data[i+1] != elem {
					return i
				}
			}
		}
	}
	return -1
}

func InterpolationSearch[E Integer](a *ArrayList[E], elem E) int {
	lower := 0
	upper := len(a.data) - 1
	for lower <= upper && elem >= a.data[lower] && elem <= a.data[upper] {
		ratio := (elem - a.data[lower]) / (a.data[upper] - a.data[lower])
		mid := lower + (upper-lower)*int(ratio)
		if elem < a.data[mid] {
			upper = mid - 1
		} else if elem > a.data[mid] {
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
