package sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBubble(t *testing.T) {
	arr := Bubble(genRandomArr(10))
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			assert.True(t, arr[i] < arr[j])
		}
	}
}

func genRandomArr(length int) []int {
	const min, max = -9999999, 9999999
	var arr = make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(max-min) + min
	}
	return arr
}
