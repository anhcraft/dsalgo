package problems

import (
	"dsalgo/utils"
)

// FindMaximumSumSW Find maximum sum of subarray length k using sliding window approach
func FindMaximumSumSW(list []int, k int) int {
	max := 0
	for _, v := range list[:k] {
		max += v
	}
	sum := max
	for i := k; i < len(list); i++ {
		sum += list[i]
		sum -= list[i-k]
		max = utils.MaxInt(max, sum)
	}
	return max
}

// FindMaximumSumKA Find maximum possible sum of a continuous subarray using Kadane algorithm
func FindMaximumSumKA(list []int) int {
	curr := 0
	max := 0
	for i := 0; i < len(list); i++ {
		curr = utils.MaxInt(0, curr+list[i])
		max = utils.MaxInt(max, curr)
	}
	return max
}

// FindMaximumPosSW Find a subarray length k combining the largest sum using sliding window approach
func FindMaximumPosSW(list []int, k int) (int, int) {
	tail := 0
	max := 0
	for _, v := range list[:k] {
		max += v
	}
	sum := max
	for i := k; i < len(list); i++ {
		sum += list[i]
		sum -= list[i-k]
		if sum > max {
			max = sum
			tail = i
		}
	}
	return tail - k + 1, tail
}

// FindMaximumPosKA Find the position of a subarray combining the largest sum using modified Kadane algorithm
func FindMaximumPosKA(list []int) (int, int) {
	curr := 0
	max := 0
	head := 0
	tail := 0
	for i := 0; i < len(list); i++ {
		next := curr + list[i]
		if next > 0 {
			curr = next
		} else {
			curr = 0
			head = i + 1
		}
		if curr > max {
			max = curr
			tail = i
		}
	}
	return head, tail
}

// FindSubarrayFromSum Find the subarray making up the given sum
func FindSubarrayFromSum(list []int, sum int) (int, int) {
	for i := 0; i < len(list); i++ {
		curr := 0
		for j := i; j < len(list); j++ {
			curr += list[j]
			if curr == sum {
				return i, j
			}
		}
	}
	return -1, -1
}

// FindSubarrayFromSumSW Find the subarray making up the given sum
// Condition: All indices are non-negative
func FindSubarrayFromSumSW(list []int, sum int) (int, int) {
	curr := 0
	head := 0
	tail := -1
	for tail < len(list) {
		for curr > sum && head < len(list) {
			curr -= list[head]
			head++
		}
		if curr == sum && head <= tail {
			return head, tail
		}
		tail++
		curr += list[tail]
	}
	return -1, -1
}

// FindSubarrayFromSumCM Find the subarray making up the given sum
// With caching map to speed up O(N^2) -> O(N)
func FindSubarrayFromSumCM(list []int, sum int) (int, int) {
	cache := make(map[int]int)
	curr := 0
	for i, n := range list {
		curr += n
		if curr == sum {
			return 0, i
		}
		v, b := cache[curr-sum]
		if b {
			return v + 1, i
		}
		cache[curr] = i
	}
	return -1, -1
}
