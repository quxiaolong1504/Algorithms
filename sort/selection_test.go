package sort

import (
	"testing"
)

func TestSelection(t *testing.T) {
	var arr = []int{40, 2, 13, 28, 16}
	Selection(arr)
	checkArrIsSorted(t, arr)
}
