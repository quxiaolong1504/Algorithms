// Implementation of basic singly linked list
// Reference: https://en.wikipedia.org/wiki/Linked_list

package singlyLinkedList

import (
	"github.com/quxiaolong1504/Algorithms/constraints"
)

type List[T constraints.Ordered] struct {
	total int
	head  *Node[T]
	tail  *Node[T]
}

type Node[T constraints.Ordered] struct {
	Next *Node[T]
	Data T
}

func NewList[T constraints.Ordered](datas ...T) *List[T] {
	head := &List[T]{}
	for i := len(datas); i > 0; i-- {
		head.Insert(datas[i-1], 0)
	}
	return head
}

func (h *List[T]) Count() int {
	return h.total
}

func (h *List[T]) Insert(d T, position int) {
	// 插入头部
	if position <= 0 {
		h.head = &Node[T]{Data: d, Next: h.head}
		if h.tail == nil {
			h.tail = h.head
		}
		h.total++
		return
	}

	// 插入尾部
	if position > h.total {
		h.tail.Next = &Node[T]{Data: d}
		h.tail = h.tail.Next
		h.total++
		return
	}

	// 插入中间
	var pre = h.head
	for count := 0; count < position-1; count++ {
		pre = pre.Next
	}
	pre.Next = &Node[T]{Data: d, Next: pre.Next}
	h.total++
}

func (h *List[T]) Foreach(fun func(node *Node[T], idx int)) {
	var p = h.head
	var idx int
	for p != nil {
		fun(p, idx)
		idx++
		p = p.Next
	}
}

func (h *List[T]) Reverse() {
	if h.Count() < 2 {
		return
	}

	// 1,2,3,4,5  -> 5,4,3,2,1
	var prev *Node[T]
	var cur = h.head
	for cur != nil {
		// 记录当前元素的下一个
		next := cur.Next

		// 修改当前元素为 prev
		cur.Next = prev

		// prev 后移
		prev = cur
		cur = next
	}
	h.head, h.tail = h.tail, h.head
}
