package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, Search[int](arr, 5), 4)
	assert.Equal(t, Search[int](arr, 1), 0)
	assert.Equal(t, Search[int](arr, 0), -1)
}

func TestSimilarSearch(t *testing.T) {
	var arr = []int{1, 10, 20, 32, 40, 48, 90}

	assert.Equal(t, SimilarSearch[int](arr, 30), 3)
	assert.Equal(t, SimilarSearch[int](arr, 11), 1)
	assert.Equal(t, SimilarSearch[int](arr, 48), 5)
}
