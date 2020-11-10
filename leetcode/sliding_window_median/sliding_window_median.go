package main

// description https://leetcode.com/problems/sliding-window-median/

import (
	"container/heap"
	"math"
	"sort"
)

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func isEven(k int) bool {
	if k%2 == 0 {
		return true
	}
	return false
}

func getMedian(lessHeap *IntMaxHeap, moreHeap *IntMinHeap, k int) float64 {
	if isEven(k) {
		return (float64((*lessHeap)[0]) + float64((*moreHeap)[0])) / 2
	}

	if lessHeap.Len() > moreHeap.Len() {
		return float64((*lessHeap)[0])
	}
	return float64((*moreHeap)[0])
}

func medianSlidingWindow(nums []int, k int) []float64 {
	window := make([]int, k)
	copy(window, nums[:k])
	sort.Ints(window)

	lessHeap := &IntMaxHeap{}
	moreHeap := &IntMinHeap{}
	heap.Init(lessHeap)
	heap.Init(moreHeap)
	for i := 0; i < int(math.Ceil(float64(k)/2)); i++ {
		heap.Push(lessHeap, window[i])
	}
	for i := int(math.Ceil(float64(k) / 2)); i < k; i++ {
		heap.Push(moreHeap, window[i])
	}

	median := make([]float64, 0, len(nums)-k+1)
	median = append(median, getMedian(lessHeap, moreHeap, k))
	for i := 1; i < len(nums)-k+1; i++ {
		deleteItem := nums[i-1]
		appendItem := nums[i+k-1]

		if lessHeap.Len() > 0 && deleteItem <= (*lessHeap)[0] {
			for i := 0; i < lessHeap.Len(); i++ {
				if (*lessHeap)[i] == deleteItem {
					heap.Remove(lessHeap, i)
					break
				}
			}
		} else {
			for i := 0; i < moreHeap.Len(); i++ {
				if (*moreHeap)[i] == deleteItem {
					heap.Remove(moreHeap, i)
					break
				}
			}
		}

		if lessHeap.Len() > 0 && appendItem <= (*lessHeap)[0] {
			heap.Push(lessHeap, appendItem)
			if lessHeap.Len()-moreHeap.Len() >= 2 {
				heap.Push(moreHeap, heap.Pop(lessHeap))
			}
		} else {
			heap.Push(moreHeap, appendItem)
			if moreHeap.Len()-lessHeap.Len() >= 2 {
				heap.Push(lessHeap, heap.Pop(moreHeap))
			}
		}
		median = append(median, getMedian(lessHeap, moreHeap, k))
	}
	return median
}
