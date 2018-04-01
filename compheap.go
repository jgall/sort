package sort

// A CompHeap is a min-heap of Comparables.
type CompHeap []Comparable

func (h CompHeap) Len() int           { return len(h) }
func (h CompHeap) Less(i, j int) bool { return h[i].Less(h[j]) }
func (h CompHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CompHeap) Push(x interface{}) {
	*h = append(*h, x.(Comparable))
}

func (h *CompHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
