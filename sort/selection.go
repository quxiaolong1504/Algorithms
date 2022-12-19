package sort

import "github.com/quxiaolong1504/Algorithms/constraints"

func Selection[T constraints.Ordered](arr []T) {
	for i := 0; i < len(arr)-1; i++ {
		var min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		// swap current i & min
		arr[i], arr[min] = arr[min], arr[i]
	}
}
