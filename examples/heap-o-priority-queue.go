//go:build ignore

package examples

import "container/heap"

/*
================================================================
Algorithm: Heap / Priority Queue
================================================================

[EASY] LeetCode 703 - Kth Largest Element in a Stream
Design a class that, on add(val), returns the kth largest seen so far.
Example: k=3, stream=4,5,8,2,3,5,10,9,4 -> outputs the running 3rd largest.
Hint: Keep a min-heap of size k. On add, push and pop while size > k. The
top is the kth largest.

[EASY] LeetCode 1046 - Last Stone Weight
Repeatedly smash the two heaviest stones; if equal, both vanish; else
diff goes back. Return the final remaining stone weight (or 0).
Example: stones = [2,7,4,1,8,1]  ->  1
Hint: Use a max-heap. Pop two, push abs(diff) if non-zero, repeat until
size <= 1.

[EASY] LeetCode 1337 - The K Weakest Rows in a Matrix
Each row has soldiers (1's) at the front. Return indices of the k weakest
rows (fewer soldiers first; ties by smaller index).
Example: mat=[[1,1,0,0,0],[1,1,1,1,0],[1,0,0,0,0],[1,1,0,0,0],[1,1,1,1,1]],
k=3  ->  [2,0,3]
Hint: Compute (count, idx) per row. Keep a max-heap of size k by (count,
idx) so largest is on top; replace when smaller arrives. Sort the final k.

[MEDIUM] LeetCode 347 - Top K Frequent Elements
Return the k most frequent elements.
Example: nums=[1,1,1,2,2,3], k=2  ->  [1,2]
Hint: Count frequencies in a map. Push (count, num) into a min-heap of
size k; replace the top if a more frequent element comes along.

[HARD] LeetCode 295 - Find Median from Data Stream
Support addNum(num) and findMedian() with median in O(log n) and O(1).
Example: addNum(1), addNum(2), findMedian()->1.5; addNum(3), findMedian()->2
Hint: Two heaps — `low` is a max-heap holding the smaller half, `high` is
a min-heap holding the larger half. Keep |low| - |high| in {0, 1} and
median is either low.top or (low.top + high.top)/2.
*/

// ===================== Answers (Optimal Solutions) =====================

// minHeap of ints (used for KthLargest and median's `high`)
type intMinHeap []int

func (h intMinHeap) Len() int            { return len(h) }
func (h intMinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intMinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intMinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

// maxHeap of ints (used for median's `low`)
type intMaxHeap []int

func (h intMaxHeap) Len() int            { return len(h) }
func (h intMaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h intMaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intMaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

// 1) [EASY] LC 703: Kth Largest Element in a Stream -- O(log k) per add
type KthLargest struct {
	k int
	h *intMinHeap
}

func ConstructorKthLargest(k int, nums []int) KthLargest {
	h := &intMinHeap{}
	heap.Init(h)
	kl := KthLargest{k: k, h: h}
	for _, v := range nums {
		kl.Add(v)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl.h, val)
	if kl.h.Len() > kl.k {
		heap.Pop(kl.h)
	}
	return (*kl.h)[0]
}

// 2) [EASY] LC 1046: Last Stone Weight -- O(n log n)
func lastStoneWeight(stones []int) int {
	h := &intMaxHeap{}
	heap.Init(h)
	for _, s := range stones {
		heap.Push(h, s)
	}
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		if a != b {
			heap.Push(h, a-b)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return (*h)[0]
}

// 3) [EASY] LC 1337: The K Weakest Rows in a Matrix -- O(m log k + k log k)
type rowItem struct {
	cnt, idx int
}
type rowMaxHeap []rowItem

func (h rowMaxHeap) Len() int { return len(h) }
func (h rowMaxHeap) Less(i, j int) bool {
	if h[i].cnt != h[j].cnt {
		return h[i].cnt > h[j].cnt
	}
	return h[i].idx > h[j].idx
}
func (h rowMaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *rowMaxHeap) Push(x interface{}) { *h = append(*h, x.(rowItem)) }
func (h *rowMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func kWeakestRows(mat [][]int, k int) []int {
	h := &rowMaxHeap{}
	heap.Init(h)
	for i, row := range mat {
		c := 0
		for _, v := range row {
			c += v
		}
		heap.Push(h, rowItem{c, i})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	out := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		out[i] = heap.Pop(h).(rowItem).idx
	}
	return out
}

// 4) [MEDIUM] LC 347: Top K Frequent Elements -- O(n log k)
type freqItem struct {
	num, cnt int
}
type freqHeap []freqItem

func (h freqHeap) Len() int            { return len(h) }
func (h freqHeap) Less(i, j int) bool  { return h[i].cnt < h[j].cnt }
func (h freqHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *freqHeap) Push(x interface{}) { *h = append(*h, x.(freqItem)) }
func (h *freqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func topKFrequent(nums []int, k int) []int {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	h := &freqHeap{}
	heap.Init(h)
	for num, c := range cnt {
		heap.Push(h, freqItem{num, c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	out := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		out[i] = heap.Pop(h).(freqItem).num
	}
	return out
}

// 5) [HARD] LC 295: Find Median from Data Stream -- O(log n) add, O(1) find
type MedianFinder struct {
	low  *intMaxHeap // bottom half
	high *intMinHeap // top half
}

func ConstructorMedianFinder() MedianFinder {
	return MedianFinder{low: &intMaxHeap{}, high: &intMinHeap{}}
}

func (m *MedianFinder) AddNum(num int) {
	heap.Push(m.low, num)
	heap.Push(m.high, heap.Pop(m.low))
	if m.high.Len() > m.low.Len() {
		heap.Push(m.low, heap.Pop(m.high))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.low.Len() > m.high.Len() {
		return float64((*m.low)[0])
	}
	return (float64((*m.low)[0]) + float64((*m.high)[0])) / 2.0
}
