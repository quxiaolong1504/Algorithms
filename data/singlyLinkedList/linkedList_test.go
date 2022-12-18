package singlyLinkedList

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestNewLink(t *testing.T) {
	head := NewList[int]()
	assert.NotNil(t, head)
	assert.Equal(t, head.total, 0)

	head = NewList[int](13, 1, 3)
	assert.Equal(t, head.Count(), 3)

}

func TestLinkedListInsert(t *testing.T) {
	head := NewList[int]()

	// insert data in  illegal position
	assert.NotNil(t, head)

	head.Insert(3, 0)
	assert.NotNil(t, head)
}

func TestLinkedListForeach(t *testing.T) {
	head := NewList[int](genRandomArr(10)...)
	head.Foreach(func(n *Node[int], idx int) {
		fmt.Printf("%d - %d\n", idx, n.Data)
	})
}

func TestListReverse(t *testing.T) {
	data := genRandomArr(100000)
	head := NewList[int](data...)
	head.Reverse()
	head.Foreach(func(n *Node[int], idx int) {
		assert.Equal(t, data[len(data)-1-idx], n.Data)
	})
}

func genRandomArr(length int) []int {
	const min, max = -9999999, 9999999
	var arr = make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(max-min) + min
	}
	return arr
}
