package search

import (
	"github.com/quxiaolong1504/Algorithms/constraints"
)

// Search 精确查找
func Search[T constraints.Ordered](arr []T, k T) int {
	var l, r = 0, len(arr) - 1
	for l <= r {
		var mid = l/2 + r/2 // 避免 int 溢出
		if arr[mid] == k {
			return mid
		} else if k < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// SimilarSearch 近似查找, Closet to T
func SimilarSearch[T constraints.Number](arr []T, k T) int {
	var l, r = 0, len(arr) - 1

	// 最后留两个元素
	for l < r-1 {
		var mid = l/2 + r/2
		if k < arr[mid] {
			r = mid
		} else {
			l = mid
		}
	}

	if k-arr[l] < arr[r]-k {
		// 左边更接近 k
		return l
	} else {
		// 右边更接近 k
		return r
	}

}
