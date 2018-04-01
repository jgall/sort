package sort

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestCompHeap(t *testing.T) {
	h := &CompHeap{Int{2}, Int{1}, Int{4}}
	heap.Init(h)
	heap.Push(h, Int{3})
	res := ""
	for h.Len() > 0 {
		res += fmt.Sprintf("%d ", heap.Pop(h))
	}
	if res != "{1} {2} {3} {4} " {
		t.Fail()
	}
}

type Int struct {
	val int
}

func (i Int) Less(j Comparable) bool {
	return i.val < j.(Int).val
}
