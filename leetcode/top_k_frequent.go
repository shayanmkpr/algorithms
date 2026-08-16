package main

import (
	"container/heap"
	"fmt"
)

// keeping the Postion of the nums in the heap
// heap.Fix()
// implementing not a normal traditional heap

type node struct {
	Key int
	Val int
}
type mapHeap struct {
	Nodes []*node
	Pos   map[int]int
}

func (h *mapHeap) Push(x any) { // push, and update the position as well
	h.Nodes = append(h.Nodes, x.(*node))
	h.Pos[x.(*node).Key] = len(h.Nodes) - 1
}

func (h *mapHeap) Pop() any {
	old := h.Nodes
	n := len(h.Nodes)
	item := old[n-1]
	h.Nodes = old[:n-1]
	delete(h.Pos, item.Key)
	return item
}

func (h *mapHeap) Len() int {
	return len(h.Nodes)
}

func (h *mapHeap) Less(i, j int) bool { // swap positions as well
	return h.Nodes[i].Val > h.Nodes[j].Val
}

func (h *mapHeap) Swap(i, j int) {
	h.Nodes[i], h.Nodes[j] = h.Nodes[j], h.Nodes[i]
	h.Pos[h.Nodes[i].Key] = i
	h.Pos[h.Nodes[j].Key] = j
}
func topKFrequentHeap(nums []int, k int) []int {
	fmt.Println(nums)
	var result []int
	seen := map[int]int{}
	newHeap := &mapHeap{
		Nodes: []*node{},
		Pos:   map[int]int{},
	}
	heap.Init(newHeap)

	for _, num := range nums {
		fmt.Println(num)
		seen[num]++
		if idx, ok := newHeap.Pos[num]; ok {
			newHeap.Nodes[idx].Val = seen[num]
			heap.Fix(newHeap, idx)
		} else {
			heap.Push(newHeap, &node{ // updates pos by itslef
				Key: num,
				Val: seen[num],
			})
		}
	}
	fmt.Println(newHeap)
	fmt.Println(seen)
	for range k {
		result = append(result, heap.Pop(newHeap).(*node).Key)
	}
	return result
}

func topKFrequent(nums []int, k int) []int {
	var result []int
	seen := map[int]int{}
	for _, val := range nums {
		seen[val]++
	}

	var buckets [][]int
	for key, cnt := range seen {
		buckets[cnt] = append(buckets[cnt], seen[key])
	}
	fmt.Println(buckets)
	return result
}
