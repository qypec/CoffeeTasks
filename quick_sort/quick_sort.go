package main

import (
	"fmt"
)

func partition(arr []int, l, r int) int {
	x := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] <= x {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[l] = arr[l], arr[j]
	return j
}

func quick_sort(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := partition(arr, l, r)
	quick_sort(arr, m + 1, r)
	quick_sort(arr, l, m - 1)
}


func main() {
	arr := []int{1, 3, 2, 4}
	quick_sort(arr, 0, len(arr) - 1)
	fmt.Println("1: ", arr)
	
	arr = []int{1, 3, 2}
	quick_sort(arr, 0, len(arr) - 1)
	fmt.Println("2: ", arr)
	
	arr = []int{1, 3, 2, 13, 0, 4, 78, 8, 1, 1, 1}
	quick_sort(arr, 0, len(arr) - 1)
	fmt.Println("3: ", arr)
}
