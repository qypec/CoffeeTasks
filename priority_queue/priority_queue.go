/*
	Первая строка входа содержит число операций.
	Каждая из последующих n строк задают операцию одного из следующих двух типов:
		Insert x
		ExtractMax
	Первая операция добавляет число x в очередь с приоритетами, вторая — извлекает максимальное число и выводит его.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func swap(a *int, b *int) { *a, *b = *b, *a }

func OutOfRange(x, l, r int) bool {
	if x >= l && x <= r {
		return false
	}
	return true
}

type Element struct {
	value int
	index int
}

type PriorityQueue struct {
	arr []int
}

func (p *PriorityQueue) siftingUp() {
	child := p.Back()
	if child == nil {
		return
	}
	for parent := p.GetParent(child); parent != nil; parent = p.GetParent(child) {
		if child.value > parent.value {
			swap(&p.arr[child.index], &p.arr[parent.index])
			swap(&child.index, &parent.index)
		} else {
			break
		}
	}
}

func (p PriorityQueue) GetParent(child *Element) *Element {
	parentIndex := int(child.index / 2)
	if parentIndex != 0 {
		return &Element{p.arr[parentIndex], parentIndex}
	}
	return nil
}

func (p *PriorityQueue) siftingDown() {
	parent := p.Front()
	if parent == nil {
		return
	}
	for child := p.GetChild(parent); child != nil; child = p.GetChild(parent) {
		if child.value > parent.value {
			swap(&p.arr[child.index], &p.arr[parent.index])
			swap(&child.index, &parent.index)
		} else {
			break
		}
	}
}

func (p PriorityQueue) GetChild(parent *Element) *Element {
	childIndexLeft, childIndexRight := int(parent.index*2), int(parent.index*2+1)
	if OutOfRange(childIndexLeft, 1, p.Len()) && OutOfRange(childIndexRight, 1, p.Len()) {
		return nil
	} else if OutOfRange(childIndexRight, 1, p.Len()) {
		return &Element{p.arr[childIndexLeft], childIndexLeft}
	} else {
		if p.arr[childIndexLeft] > p.arr[childIndexRight] {
			return &Element{p.arr[childIndexLeft], childIndexLeft}
		} else {
			return &Element{p.arr[childIndexRight], childIndexRight}
		}
	}
	return nil
}

func (p *PriorityQueue) Init() {
	p.arr = make([]int, 1)
	p.arr[0] = 0
}

func (p PriorityQueue) Len() int { return p.arr[0] }

func (p PriorityQueue) Back() *Element {
	if p.Len() == 0 {
		return nil
	}
	return &Element{p.arr[p.Len()], p.Len()}
}

func (p PriorityQueue) Front() *Element {
	if p.Len() == 0 {
		return nil
	}
	return &Element{p.arr[1], 1}
}

func (p *PriorityQueue) Insert(x int) {
	p.arr = append(p.arr, x)
	p.arr[0]++
	p.siftingUp()
}

func (p *PriorityQueue) ExtractMax() int {
	if p.Len() == 0 {
		return 0
	}
	max := p.arr[1]
	swap(&p.arr[1], &p.arr[p.Back().index])
	p.arr = p.arr[:p.Len()]
	p.arr[0]--
	p.siftingDown()
	return max
}

func main() {
	var pQueue PriorityQueue

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	pQueue.Init()
	result := make([]int, 0)
	for scanner.Scan() {
		switch scanner.Text() {
		case "Insert":
			scanner.Scan()
			x, _ := strconv.Atoi(scanner.Text())
			pQueue.Insert(x)
		case "ExtractMax":
			result = append(result, pQueue.ExtractMax())
		}
	}

	for _, elem := range result {
		fmt.Println(elem)
	}
}
