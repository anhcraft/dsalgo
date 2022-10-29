package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaximumSumSW(t *testing.T) {
	assert.Equal(t, 21, FindMaximumSumSW([]int{3, 0, 7, 4, 9, 8, 1, 5}, 3))
	assert.Equal(t, 73, FindMaximumSumSW([]int{-2, -14, -24, 21, 21, 16, -2, 17, -6, -17, 19, -13, -21, 19, 3, 16, -7, -17, 3, -11, 17, 15, -7, 25, -15, 5, 24, -11, -23, -10}, 5))
}

func TestFindMaximumPosSW(t *testing.T) {
	a, b := FindMaximumPosSW([]int{3, 0, 7, 4, 9, 8, 1, 5}, 3)
	assert.Equal(t, 3, a)
	assert.Equal(t, 5, b)

	a, b = FindMaximumPosSW([]int{-2, -14, -24, 21, 21, 16, -2, 17, -6, -17, 19, -13, -21, 19, 3, 16, -7, -17, 3, -11, 17, 15, -7, 25, -15, 5, 24, -11, -23, -10}, 5)
	assert.Equal(t, 3, a)
	assert.Equal(t, 7, b)
}

func TestFindMaximumSumKA(t *testing.T) {
	assert.Equal(t, 7, FindMaximumSumKA([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	assert.Equal(t, 18, FindMaximumSumKA([]int{-5, -5, 6, -9, -6, 1, 9, -7, -5, -9, 0, -7, 6, -5, 8, -9, -2, 0, -2, 3, 8, -8, -2, -2, 0, 10, -2, 8, -1, 3}))
}

func TestFindMaximumPosKA(t *testing.T) {
	a, b := FindMaximumPosKA([]int{-2, -3, 4, -1, -2, 1, 5, -3})
	assert.Equal(t, 2, a)
	assert.Equal(t, 6, b)

	a, b = FindMaximumPosKA([]int{-5, -5, 6, -9, -6, 1, 9, -7, -5, -9, 0, -7, 6, -5, 8, -9, -2, 0, -2, 3, 8, -8, -2, -2, 0, 10, -2, 8, -1, 3})
	assert.Equal(t, 25, a)
	assert.Equal(t, 29, b)
}

func TestFindSubarrayFromSum(t *testing.T) {
	a, b := FindSubarrayFromSum([]int{7, 8, 2, 1, 4, 5, 3}, 10)
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, b)

	a, b = FindSubarrayFromSum([]int{7, 2, 3, 1, 4, 5, 3}, 8)
	assert.Equal(t, 2, a)
	assert.Equal(t, 4, b)

	a, b = FindSubarrayFromSum([]int{6, 3, 0, 0, 1, 2, 4}, 4)
	assert.Equal(t, 1, a)
	assert.Equal(t, 4, b)

	a, b = FindSubarrayFromSum([]int{6, 3, 0, 1, 1, 2, 4}, 4)
	assert.Equal(t, 1, a)
	assert.Equal(t, 3, b)

	a, b = FindSubarrayFromSum([]int{6, 3, 0, 9, 1, 2, 4}, 4)
	assert.Equal(t, 6, a)
	assert.Equal(t, 6, b)

	a, b = FindSubarrayFromSum([]int{6, 3, 0, 9, 1, 2, 0}, 4)
	assert.Equal(t, -1, a)
	assert.Equal(t, -1, b)

	a, b = FindSubarrayFromSum([]int{4, 2, -3, 8, -7, -2, 0, 5, -6}, 5)
	assert.Equal(t, 2, a)
	assert.Equal(t, 3, b)

	a, b = FindSubarrayFromSum([]int{3, 9, -7, -6, -2, 8, 1, -1, 2, 4}, 0)
	assert.Equal(t, 3, a)
	assert.Equal(t, 5, b)
}

func TestFindSubarrayFromSumSW(t *testing.T) {
	a, b := FindSubarrayFromSumSW([]int{1, 0, 3, 4, 5, 4, 2}, 7)
	assert.Equal(t, 1, a)
	assert.Equal(t, 3, b)

	a, b = FindSubarrayFromSumSW([]int{2, 6, 8, 4, 3, 5, 9}, 7)
	assert.Equal(t, 3, a)
	assert.Equal(t, 4, b)

	a, b = FindSubarrayFromSumSW([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 1)
	assert.Equal(t, 8, a)
	assert.Equal(t, 8, b)
}

func TestFindSubarrayFromSumCM(t *testing.T) {
	a, b := FindSubarrayFromSumCM([]int{-5, 4, 3, -2, -1, 8}, 7)
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, b)
}
