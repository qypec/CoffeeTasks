/*
	Найти максимум в каждом окне размера m данного массива чисел A[1 . . . n].

	Вход. Массив чисел A[1 . . . n] и число 1 ≤ m ≤ n

	Выход. n − m + 1 максимумов, разделённых пробелами.
*/

package main

import (
	"container/list"
	"fmt"
	"math"
)

type StackValue struct {
	elem int
	max  int
}

// implementation of a queue on two stacks
type MyQueue struct {
	addStack *list.List
	delStack *list.List
}

func (this *MyQueue) Init() {
	this.addStack = list.New()
	this.delStack = list.New()
}

// return max element of addStack
func (this MyQueue) getMaxAddStack() int { return this.addStack.Back().Value.(*StackValue).max }

// return max element of delStack
func (this MyQueue) getMaxDelStack() int { return this.delStack.Back().Value.(*StackValue).max }

func (MyQueue) pushBack(stack *list.List, elem int) {
	nextMax := elem
	if stack.Len() != 0 && nextMax < stack.Back().Value.(*StackValue).max {
		nextMax = stack.Back().Value.(*StackValue).max
	}
	stack.PushBack(&StackValue{elem, nextMax})
}

// move all elements from addStack to delStack
func (this *MyQueue) moveToDelStack() {
	for e := this.addStack.Back(); e != nil; e = e.Prev() {
		this.pushBack(this.delStack, e.Value.(*StackValue).elem)
	}
	this.addStack = list.New()
}

// push element to the end of queue
func (this *MyQueue) PushBack(elem int) { this.pushBack(this.addStack, elem) }

// remove first element of queue
func (this *MyQueue) PopFront() {
	if this.delStack.Len() == 0 && this.addStack.Len() != 0 {
		this.moveToDelStack()
	}
	this.delStack.Remove(this.delStack.Back())
}

// return max element of queue
func (this *MyQueue) GetMax() int {
	if this.addStack.Len() == 0 {
		return this.delStack.Back().Value.(*StackValue).max
	} else if this.delStack.Len() == 0 {
		return this.addStack.Back().Value.(*StackValue).max
	} else {
		return int(math.Max(float64(this.getMaxAddStack()), float64(this.getMaxDelStack())))
	}
}

func main() {
	var q MyQueue
	var N, M int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&M)

	maxOfSlidingWindow := make([]int, 0)
	q.Init()
	for i := 0; i < M-1; i++ { // push the elements of the first window to the queue
		q.PushBack(arr[i])
	}
	for i := M - 1; i < N; i++ {
		q.PushBack(arr[i])
		maxOfSlidingWindow = append(maxOfSlidingWindow, q.GetMax())
		q.PopFront()
	}

	for i, elem := range maxOfSlidingWindow {
		fmt.Printf("%d", elem)
		if i+1 != len(maxOfSlidingWindow) {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}
