package main

import (
	"container/heap"
	"fmt"
)

type myPeap []int

func (h myPeap) Len() int           { return len(h) }
func (h myPeap) Less(i, j int) bool { return h[i] < h[j] }
func (h myPeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *myPeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *myPeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type myHeap []int

func (h myHeap) Len() int           { return len(h) }
func (h myHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h myHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *myHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *myHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func runHeap(arr []int) []int {
	h := &myHeap{}
	p := &myPeap{}
	heap.Init(h)
	heap.Init(p)
	for i := range arr {
		heap.Push(h, arr[i])
		heap.Push(p, arr[i])
		fmt.Println(h, p)
	}
	return arr
}
