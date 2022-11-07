package collections

import (
	"dsalgo/utils"
	"math"
)

type Grid[T utils.Ordered] struct {
	Data [][]T
}

func NewGrid[T utils.Ordered](data [][]T) Grid[T] {
	return Grid[T]{
		Data: data,
	}
}

func (g *Grid[T]) Width() int {
	return len(g.Data[0])
}

func (g *Grid[T]) Height() int {
	return len(g.Data)
}

func (g *Grid[T]) PosToOrder(x int, y int) int {
	return y*g.Width() + x
}

func (g *Grid[T]) OrderToPos(i int) (int, int) {
	return i % g.Width(), i / g.Width()
}

func (g *Grid[T]) GetByPos(x int, y int) T {
	return g.Data[y][x]
}

func (g *Grid[T]) GetByOrder(i int) T {
	m, n := g.OrderToPos(i)
	return g.Data[n][m]
}

func (g *Grid[T]) BinarySearch(elem T) int {
	lower := 0
	upper := g.Width()*g.Height() - 1
	for lower <= upper {
		mid := (lower + upper) >> 1
		m := g.GetByOrder(mid)
		if elem == m {
			return mid
		} else if elem < m {
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}
	return -1
}

func (g *Grid[T]) BubbleSort() {
	n := g.Width() * g.Height()
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			x1, y1 := g.OrderToPos(j)
			x2, y2 := g.OrderToPos(j + 1)
			if g.Data[y1][x1] > g.Data[y2][x2] {
				t := g.Data[y1][x1]
				g.Data[y1][x1] = g.Data[y2][x2]
				g.Data[y2][x2] = t
			}
		}
	}
}

// FindMaximumSumArea Kadane algorithm for finding maximum sum of area
// Returns: top, left, bottom, right, max
func FindMaximumSumArea(g Grid[int]) (int, int, int, int, int) {
	maxGlobal := math.MinInt
	top := -1
	left := -1
	bottom := -1
	right := -1
	for i := 0; i < g.Width(); i++ {
		temp := make([]int, g.Height())
		for j := i; j < g.Width(); j++ {
			up := -1
			down := -1
			maxLocal := math.MinInt
			m := 0
			n := 0
			curr := 0
			for r := 0; r < g.Height(); r++ {
				temp[r] += g.Data[r][j]
				curr += temp[r]
				if curr < 0 {
					curr = 0
					m = r + 1
					n = m
				} else if curr > maxLocal {
					n = r
					maxLocal = curr
					up = m
					down = n
				}
			}
			if maxLocal > maxGlobal {
				maxGlobal = maxLocal
				top = up
				left = i
				bottom = down
				right = j
			}
		}
	}
	return top, left, bottom, right, maxGlobal
}
