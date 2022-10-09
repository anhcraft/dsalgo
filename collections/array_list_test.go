package collections

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestArrayList1(t *testing.T) {
	arr := NewArrayList[int]()
	arr.Append(3)
	arr.Append(-5)
	arr.Append(-1)
	arr.Append(7)
	arr.Append(4)
	assert.Equal(t, []int{3, -5, -1, 7, 4}, arr.data)

	arr.Insert(0, 8)
	assert.Equal(t, []int{8, 3, -5, -1, 7, 4}, arr.data)
	arr.Insert(1, -6)
	assert.Equal(t, []int{8, -6, 3, -5, -1, 7, 4}, arr.data)

	assert.Equal(t, []int{-6, 3, -5}, arr.CopySlice(1, 3).data)
	assert.Equal(t, []int{-1, 7, 4}, arr.CopySlice(4, -1).data)

	arr.Delete(0)
	assert.Equal(t, []int{-6, 3, -5, -1, 7, 4}, arr.data)

	arr.Delete(3)
	assert.Equal(t, []int{-6, 3, -5, 7, 4}, arr.data)
}

func TestArrayList2(t *testing.T) {
	arr := NewArrayList[int]()
	arr.Append(-7)
	arr.Append(-2)
	arr.Append(4)
	arr.Append(6)
	arr.Append(11)
	assert.Equal(t, 4, BinarySearch(arr, 11))
	assert.Equal(t, 3, BinarySearch(arr, 6))
	assert.Equal(t, 2, BinarySearch(arr, 4))
	assert.Equal(t, -1, BinarySearch(arr, -3))

	arr.Append(11)
	arr.Append(13)
	arr.Append(17)
	arr.Append(17)
	arr.Append(19)
	assert.Equal(t, []int{-7, -2, 4, 6, 11, 11, 13, 17, 17, 19}, arr.data)
	assert.Equal(t, 4, BinarySearchFirstOccurrence(arr, 11))
	assert.Equal(t, 7, BinarySearchFirstOccurrence(arr, 17))
	assert.Equal(t, 5, BinarySearchLastOccurrence(arr, 11))
	assert.Equal(t, 8, BinarySearchLastOccurrence(arr, 17))
}

func TestArrayList3(t *testing.T) {
	arr := &ArrayList[int]{
		data: []int{0, 10, 20, 30, 40, 50, 60},
	}
	assert.Equal(t, 1, InterpolationSearch(arr, 10))

	arr.Append(100)
	arr.Append(200)
	arr.Append(300)
	arr.Append(500)
	assert.Equal(t, 8, InterpolationSearch(arr, 200))
}

func TestArrayList4(t *testing.T) {
	arr := &ArrayList[int]{
		data: []int{1, 3, 7, 4, 2},
	}
	assert.Equal(t, []int{4, 2, 1, 3, 7}, arr.Rotate(2).data)
	assert.Equal(t, []int{7, 4, 2, 1, 3}, arr.Rotate(3).data)
	assert.Equal(t, []int{3, 7, 4, 2, 1}, arr.Rotate(4).data)
	assert.Equal(t, []int{1, 3, 7, 4, 2}, arr.Rotate(5).data)
	assert.Equal(t, []int{7, 4, 2, 1, 3}, arr.Rotate(-2).data)
	assert.Equal(t, []int{4, 2, 1, 3, 7}, arr.Rotate(-3).data)
	assert.Equal(t, []int{2, 1, 3, 7, 4}, arr.Rotate(-4).data)
	assert.Equal(t, []int{1, 3, 7, 4, 2}, arr.Rotate(-5).data)

	assert.Equal(t, 2, arr.SearchAfterRotation(0, 2))
	assert.Equal(t, 2, arr.SearchAfterRotation(4, 3))
	assert.Equal(t, 3, arr.SearchAfterRotation(0, -2))
	assert.Equal(t, 1, arr.SearchAfterRotation(4, -3))
}

func TestArrayList5(t *testing.T) {
	arr1 := []int{2, 7, 9, 11, 14}
	arr2 := []int{0, 3, 5, 8, 13, 19}
	assert.Equal(t, []int{0, 2, 3, 5, 7, 8, 9, 11, 13, 14, 19}, MergeOrdered(arr1, arr2))
}

func TestArrayList6(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 7, 8}, MergeSort([]int{3, 0, 1, 2, 7, 8}))
	assert.Equal(t, []int{2, 2, 4, 5, 6, 7, 8, 9, 11, 13}, MergeSort([]int{9, 6, 2, 5, 11, 13, 4, 2, 7, 8}))
}

func TestArrayList7(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 7, 8}, QuickSort([]int{3, 0, 1, 2, 7, 8}))
	assert.Equal(t, []int{2, 2, 4, 5, 6, 7, 8, 9, 11, 13}, QuickSort([]int{9, 6, 2, 5, 11, 13, 4, 2, 7, 8}))
}

func TestArrayList8(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 7, 8}, InsertionSort([]int{3, 0, 1, 2, 7, 8}))
	assert.Equal(t, []int{2, 2, 4, 5, 6, 7, 8, 9, 11, 13}, InsertionSort([]int{9, 6, 2, 5, 11, 13, 4, 2, 7, 8}))
}

func TestArrayList9(t *testing.T) {
	list := NewArrayList[int]()
	n := 1_000_000
	for i := 0; i < n; i++ {
		list.Append(rand.Intn(n << 1))
	}
	last := 0
	for _, v := range MergeSort(list.data) {
		if v < last {
			t.FailNow()
		}
		last = v
	}
}

func TestArrayList10(t *testing.T) {
	list := NewArrayList[int]()
	n := 1_000_000
	for i := 0; i < n; i++ {
		list.Append(rand.Intn(n << 1))
	}
	last := 0
	for _, v := range QuickSort(list.data) {
		if v < last {
			t.FailNow()
		}
		last = v
	}
}

func TestArrayList11(t *testing.T) {
	list := NewArrayList[int]()
	n := 100_000
	for i := 0; i < n; i++ {
		list.Append(rand.Intn(n << 1))
	}
	last := 0
	for _, v := range InsertionSort(list.data) {
		if v < last {
			t.FailNow()
		}
		last = v
	}
}
