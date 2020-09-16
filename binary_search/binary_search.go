package main

import (
	"fmt"
)

// binary_search findes first `toFind` in array `arr`
// `arr` must be ascending sorted
// returns len(arr) if `toFind` not found
// Complexity: log(n)
func binary_search(arr []int, toFind int) int {
	// return rec_binary_search(arr, toFind, 0, len(arr) - 1)

	l := 0
	r := len(arr) - 1
	fix := len(arr)
	for l <= r {
		mid := l + int((r-l)/2)
		if arr[mid] >= toFind {
			if arr[mid] == toFind {
				fix = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return fix
}

func main() {
	fmt.Println("1: ", binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println("2: ", binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9))
	fmt.Println("3: ", binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1))

	fmt.Println("4: ", binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println("5: ", binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6))
	fmt.Println("6: ", binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println("7: ", binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))

	fmt.Println("8: ", binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100))
}
