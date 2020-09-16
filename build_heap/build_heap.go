/*
	Переставить элементы заданного массива чисел так,
	чтобы он удовлетворял свойству мин-кучи.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func outOfRange(x, l, r int) bool {
	if x >= l && x <= r {
		return false
	}
	return true
}

const MinHeap = 1

type Heap struct {
	arr  []element
	size int
	kind int

	swapElem [][]int
}

type element struct {
	index    int
	priority int

	value interface{}
}

// Builds a min-heap from a priority array.
// The value of the element with priority[i] is equal to values[i]
// Complexity: n
// size of priority and values must be equals
func BuildMinHeap(priority []int, values []interface{}) *Heap {
	var h Heap

	h.Init(MinHeap)
	for i := 0; i < len(priority); i++ {
		h.arr = append(h.arr, element{i + 1, priority[i], values[i]})
		h.size++
	}
	h.heapify()
	return &h
}

// initializes a heap of size 0.
func (h *Heap) Init(heapType int) *Heap {
	h.arr = make([]element, 1)
	h.arr[0] = element{0, 0, nil}
	h.size = 0
	h.kind = heapType
	h.swapElem = make([][]int, 0)
	return h
}

// normalizes a heap. Element order respects heap properties
func (h *Heap) heapify() {
	for i := int(h.size / 2); i > 0; i-- {
		h.siftingDown(i)
	}
}

func (h *Heap) siftingSwapCondition(childPriority, parentPriority int) bool {
	if h.kind == MinHeap {
		return childPriority < parentPriority
	} else {
		return childPriority > parentPriority
	}
}

// returns the largest(MaxHeap)/less(MinHeap) child of this element
func (h Heap) getChild(parent *element) *element {
	childIndexLeft, childIndexRight := int(parent.index*2), int(parent.index*2+1)
	if outOfRange(childIndexLeft, 1, h.size) && outOfRange(childIndexRight, 1, h.size) {
		return nil
	} else if outOfRange(childIndexRight, 1, h.size) {
		return &h.arr[childIndexLeft]
	} else {
		if h.siftingSwapCondition(h.arr[childIndexLeft].priority, h.arr[childIndexRight].priority) {
			return &h.arr[childIndexLeft]
		} else {
			return &h.arr[childIndexRight]
		}
	}
}

// swap elements inside heap.arr
// heap.index remains untouched
func (h *Heap) swap(child, parent **element) {
	childIndex := (*child).index
	parentIndex := (*parent).index

	h.swapElem = append(h.swapElem, []int{parentIndex - 1, childIndex - 1})

	h.arr[childIndex], h.arr[parentIndex] = h.arr[parentIndex], h.arr[childIndex]
	h.arr[parentIndex].index, h.arr[childIndex].index = parentIndex, childIndex

	*child = &h.arr[parentIndex]
	*parent = &h.arr[childIndex]
}

// moves an element down until it satisfies the properties of the heap
func (h *Heap) siftingDown(index int) {
	if outOfRange(index, 1, h.size) {
		return
	}

	parent := &h.arr[index]
	if parent == nil {
		return
	}
	for child := h.getChild(parent); child != nil; child = h.getChild(parent) {
		if h.siftingSwapCondition(child.priority, parent.priority) {
			h.swap(&child, &parent)
		} else {
			break
		}
	}
}

func GetValues(priority []int) []interface{} {
	values := make([]interface{}, len(priority))

	for i := 0; i < len(priority); i++ {
		values[i] = priority[i]
	}
	return values
}

func buildHeap(arr []int) [][]int {
	heap := BuildMinHeap(arr, GetValues(arr))
	return heap.swapElem
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	result := buildHeap(arr)
	fmt.Println(len(result))
	for _, elem := range result {
		fmt.Println(elem[0], elem[1])
	}
}
